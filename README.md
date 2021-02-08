# desafioloca

utilizando GO para subir o serviço e consumir o apache kafka
ddd -> para separar a complexidade do dominio(negocio) e a complexidade.
Buscando uma aplicação flexivel para outros formatos de comunicação
utilizando assim ports and adapters

Camada de aplicação
Factory - instancia objetos com muitas dependencias
Kafka - consumo e processamento de trasações com o apache Kafka
model - Estruturas que vão receber  as requisições externas
usecase - executa o fluxo de ações de acordo com as regras do negocio

CMD
Comandos registrados para iniciars serviços(CLI)

Domain
dominio da aplicação e suas regras de negocio

Infrastructure
db - configurações e interface com  o banco de dados
repository - realiza a persistencia dos dados
--------------------------------------------------------
Front-End
utilização do nextJS para ter uma das tecnologias para tornar a experiência do usuarios mais rapido

API
Utilização do NestJS para cumprir o papel de bff