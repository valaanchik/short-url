package serv

import (
	"math"
	"strings"
)

const alp = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890_"
const base = 63
const size = 10

func GenerateShortUrl(id uint64) string {
	if id == 0 {
		return strings.Repeat(string(alp[0]), size)
	}

	var shortUrl strings.Builder
	for id > 0 && shortUrl.Len() < size {
		newr := id % base
		shortUrl.WriteByte(alp[newr])
		id /= base
	}
	for shortUrl.Len() < size {
		shortUrl.WriteByte(alp[0])
	}

	return shortUrl.String()
}

func GenerateUrl(longUrl string) string {
	hash := hash(longUrl)
	return GenerateShortUrl(hash)
}

func hash(s string) uint64 {
	var hash uint64
	for i, char := range s {
		hash += uint64(char) * uint64(math.Pow(2, float64(i%22)))
	}
	return hash
}
