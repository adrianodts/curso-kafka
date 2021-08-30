# curso-kafka
Curso de introdução ao kafka.

## Lib para interface com o kafka
Github [librdkafka](https://github.com/edenhill/librdkafka)

## Criando tópico via linha de comando:
kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3

## Consumindo tópico via linha de comando:
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

## Produzindo mensagem em um tópico via linha de comando:
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste
>[digitar mensagem] + enter

