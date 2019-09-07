CURRENT_TAG=$(shell git describe --abbrev=0 --tags)
TARGET:=$${HOME}/auth-sample-jwt

keys:
	@-mkdir ./src/keys
	@openssl genrsa 4096 > ./src/keys/private-key.pem
	@openssl rsa -in ./src/keys/private-key.pem -pubout -out ./src/keys/public-key.pem

deploy:
	dep ensure
	GOOS=linux go build  -a -tags netgo -installsuffix netgo -o ./docker/api/api ./src/main.go
	cd docker/api;docker build -t jwt-sample-api:latest .
	docker save jwt-sample-api:latest > api.tar
	-mkdir -p ${TARGET}
	mv api.tar ${TARGET}/
	git describe --abbrev=0 --tags > ./version_api
	-mkdir -p ${TARGET}/versions
	mv ./version_api ${TARGET}/versions/
	cp ./docker/docker-compose.yml ${TARGET}/
	-mkdir ${TARGET}/keys
	if [ ! -f ${TARGET}/keys/private-key.pem ]; then\
		cd ${TARGET}/keys; openssl genrsa 4096 > ./private-key.pem; openssl rsa -in ./private-key.pem -pubout -out ./public-key.pem;\
	fi

start:
	cd ${TARGET}; docker-compose up -d

stop:
	cd ${TARGET}; docker-compose down
