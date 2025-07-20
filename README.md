# simple-go-kafka
simple Go REST API (Kafka Producer) and  Go Kafka Consumer

#Producer Set up
docker-compose up --build 

#GO server running
go run main.go

#Consumer Set up
docker-compose ps

#consumer check
docker-compose logs -f consumer
