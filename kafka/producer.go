package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/snappy"
	k "github.com/segmentio/kafka-go"
)

func Produce1(brokers []string, topic string, n int) {
	w := k.NewWriter(k.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &k.RoundRobin{}, // ⬅️ jelas: round-robin ke semua partition
	})
	defer w.Close()
	start := time.Now()

	for i := 0; i < n; i++ {
		msg := k.Message{
			Key:   nil, // ⬅️ kosongin biar balancer bekerja
			Value: []byte(fmt.Sprintf("value-%d", i)),
		}

		if err := w.WriteMessages(context.Background(), msg); err != nil {
			log.Fatal("failed to write:", err)
		}
		log.Printf("Message[%d] sent successfully!", i)
	}

	elapsed := time.Since(start)
	fmt.Printf("✅ Produced %d messages in %s (%.2f msg/sec)\n",
		n, elapsed, float64(n)/elapsed.Seconds())
}

func Produce2(brokers []string, topic string, n int) {
	w := k.NewWriter(k.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &k.RoundRobin{},
		// ⬇️ tambahin config untuk batch
		BatchSize:    100,                   // kumpulin 100 msg per flush
		BatchTimeout: 50 * time.Millisecond, // atau flush tiap 50ms
	})
	defer w.Close()
	start := time.Now()

	// kumpulin messages
	msgs := make([]k.Message, 0, n)
	for i := 0; i < n; i++ {
		msgs = append(msgs, k.Message{
			Key:   nil,
			Value: []byte(fmt.Sprintf("value-%d", i)),
		})
	}
	log.Printf("Total:%d", len(msgs))

	// kirim sekaligus
	if err := w.WriteMessages(context.Background(), msgs...); err != nil {
		log.Fatal("failed to write:", err)
	}

	elapsed := time.Since(start)
	log.Printf("✅ Produced %d messages in %s (%.2f msg/sec)\n",
		n, elapsed, float64(n)/elapsed.Seconds())
}

func ProduceCompressed(brokers []string, topic, data string, n int) {
	writer := &k.Writer{
		Addr:         k.TCP(brokers...),
		Topic:        topic,
		Balancer:     &k.LeastBytes{},       // Load balancing: kirim message ke partition yang paling sedikit datanya
		RequiredAcks: k.RequireOne,          // Acknowledgement: Producer menunggu ack dari minimal 1 broker
		Compression:  k.Snappy,              // Compression pakai Snappy (lebih hemat storage & bandwidth)
		BatchSize:    200,                   // Batching: kirim message setelah terkumpul x record
		BatchTimeout: 50 * time.Millisecond, // atau setelah 50ms walaupun belum penuh x record
	}
	defer writer.Close()

	start := time.Now()
	var totalOriginalSize int
	var totalCompressedSize int
	for i := 0; i < n; i++ {
		payload:=fmt.Sprintf("%s-%d",data,i)
		msg := k.Message{Value: []byte(payload)}
		totalOriginalSize += len(payload)
		totalCompressedSize += len(snappy.Encode(nil, []byte(payload)))

		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			fmt.Println("error:", err)
		}
	}

	elapsed := time.Since(start)
	log.Printf("✅ Produced %d messages with (totalOriginalSize: %d bytes and totalCompressedSize:%d bytes) in %s (%.2f msg/sec)\n",
		n, totalOriginalSize, totalCompressedSize, elapsed, float64(n)/elapsed.Seconds())
}
