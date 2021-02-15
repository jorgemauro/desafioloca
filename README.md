# desafioloca

Buscando uma aplicação flexivel para outros formatos de comunicação
busquei utilizar o padrão ports and adapters
- Estrutura do serverloca:
  - Camada de aplicação
  - Factory - instancia objetos com muitas dependencias
  - Kafka - consumo e processamento de trasações com o apache Kafka
  - grpc - Estabelecimento das dos protocolos de gRPC para facilitar consultas com alguns protocolos provenientes da tecnologia
  - model - Estruturas que vão receber  as requisições externas
  - usecase - executa o fluxo de ações de acordo com as regras do negocio

- CMD
  - Comandos registrados para iniciars serviços(CLI)

- Domain
  - dominio da aplicação e suas regras de negocio

- Infrastructure
 - db -> configurações e interface com  o banco de dados
 - repository -> realiza a persistencia dos dados
--------------------------------------------------------
- location-api ---> utilização do nesteJs responsavel por ser o intermediario entre nosso dominio e o front-end
--------------------------------------------------------
- location-frontend ---> utilização do nextJS para ter uma das tecnologias para tornar a experiência do usuarios mais rapido aproveitando o com server side rendering