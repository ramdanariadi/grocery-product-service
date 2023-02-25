go-run:
	./grocery-product-service DB_USERNAME=postgres DB_PASS=secret DB_NAME=grocery-product-service DB_HOST=localhost REDIS_HOST=localhost REDIS_PORT=6379

go-run-dev:
	DB_USERNAME=postgres DB_PASS=secret DB_NAME=grocery-product-service DB_HOST=localhost REDIS_HOST=localhost REDIS_PORT=6379 go run main.go