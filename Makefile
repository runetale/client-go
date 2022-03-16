BINDIR:=bin

.PHONY: store
store:
	go build -o $(BINDIR)/wissy cmd/wissy/main.go
	nix-store --add $(BINDIR)/wissy
