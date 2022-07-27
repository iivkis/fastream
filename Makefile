server-run:
	go run ./cmd/app/main.go

server-build:
	go build ./cmd/app/main.go

front-run:
	cd ./web && npm run dev

front-build:
	cd ./web && npm run build

full-build: 
	make server-build
	make front-build