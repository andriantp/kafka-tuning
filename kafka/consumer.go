package kafka

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	k "github.com/segmentio/kafka-go"
)

func Consume1(brokers []string, topic string, group string, expect int) {
	r := k.NewReader(k.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: group,
	})
	defer r.Close()
	fmt.Println("🚀 Consumer started...")
	timeout := time.After(20 * time.Minute) // ⏰ stop kalau gak nyampe juga
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	real := 0
	start := time.Now()
	for {
		select {
		case <-timeout:
			log.Printf("⏰ Timeout reached. Got %d/%d messages (%.2f%%)",
				real, expect, float64(real)/float64(expect)*100)
			return

		default:
			_, err := r.ReadMessage(ctx)
			if err != nil {
				log.Fatal("failed to read:", err)
			}
			real++
			// log.Printf("[%d] 📩 partition=%d offset=%d value=%s\n",
			// 	real, m.Partition, m.Offset, string(m.Value))

			// progress log
			// if real%100 == 0 || real == expect {
			// 	percentage := float64(real) / float64(expect) * 100
			// 	log.Printf("✅ Progress: %.2f%% (%d/%d)",
			// 		percentage, real, expect)
			// }

			// selesai
			if real >= expect {
				elapsed := time.Since(start)
				log.Printf("🎯 Done. Total=%d messages Duration:%.4f second", real, elapsed.Seconds())
				return
			}
		}
	}
}

func Consume2(brokers []string, topic string, group string, expect int) {
	r := k.NewReader(k.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  group,
		MinBytes: 1,
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()
	fmt.Println("🚀 Consumer started...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute) // ⏰ timeout
	defer cancel()

	real := 0
	start := time.Now()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Printf("⏰ Timeout reached. Got %d/%d messages (%.2f%%)",
					real, expect, float64(real)/float64(expect)*100)
				return
			}
			log.Fatal("failed to read:", err)
		}

		real++

		// progress log
		if real%100 == 0 || real == expect {
			percentage := float64(real) / float64(expect) * 100
			log.Printf("✅ Progress: %.2f%% (%d/%d)",
				percentage, real, expect)
		}

		// proses tiap pesan di goroutine biar non-blocking
		go func(m k.Message, idx int) {
			log.Printf("[%d] 📩 partition=%d offset=%d value=%s\n",
				idx, m.Partition, m.Offset, string(m.Value))
		}(m, real)

		// selesai
		if real >= expect {
			elapsed := time.Since(start)
			log.Printf("🎯 Done. Total=%d messages Duration:%.4f second", real, elapsed.Seconds())
			return
		}
	}
}

func Consume3(brokers []string, topic string, group string, expect int) {
	r := k.NewReader(k.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  group,
		MinBytes: 1,
		// MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()
	fmt.Println("🚀 Consumer started...")
	real := 0
	start := time.Now()
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read:", err)
		}
		real++
		log.Printf("[%d] 📩 partition=%d offset=%d value=%s\n",
			real, m.Partition, m.Offset, string(m.Value))

		// progress log
		percentage := float64(real) / float64(expect) * 100
		elapsed := time.Since(start)
		log.Printf("✅ Progress: %.2f%% (%d/%d) Duration:%.4f second",
			percentage, real, expect, elapsed.Seconds())

	}
}
