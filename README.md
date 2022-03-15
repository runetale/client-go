# wissy

## for install

## for build

## for development

## for NixOS
if you want to run it in a daemon, add the following.

```
  systemd.services.wissy = {
    description = "wissy daemon";
    wants  = [ "network-online.target" "systemd-networkd-wait-online.service"];
    after = [ "network-online.target" ];
    serviceConfig = {
      User = "root";
      Type = "simple";
      ExecStart = "/bin/wissy up";
      Restart = "on-failure";
      RestartSec = "15";
    };
    wantedBy = [ "multi-user.target" ];
  };
  
  systemd.services.wissy.enable = true;
```

TODO: flake.nix, etc. to add daemon service directly
