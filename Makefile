BINDIR:=bin

# TODO: (shintard) it will disappear once flake is supported
.PHONY: store
store:
	go build -o $(BINDIR)/dotshake cmd/dotshake/main.go
	nix-store --add $(BINDIR)/dotshake

# docker
## build the server
build-server:
	docker-compose -f docker/docker-compose.server.yml build

up-server:
	docker-compose -f docker/docker-compose.server.yml up -d

down-server:
	docker-compose -f docker/docker-compose.server.yml down

## build the signaling server
build-signaling:
	docker-compose -f docker/docker-compose.signaling.yml build

up-signaling:
	docker-compose -f docker/docker-compose.signaling.yml up -d

down-signaling:
	docker-compose -f docker/docker-compose.server.yml down

## build the dotshake
build-dotshake:
	docker-compose -f docker-compose.yml build

up-dotshake:
	docker-compose -f docker-compose.yml up -d
