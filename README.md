# Beyond limits

### API:

- `POST /admin/login` - получение токена для доступа к админским методам

```json
{
  "login": "admin",
  "password": "secret123"
}
```

Ответ:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

Для доступа к защищенным методам:

```
Authorization: Bearer <token>
Content-Type: application/json
```

В системе предусмотрен только один админ, значения login и password хранятся в .env переменных.

- `GET /pictures` - получение списка картин с фильтрами и пагинацией

Параметры запроса:

| Параметр    | Тип    | Описание                                                |
|-------------|--------|---------------------------------------------------------|
| page        | number | Номер страницы (по умолчанию 1)                         |
| limit       | number | Картин на странице (по умолчанию 10)                    |
| genre       | number | Фильтр по ID жанра                                      |
| author      | number | Фильтр по ID автора                                     |
| minprice    | number | Минимальная цена                                        |
| maxprice    | number | Максимальная цена                                       |
| dimensions  | number | Фильтр по ID размера                                    |
| technique   | number | Фильтр по ID техники                                    |
| search      | string | Поиск по названию                                       |
| sort        | string | Сортировка (priceasc, pricedesc, dateasc, datedesc)     |

Ответ:

```json
{
  "data": [
    {
      "id": 1,
      "title": "Звёздная ночь",
      "price": 5000,
      "author": { "id": 1, "full_name": "Ван Гог" },
      "dimensions": { "id": 1, "width": 73, "height": 92 },
      "work_technique": { "id": 1, "name": "Масло" },
      "genre": { "id": 1, "name": "Пейзаж" },
      "photo": { "id": 1, "url": "/images/1.jpg", "mime": "image/jpeg" },
      "gallery": [
        { "id": 2, "url": "/images/1-1.jpg", "mime": "image/jpeg" }
      ],
      "created_at": "2024-01-15T10:00:00Z"
    }
  ],
  "meta": {
    "total_items": 100,
    "current_page": 1,
    "items_per_page": 10,
    "total_pages": 10,
    "next_page": 2,
    "prev_page": null
  }
}
```

- `GET /news` - cписок новостей с пагинацией

| Параметр    | Тип    | Описание                                                |
|-------------|--------|---------------------------------------------------------|
| page        | number | Номер страницы (по умолчанию 1)                         |
| limit       | number | Картин на странице (по умолчанию 5)                     |

Ответ:

```json
{
  "data": [
    {
      "id": 1,
      "title": "Новая выставка",
      "content": "Текст новости...",
      "created_at": "2024-02-20T12:00:00Z"
    }
  ],
  "meta": { ... } // Аналогично картинам
}
```

- Справочники (публичные)

`GET /genres`
`GET /authors`
`GET /dimensions`
`GET /work-techniques`

Ответ (пример для genres):

```json
[
  { "id": 1, "name": "Пейзаж" },
  { "id": 2, "name": "Портрет" }
]
```

- Админские методы

Картины

| Метод  | Путь                                      | Описание                           |
|--------|-------------------------------------------|------------------------------------|
| POST   | `/admin/pictures`                         | Добавление картины (без фото)      |
| PATCH  | `/admin/pictures/{id}`                    | Обновление данных                  |
| DELETE | `/admin/pictures/{id}`                    | Удаление                           |
| POST   | `/admin/pictures/{id}/photo`              | Загрузка основного фото            |
| POST   | `/admin/pictures/{id}/gallery`            | Добавление фото в галерею          |
| DELETE | `/admin/pictures/{id}/gallery/{photo-id}` | Удаление фото из галереи           |

Новости

| Метод  | Путь                   | Описание       |
|--------|------------------------|----------------|
| POST   | `/admin/news`          | Создание       |
| PATCH  | `/admin/news/{id}`     | Обновление     |
| DELETE | `/admin/news/{id}`     | Удаление       |

Справочники

| Метод  | Путь                  | Описание         |
|--------|-----------------------|------------------|
| POST   | `/admin/genres`       | Добавление жанра |
| DELETE | `/admin/genres/{id}`  | Удаление         |

(аналогично для authors, dimensions, work-techniques)
