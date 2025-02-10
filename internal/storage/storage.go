package storage

type Storage interface {
	Exist(stringUrl string) (bool, string)
	Get(stringUrl, typeUrl string) (string, error)
	Save(shortUrl, longUrl string) error
}
