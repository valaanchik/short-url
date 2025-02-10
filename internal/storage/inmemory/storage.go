package inmemory

import (
	"errors"
	"sync"
)

type Storage struct {
	data map[string]string
	mut  sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Get(stringUrl, typeUrl string) (string, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	if typeUrl == "short" {
		longUrl, _ := s.data[stringUrl]
		return longUrl, nil
	} else {
		for shortUrl, longUrl := range s.data {
			if longUrl == stringUrl {
				return shortUrl, nil
			}
		}
	}
	return "", errors.New("url не найден")
}

func (s *Storage) Save(shortUrl, longUrl string) error {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.data[shortUrl] = longUrl
	return nil
}

func (s *Storage) Exist(stringUrl string) (bool, string) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	_, exist := s.data[stringUrl]
	if exist {
		return exist, "short"
	}
	for _, v := range s.data {
		if v == stringUrl {
			return true, "long"
		}
	}
	return false, ""
}
