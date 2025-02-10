package main

import (
	"log"
	"net"
	"os"

	"test-shorturl-ozon/internal/serv"
	"test-shorturl-ozon/internal/storage"
	"test-shorturl-ozon/internal/storage/db"
	"test-shorturl-ozon/internal/storage/inmemory"
	pb "test-shorturl-ozon/proto"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	var storageType = os.Getenv("STORAGE_TYPE")
	var store storage.Storage

	if storageType == "db" {
		dbConn, err := db.ConnectDB()
		if err != nil {
			log.Fatalf("Ошибка подключения к БД: %v", err)
		}
		store = db.NewStorage(dbConn)
	} else {
		store = inmemory.NewStorage()
	}

	grpcServer := grpc.NewServer()
	shortenerService := serv.NewShortUrlServ(store)

	pb.RegisterShortenerServiceServer(grpcServer, shortenerService)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	log.Println("gRPC-сервер запущен на порту 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске gRPC: %v", err)
	}
}
