FROM golang:1.22.2-alpine

WORKDIR /app

# Копируем все файлы проекта
COPY . .

# Устанавливаем зависимости Go
RUN go mod download

# Устанавливаем curl для загрузки grpcurl
RUN apk add --no-cache curl bash

# Скачиваем grpcurl с официального репозитория GitHub
RUN curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.8.5/grpcurl_1.8.5_linux_x86_64.tar.gz -o grpcurl.tar.gz && \
    tar -xvzf grpcurl.tar.gz && \
    mv grpcurl /usr/local/bin/ && \
    rm -f grpcurl.tar.gz

# Сборка приложения
RUN go build -o main ./run

CMD ["./main"]
