package main

import (
	"flag"
	"log"
	"os"
	"perfomance-kafka_go/kafka"
)

func main() {
	// brokers := []string{"kafka:9092"} // []string{"localhost:29092"}
	// topic := "kafka-performance"
	// topic := "kafka-scalling"

	brokers := []string{"kafka1:9092"}
	// brokers := []string{"kafka1:9092", "kafka2:9094", "kafka3:9096"}
	topic := "kafka-distribusi"

	flag.Usage = func() {
		log.Printf("Usage:")
		log.Println(" ================== consumer ================== ")
		log.Printf("go run . consume1")
		log.Printf("go run . consume2")
		log.Printf("go run . consume3 group")
		log.Println(" ================== producer ================== ")
		log.Printf("go run . produce1")
		log.Printf("go run . produce2")
		log.Printf("go run . produce-compress")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	totalData := 1000 // kirim x pesan
	switch flag.Args()[0] {
	case "produce1":
		kafka.Produce1(brokers, topic, totalData)
	case "produce2":
		kafka.Produce2(brokers, topic, totalData)
	case "produce-compress":
		kafka.ProduceCompressed(brokers, topic, "aaaaaaaaaa adalah contoh payload untuk uji coba normal dan compress aaaaaaaaaa", totalData)

	case "consume1":
		kafka.Consume1(brokers, topic, "baseline-group", totalData)
	case "consume2":
		kafka.Consume2(brokers, topic, "baseline-group", totalData)
	case "consume3":
		kafka.Consume3(brokers, topic, flag.Args()[1], totalData)
	default:
		panic("unknown command")
	}
}
