# Kafka Tuning with Golang

This repository contains experimental code for testing **Kafka performance tuning and scaling** using **Golang**.  
It supports the article:  
📖 [Golang x Kafka #4: Performance Tuning and Scaling](https://medium.com/@andriantriputra/golang-x-kafka-4-performance-tuning-and-scaling-39292e062bdc)

---

## 🧩 Overview

The goal of this project is to explore how different configurations affect **throughput**, **latency**, and **scalability** in Kafka when integrated with Go.  
It includes experiments for:
- Producer tuning (batching, async writes)
- Consumer tuning (fetch size, concurrency)
- Compression testing (Snappy)
- Scaling with partitions, consumers, and brokers

---

## 🚀 Getting Started

### Prerequisites
- Go 1.20+
- Running Kafka cluster (local or via Docker)
- Basic knowledge of Kafka topics and partitions

### Installation
```bash
git clone https://github.com/andriantp/kafka-tuning.git
cd kafka-tuning
go mod tidy
```

### Run Experiments

Update broker and topic in main.go

Run different modes:
``` bash
# Basic producer/consumer
go run main.go produce1
go run main.go consume1

# Producer tuning
go run main.go produce2

# Consumer tuning
go run main.go consume2

# Scaling test (multiple consumers)
go run main.go consume3 <groupName>
```

### ⚙️ Experiments

| Category    | Focus                   | Examples                  |
| ----------- | ----------------------- | ------------------------- |
| Producer    | Batching, async mode    | `produce1`, `produce2`    |
| Consumer    | Fetch size, concurrency | `consume1`, `consume2`    |
| Scaling     | Multi-partition & group | `consume3`                |
| Compression | Snappy codec test       | Enable in producer config |


### 📈 Key Findings

- One-by-one messages = very low throughput
- Batch & async producer = 5–10x performance boost
- Larger fetch size improves consumer throughput
- Snappy gives a good balance between compression and latency
- Multi-partition scaling distributes load effectively
Detailed analysis is available in the Medium article linked above.

## 🧠 Next Steps

Ideas for future exploration:
- Compare with other codecs (LZ4, ZSTD)
- Add metrics visualization (Prometheus/Grafana)
- Test multi-broker Kafka setup
- Benchmark against Python/Java clients

### 🧾 License

This project is licensed under the Apache 2.0 License.

See the LICENSE file for details.

