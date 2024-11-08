**Go e Clean Architecture**  

**Descrição:**  
Este projeto tem como objetivo a inserção e consulta de registros em um banco de dados Mysql utilizando a linguagem GO e abordagem clean architecture.  
Ao criar um pedido além de gravar os dados no Mysql também será criada uma mensagem no sistema de mensageria RabbitMQ.  
As consultas e as inserções poderão ser realizadas via Api Rest, GRPC e GRAPHQL.


**Execução**

O mysql, o rabbitMQ e a aplicação serão inicializados com o comando docker compose up -d.

**Acesso as funcionalidades**

* REST: utilize o arquivo order.http na pasta api, que possui exemplos de como realizar uma requisição POST e GET, para inserção e consulta respectivamente.
* GRAPHQL: acesse o endereço http://localhost:8080/ para utilizar o playground, e então acionar o botão na tela e escolher queryOrders para consultar o banco de dados ou createOrders para inserir um registro.  ![playground graphql](/imagens/playground_graphql.png)  
* GRPC: utilize um cliente GRPC como o evans, para acessar os serviços.
  * inicialize o evans com o comando 'evans -r repl'.  ![evans iniciar](/imagens/evans_inicio.png)
  * para listar os registros será necessário os seguintes comandos: package pb, service OrderService, call ListOrder.  ![evans ListOrder](/imagens/evans_ListOrder.png)
  * para inserir os registros será necessário os seguintes comandos: package pb, service OrderService, call CreateOrder.  ![evans CreateOrder](/imagens/evans_CreateOrder.png)


