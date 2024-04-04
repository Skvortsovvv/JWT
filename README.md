## Первый запуск
1. ```go install github.com/pressly/goose/v3/cmd/goose@latest```
2. Создать БД c ```host: localhost port: 5432 user: postgres dbname: mydb``` 
3. Выполнить ```make migrate-up``` для создания таблицы ```users```