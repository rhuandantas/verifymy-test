# verifymy-test challenge

## Requirements

- Run swag to generate api documentation
```sh
swag init
```
- Run this command to generate dependency injection and mock test files
```sh
go generate ./...
```

- set env vars
```sh
export AUTH_SECRET={your_secret}
export DB_USER_PASS={db_password}
export DB_USER_NAME={db_name}
```
- some application configurations can be set into ``resources/config.yml``
- to build database (myqsl) container run ``docker-compose up -d``
---
### to run application
```sh
go main .
```

### to run tests
```sh
ginkgo -v ./...
```

### to access swagger doc
```
http://localhost:3000/swagger/index.html
```

## endpoints

- before access ``/users`` endpoints you should get authentication token calling ``curl --request GET \
  --url http://localhost:3000/token`` and pass it through _Bearer Authentication_ or header['token']
- besides swagger doc you can also use cURL provided into ``resources/curls.json``