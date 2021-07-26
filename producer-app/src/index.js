const { Kafka } = require('kafkajs');

const kafka = new Kafka({
    clientId: 'nodejs-kafka-app',
    brokers: ['kafka_kafka_1:9092'],
    connectionTimeout: 10000,
    retry: {
        initialRetryTime: 600,
        retries: 10
    }
})

const producer = kafka.producer()
const consumer = kafka.consumer({ groupId: 'test-group' })

const run = async () => {

    // Producing
    await producer.connect()
    await producer.send({
      topic: 'test-topic',
      messages: [
        { value: 'Hello KafkaJS user!' },
      ],
    })

    // Consuming
    await consumer.connect()
    await consumer.subscribe({ topic: 'test-topic', fromBeginning: true })
    
    await consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
        console.log({
            partition,
            offset: message.offset,
            value: message.value.toString(),
        })
        },
    })
}
 
run().catch(console.error)