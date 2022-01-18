{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  # TODO: (shintard) support aarch64 package
  outputs = { self, nixpkgs, flake-utils, ... }:
    let systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
     in flake-utils.lib.eachSystem systems (system:
      let
        buildDeps = with nixpkgs; [ git go ];
        devDeps = with nixpkgs;
          buildDeps ++ [
            protoc-gen-go
            go_1_17
            goimports
            gopls
            protobuf
            protoc-gen-go-grpc
          ];
      in
      { devShell = nixpkgs.mkShell {
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
