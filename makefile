.PHONY: build run clean create

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

create:
	docker exec -it postgresdata-db-1 bash -c "psql -U Misfio32 -c '\c postgres' -c '\i ./mnt/createTable.sql' -c '\i ./mnt/limitation.sql'"

delete:
	docker exec -it postgresdata-db-1 bash -c "psql -U Misfio32 -c '\c postgres' -c '\i ./mnt/dropTable.sql'"
