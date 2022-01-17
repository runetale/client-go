{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        buildDeps = with pkgs; [ git go ];
        devDeps = with pkgs;
          buildDeps ++ [
            protoc-gen-go
            go_1_17
            goimports
            gopls
            protobuf
            protoc-gen-go-grpc
          ];
      in
      { devShell = pkgs.mkShell {
        buildInputs = devDeps; 

        shellHook = ''
          export GOPATH=$GOPATH
          PATH=$PATH:$GOPATH/bin
          export GO111MODULE=on
        '';
      };
    }
  );
}
