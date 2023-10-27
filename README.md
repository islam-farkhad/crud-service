# Homework 5

### Getting started

Для начала нужно запустить 

`make build-app-image`

Затем выполнить:

`make start-test-service`

(В первый раз будет ошибка, так как база будет не готова принимать соединение. Поэтому придется выполнить команду **еще раз**)

Теперь нужно накатить миграции:

`make test-migration-up`

Готово. Теперь можно запускать тесты:

`make run-integration-tests`

`make run-unit-tests`

### curl запросы:

- curl -X POST localhost:8080/post -d '{"content":"hello world!","likes":17}' -i
- curl -X POST localhost:8080/post/1/comment -d '{"content":"hello!"}' -i
- curl -X POST localhost:8080/post/1/comment -d '{"content":"hi!"}' -i
- curl -X GET localhost:8080/post/1 -i
- curl -X PUT localhost:8080/post -d '{"content":"hello universe!","likes":17,"id":1}' -i
- curl -X DELETE localhost:8080/post/1 -i


## Задание

* Отрефакторить код дз 3 недели(если необходимо). Применить подходит DI, переписать методы бд на использование интерфейсов

* Покрыть юнит тестами хендлеры из дз 3 . Минимальное покрытие - 40%.

* Покрыть интеграционными тестами хендлеры дз 3 недели. Минимальные тест кейсы - успешное выполнение, получение ошибки из-за передачи некорректных данных.

* Подготовить Makefile, в котором будут след команды: запуск тестового окружения при помощи docker-compose, запуск интеграционных тестов, запуск юнит тестов, запуск скрипта миграций, очищение базы от тествых данных