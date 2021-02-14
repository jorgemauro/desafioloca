package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/jorgemauro/desafioloca/serverloca/application/factory"
	appmodel "github.com/jorgemauro/desafioloca/serverloca/application/model"
	"github.com/jorgemauro/desafioloca/serverloca/application/usecase"
	"github.com/jorgemauro/desafioloca/serverloca/domain/model"
)

type KafkaProcessor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}
	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaLocationTopic"), os.Getenv("kafkaLocationConfirmationTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
			k.processMessage(msg)
		}
	}
}

func (k *KafkaProcessor) processMessage(msg *ckafka.Message) {
	locationsTopic := "locations"
	locationConfirmationTopic := "location_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case locationsTopic:
		k.processLocation(msg)
	case locationConfirmationTopic:
		k.processLocationConfirmation(msg)
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (k *KafkaProcessor) processLocation(msg *ckafka.Message) error {
	location := appmodel.NewLocation()
	err := location.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	locationUseCase := factory.LocationUseCaseFactory(k.Database)

	createdLocation, err := locationUseCase.Register(
		location.AccountID,
		location.Amount,
		location.PixKeyTo,
		location.PixKeyKindTo,
		location.Description,
		location.ID,
	)
	if err != nil {
		fmt.Println("error registering location", err)
		return err
	}

	topic := "bank" + createdLocation.PixKeyTo.Account.Bank.Code
	location.ID = createdLocation.ID
	location.Status = model.LocationPending
	locationJson, err := location.ToJson()

	if err != nil {
		return err
	}

	err = Publish(string(locationJson), topic, k.Producer, k.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func (k *KafkaProcessor) processLocationConfirmation(msg *ckafka.Message) error {
	location := appmodel.NewLocation()
	err := location.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	locationUseCase := factory.LocationUseCaseFactory(k.Database)

	if location.Status == model.LocationConfirmed {
		err = k.confirmLocation(location, locationUseCase)
		if err != nil {
			return err
		}
	} else if location.Status == model.LocationCompleted {
		_, err := locationUseCase.Complete(location.ID)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (k *KafkaProcessor) confirmLocation(location *appmodel.Location, locationUseCase usecase.LocationUseCase) error {
	confirmedLocation, err := locationUseCase.Confirm(location.ID)
	if err != nil {
		return err
	}

	topic := "bank" + confirmedLocation.AccountFrom.Bank.Code
	locationJson, err := location.ToJson()
	if err != nil {
		return err
	}

	err = Publish(string(locationJson), topic, k.Producer, k.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
