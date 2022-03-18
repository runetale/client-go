BINDIR:=bin

# TODO: (shintard) it will disappear once flake is supported
.PHONY: store
store:
	go build -o $(BINDIR)/wissy cmd/wissy/main.go
	nix-store --add $(BINDIR)/wissy

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

## build the wissy
build-wissy:
	docker-compose -f docker-compose.yml build

up-wissy:
	docker-compose -f docker-compose.yml up -d
