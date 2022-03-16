# wissy

## for install

## for build

## for development

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
