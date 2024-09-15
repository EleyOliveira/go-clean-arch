# Go e Clean Architecture

Este projeto tem como objetivo a inserção e consulta de informações referente a um pedido no banco de dados Mysql utilizando a linguagem GO e clean architecture.
Ao criar um pedido além de gravar os dados no Mysql também será criada uma mensagem no sistema de mensageria RabbitMQ.
As consultas e as inserções poderão ser realizadas via Api Rest, GRPC e GRAPHQL.
Tanto o Mysql quando o RabbitMQ estão utilizando imagens Docker, conforme descrito no arquivo docker-compose.yaml.

Quando os containers do MySql e do RabbitMQ estiverem em execução, pode-se utilizar o arquivo Makefile para realizar a migration no MySQL.

Para subirem os servidores da Api, do GRPC e do GRAPHQL executar o comando "go run main.go wire_gen.go" dentro da pasta "cmd".

Esses serviços executam no localhost, porta 8000 para a API Rest, porta 50051 GRPC e porta 8080 para GRAPHQL.
Na pasta "api", o arquivo order.http possui um exemplo de chamada para inserção e consulta.


