package serv

import (
	"context"
	"errors"
	"log"
	"test-shorturl-ozon/internal/storage"

	//"internal/serv"
	pb "test-shorturl-ozon/proto"
)

type ShortUrlServ struct {
	storage storage.Storage
	pb.UnimplementedShortenerServiceServer
}

func NewShortUrlServ(storage storage.Storage) *ShortUrlServ {
	return &ShortUrlServ{storage: storage}
}

func (s *ShortUrlServ) Post(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	if req.LongUrl == "" {
		return nil, errors.New("Пустой запрос")
	}

	exists, urlType := s.storage.Exist(req.LongUrl)
	if exists {
		shortUrl, _ := s.storage.Get(req.LongUrl, urlType)
		return &pb.PostResponse{ShortUrl: shortUrl}, nil
	}
	shortUrl := GenerateUrl(req.LongUrl)
	err := s.storage.Save(shortUrl, req.LongUrl)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{ShortUrl: shortUrl}, nil
}

func (s *ShortUrlServ) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.ShortUrl == "" {
		return nil, errors.New("Пустой запрос")
	}

	exists, urlType := s.storage.Exist(req.ShortUrl)
	if !exists {
		return nil, errors.New("Url не найден")
	}

	LongUrl, err := s.storage.Get(req.ShortUrl, urlType)
	if err != nil {
		log.Printf("Ошибка при получении URL: %v\n", err)
		return nil, err
	}

	return &pb.GetResponse{LongUrl: LongUrl}, nil
}
