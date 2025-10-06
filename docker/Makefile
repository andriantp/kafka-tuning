# ===============================
# Makefile Kafka Zookeeper 
# ===============================

# Bersihkan docker system & volume
clean:
	@echo "📋 Cleaning Docker system & volumes..."
	docker system prune -a -f
	docker volume prune -f
	rm -rf ./data/*
	mkdir -p ./data/kafka1 ./data/kafka2 ./data/kafka3
	chmod 777 ./data/*

# Deploy (up)
up: clean
	@echo "🚀 Starting Kafka + AKHQ..."
	docker compose -f docker-compose.yml up --force-recreate -d --build

# Stop (down)
down:
	@echo "✅ Stopping Kafka  ..."
	docker compose -f docker-compose.yml down -v
	$(MAKE) clean

# Status
status:
	docker compose -f docker-compose.yml ps

# Logs 
logs:
	docker compose -f docker-compose.yml logs -f
