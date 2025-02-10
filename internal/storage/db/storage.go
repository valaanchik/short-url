package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Save(shortUrl, longUrl string) error {
	_, err := s.db.Exec("INSERT INTO url (shortUrl, longUrl) VALUES ($1, $2)", shortUrl, longUrl)
	if err != nil {
		log.Printf("Ошибка при записи в БД: %v\n", err)
		return err
	}
	return nil
}

func (s *Storage) Exist(urlString string) (bool, string) {
	var shortUrlExists, longUrlExists bool

	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM url WHERE shortUrl = $1), EXISTS(SELECT 1 FROM url WHERE longUrl = $1)", urlString, urlString).
		Scan(&shortUrlExists, &longUrlExists)
	if err != nil {
		log.Printf("Ошибка при проверке существования URL: %v\n", err)
		return false, ""
	}

	if shortUrlExists {
		return true, "short"
	}
	if longUrlExists {
		return true, "long"
	}
	return false, ""
}

func (s *Storage) Get(urlString, typeUrl string) (string, error) {
	var resultUrl string
	if typeUrl == "short" {
		err := s.db.QueryRow("SELECT longUrl FROM url WHERE shortUrl = $1", urlString).Scan(&resultUrl)
		if err != nil {
			log.Printf("Ошибка при извлечении из базы данных: %v\n", err)
			return "", err
		} else {
			return resultUrl, nil
		}
	} else {
		err := s.db.QueryRow("SELECT shortUrl FROM url WHERE longUrl = $1", urlString).Scan(&resultUrl)
		if err != nil {
			log.Printf("Ошибка при извлечении из базы данных: %v\n", err)
			return "", err
		} else {
			return resultUrl, nil
		}
	}

	//return resultUrl, nil
}
