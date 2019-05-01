DOCKERVER :=`cat VERSION`
.DEFAULT_GOAL := reciever

reciever:
	cd cmd/reciever ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o reciever

consumer:
	cd cmd/consumer ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o consumer 

sender:
	cd cmd/sender ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o sender 

docker:
	docker build  --tag="fcore/golden:$(DOCKERVER)"  --file=./build/Dockerfile . 

latest:
	docker build  --tag="fcore/golden:$(DOCKERVER)"  --file=./build/Dockerfile . ; \
	docker tag fcore/golden:$(DOCKERVER) fcore/golden:latest
