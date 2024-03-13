# CRUD service

### Getting started

Для начала нужно запустить 

`make build-app-image`

Затем выполнить:

`make start-test-service`

(В первый раз скорее всего будет ошибка, так как база будет не готова принимать соединение. Поэтому придется выполнить команду **еще раз**)

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
