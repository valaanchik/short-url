# ShortURL 
Это сервис для сокращения URL, который использует gRPC для взаимодействия. Он предоставляет API для создания коротких ссылок и получения оригинальных URL-адресов по коротким ссылкам. 
Реализованы: 
  Метод Post, который сохраняет оригинальный URL в базе и возвращает сокращённый.
  Метод Get, который принимает сокращённый URL и возвращает оригинальный.

## Установка
Сборка Docker-образа и запуск через Docker Compose
Для запуска сервиса и базы данных используйте Docker Compose:
  docker-compose up --build
### Описание переменных окружения
STORAGE_TYPE: Выбор типа хранилища. Два доступных варианта:
  postgresql — для использования PostgreSQL.
  inmemory — для использования In-Memory хранилища.
DB_HOST: Адрес хоста 
DB_PORT: Порт для подключения к БД
DB_USER: Имя пользователя для подключения к БД
DB_PASSWORD: Пароль для подключения к БД.
DB_NAME: Название БД
DB_SSLMODE: Режим SSL для подключения к базе данных PostgreSQL. Установите в disable для локальной разработки.
### Проверка работы сервиса через gRPC
Сокращение URL
Для сокращения длинного URL, используйте команду grpcurl:
    grpcurl -plaintext -d '{"longUrl": "https://example.com"}' localhost:50051 Short_url.ShortenerService/Post
 
    ![ShortenerService](https://github.com/user-attachments/assets/8e49dd43-228b-4941-b958-1ae107bb1ff4)
  Для получения длинного URL, по сокращенному используйте команду grpcurl:
    grpcurl -v -plaintext -d '{"shortUrl": "GudobAAAAA"}' localhost:50051 short_url.ShortenerService/Get
    ![ShortenerService](https://github.com/user-attachments/assets/51b43092-4a96-440f-8c21-0c8c731fd9c2)


    
