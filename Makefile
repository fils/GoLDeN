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
	podman build  --tag="fcore/golden:$(DOCKERVER)"  --file=./build/Dockerfile .

latest:
	podman build  --tag="fcore/golden:$(DOCKERVER)"  --file=./build/Dockerfile . ; \
	podman tag fcore/golden:$(DOCKERVER) fcore/golden:latest

push:
	podman push localhost/fcore/golden:$(DOCKERVER) fils/golden:$(DOCKERVER)
	podman push localhost/fcore/golden:$(DOCKERVER) fils/golden:latest

