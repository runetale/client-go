# wissy


## env
TODO: (shintard) preparing enviroment variables

## for install
TODO: (shintard) preparing how install for wissy command
### linux

### darwin

### freebsd

### windows

## for build without docker
TODO: (shintard) preparing arguments for each commands
check the command line arguments carefully!

### server
start the server without docker
`sudo go run cmd/server/server.go`

### signaling
start the signaling without docker
`go run cmd/signaling/signaling.go`

### wissy
start the wissy without docker
`sudo go run cmd/wissy/up.go --key <your setup key>`

## for build with docker
start the server using docker-compose
`make up-server`
if you want to build
`make build-server`

start the signaling server using docker-compose
`make up-signaling`
if you want to build
`make build-signaling`

start the wissy using docker-compose
`make up-wissy`
if you want to build
`make build-wissy`

## for development
if you want develop server
`cmd/server` is the server that manages the peer's information.

if you want develop signaling server
`cmd/signaling` is the server to negotiate peer.

if you want develop wissy
`cmd/wissy` is the connects to peer clients, signaling servers and servers, and performs peer communication.

## for NixOS
TODO: Instead of using make, use flake to create the wissy binary. This will make development much easier since the build binaries can be placed directly in the store without using the nix-store command.

if you want to run it in a daemon,
1. run `make store` to create a wissy build binary in store
2. add the store path as follows and rebuild nixos

```
  { pkgs, ... }:

  let
  in {
    nixpkgs.overlays = [(self: super: {
      wissy = pkgs.writeScriptBin "wissy" ''
        #! ${pkgs.stdenv.shell} -e
        exec </your store binary> up // place your own built binaries in the store
      '';
    })];
  
    # for development wissy
    systemd.services.wissy = {
      description = "wissy daemon";
      wants  = [ "network-online.target" "systemd-networkd-wait-online.service"];
      after = [ "network-online.target" ];
      path = [ pkgs.iproute ];
      serviceConfig = {
        User = "root";
        Type = "simple";
        ExecStart = "${pkgs.wissy}/bin/wissy";
        Restart = "on-failure";
        RestartSec = "15";
      };
      wantedBy = [ "multi-user.target" ];
    };
  }

```

enjoy the daemons.
