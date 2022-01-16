{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_17 goimports gopls protobuf protoc-gen-go
  ];
}
