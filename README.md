# Приложение для мониторинга Docker

Это приложение отслеживает контейнеры Docker, проверяя их с регулярными интервалами и сохраняя данные в базе данных PostgreSQL. Затем данные отображаются на динамически создаваемой веб-странице.

## Сервисы

1. **Серверная часть**: Предоставляет RESTful API для запроса и добавления данных в базу данных.
2. **Интерфейс**: Отображает данные в табличном формате с помощью React.
3. **Database**: база данных PostgreSQL для хранения данных о состоянии контейнеров.
4. **Pinger**: Пингует контейнеры Docker и отправляет данные в базу данных через внутренний API.

## Запуск приложения

1. Клонируйте репозиторий.
2. Перейдите в каталог проекта.
3. Выполните следующую команду:

``bash
- docker-compose up --build
- http://localhost:8090/api/containers

