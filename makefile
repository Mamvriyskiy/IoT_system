.PHONY: build run clean run-psql

build: run-psql
	@echo "Сборка сервера..."
	docker build -t my-golang-server .

run:
	docker-compose -f ./postgresData/docker-compose.yaml start
	@echo "Запуск сервера..."
	docker run -d --name=my-server -p 80:8080 my-golang-server

run-psql:
	@echo "Создание контейнера PostgreSQL..."
	docker-compose -f ./postgresData/docker-compose.yaml create

clean:
	@echo "Очистка..."
	docker stop my-server || true
	docker-compose -f ./postgresData/docker-compose.yaml down --remove-orphans
	docker rmi my-golang-server || true
