dev:
	SERVER_ADDR=127.0.0.1 SERVER_PORT=3000 go run main.go

build:
	go build -o bin/godraft -ldflags "-X main.Version=$$VERSION" main.go

run:
	bin/godraft

dockerbuild:
	docker build -t zwindler/godraft:$$TAG --build-arg VERSION=$$TAG .

dockerpush: dockerbuild
	docker push zwindler/godraft:$$TAG

dockerrun:
	docker run -p 3000:3000 zwindler/godraft:$$TAG
