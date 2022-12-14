# REST API simple-app

### Использую следующие концепции:
<ul>
  <li>REST API</li>
  <li>Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>. </li>
  <li>Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.</li>
  <li>Работа с БД Postgres. Запуск из Docker(multi-stage builds). Генерация файлов миграций.</li>
  <li>Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>. Работа с переменными окружения.</li>
  <li>Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.</li>
  <li>Организация обработки транзакций в Service Layer.</li>
  <li>Регистрация и аутентификация. Работа с JWT. Middleware.</li>
  <li>Написание SQL запросов.</li>
  <li>Graceful Shutdown</li>
</ul>

### Для запуска приложения:
```make build && make run```
