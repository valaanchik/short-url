package test

import (
	"context"
	"errors"
	"test-shorturl-ozon/internal/serv"
	pb "test-shorturl-ozon/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	data map[string]string
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		data: make(map[string]string),
	}
}

func (m *MockStorage) Save(shortUrl, longUrl string) error {
	m.data[shortUrl] = longUrl
	return nil
}

func (m *MockStorage) Exist(urlString string) (bool, string) {
	for short, long := range m.data {
		if short == urlString {
			return true, "short"
		}
		if long == urlString {
			return true, "long"
		}
	}
	return false, ""
}

func (m *MockStorage) Get(urlString, typeUrl string) (string, error) {
	for short, long := range m.data {
		if typeUrl == "short" && short == urlString {
			return long, nil
		}
		if typeUrl == "long" && long == urlString {
			return short, nil
		}
	}
	return "", errors.New("url не найден")
}

func TestPost(t *testing.T) {
	mockStorage := NewMockStorage()
	service := serv.NewShortUrlServ(mockStorage)

	req := &pb.PostRequest{LongUrl: "https://google.com"}
	res, err := service.Post(context.Background(), req)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.ShortUrl)

	exists, urlType := mockStorage.Exist(res.ShortUrl)
	assert.True(t, exists)
	assert.Equal(t, "short", urlType)
}

func TestGet(t *testing.T) {
	mockStorage := NewMockStorage()
	service := serv.NewShortUrlServ(mockStorage)

	mockStorage.Save("short123456", "https://google.com")

	req := &pb.GetRequest{ShortUrl: "short123456"}
	res, err := service.Get(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "https://google.com", res.LongUrl)
}
