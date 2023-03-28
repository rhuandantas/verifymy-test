run: set-vars
	swag init
	go generate ./...
	go run .

run-test:
	ginkgo -v -p ./...

set-vars:
	export DB_USER_PASS=12345678
	export DB_USER_NAME=root
	export AUTH_SECRET=Testando
