CURRENT_TAG=$(shell git describe --abbrev=0 --tags)
TARGET=${HOME}/auth-sample-jwt/

keys:
	-mkdir ./src/keys
	openssl genrsa 4096 > ./src/keys/private-key.pem
	openssl rsa -in ./src/keys/private-key.pem -pubout -out ./src/keys/public-key.pem

test:
	go test ./src/...

deploy:
	GOOS=linux go build  -a -tags netgo -installsuffix netgo -o ./docker/api/api ./src/main.go
	cd docker/api;docker build -t jwt-sample-api:latest .
	docker save api:latest > api.tar
	-mkdir -p ${TARGET}
	mv api.tar ${TARGET}/
	git describe --abbrev=0 --tags > ./version_api
	-mkdir -p ${TARGET}/versions
	mv ./version_api ${TARGET}/versions/
	cp ./docker/docker-compose.yml ${TARGET}/
	-mkdir ./src/keys
	openssl genrsa 4096 > ${TARGET}/keys/private-key.pem
	openssl rsa -in ${TARGET}/keys/private-key.pem -pubout -out ${TARGET}/keys/public-key.pem

start:
	cd ${TARGET}; docker-compose up -d
