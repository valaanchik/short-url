package test

import (
	"strings"
	"test-shorturl-ozon/internal/serv"
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	t.Run("Normal Case", func(t *testing.T) {
		result := serv.GenerateShortUrl(1)
		if len(result) != 10 {
			t.Fatalf("Длина короткого URL должна быть 10, но получена: %d", len(result))
		}
		if !strings.ContainsAny(result, "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890_") {
			t.Fatalf("Короткий URL содержит недопустимые символы: %v", result)
		}
	})

}

func TestGenerateUrl(t *testing.T) {
	t.Run("Valid URL", func(t *testing.T) {
		longUrl := "http://example.com"
		result := serv.GenerateUrl(longUrl)
		if len(result) != 10 {
			t.Fatalf("Длина короткого URL должна быть 10, но получена: %d", len(result))
		}
		if !strings.ContainsAny(result, "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890_") {
			t.Fatalf("Короткий URL содержит недопустимые символы: %v", result)
		}
	})
}
