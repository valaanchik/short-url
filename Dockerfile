FROM golang:1.22.2-alpine

WORKDIR /app

COPY . .

RUN go mod download

#  загрузки grpcurl
RUN apk add --no-cache curl bash

# Скачиваем grpcurl 
RUN curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.8.5/grpcurl_1.8.5_linux_x86_64.tar.gz -o grpcurl.tar.gz && \
    tar -xvzf grpcurl.tar.gz && \
    mv grpcurl /usr/local/bin/ && \
    rm -f grpcurl.tar.gz

RUN go build -o main ./run

CMD ["./main"]
