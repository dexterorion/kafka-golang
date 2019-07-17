https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26

- MY_IP=your-ip docker-compose up
- ifconfig -a
- docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic foo --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
- brew install kafkakat
- kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0
- new terminal
- echo 'publish to partition 0' | kafkacat -P -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0