ALL: bin/server

bin/server: cmd/server/main.go process.go $(shell find config -type f)
	go build -o $@ ./cmd/server