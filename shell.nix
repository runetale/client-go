{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_17 goimports gopls protobuf protoc-gen-go
  ];

  shellHook = ''
    export GOPATH=$GOPATH
    PATH=$PATH:$GOPATH/bin
    export GO111MODULE=on
  '';
}
