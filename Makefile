# билд и запуск приложения
build_run:
	docker compose -f "docker-compose.yml" up -d --build  && migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5436/postgres?sslmode=disable' up

# Остановка и удаление контейнера
stop:
	docker compose -f "docker-compose.yml" down
