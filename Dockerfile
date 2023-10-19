# Установка базового образа Golang
FROM golang:latest

# Установка переменной окружения GOPATH
ENV GOPATH /go

RUN mkdir pictures

# Копирование исходного кода в контейнер
COPY . /go/src/app

# Переход в директорию с исходным кодом
WORKDIR /go/src/app

# Сборка приложения
RUN go build -o main .

# Открытие порта 8080 для взаимодействия с приложением
EXPOSE 8081

# Запуск приложения при старте контейнера
CMD ["/app/main"]