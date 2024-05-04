.PHONY: build run clean run-psql

build:
	@echo "Сборка сервера..."
	docker build -t my-golang-server .

run:
	docker-compose -f ./postgresData/docker-compose.yaml up -d 

clean:
	@echo "Очистка..."
	docker stop my-server || true
	docker-compose -f ./postgresData/docker-compose.yaml down --remove-orphans
	docker rmi my-golang-server || true
