{
  description = "Rust development nix flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        overlays = [ 
        ];
        pkgs = import nixpkgs { inherit system overlays; };
     in {
        stdenv = pkgs.fastStdenv;
        devShell = pkgs.mkShell {
          LIBCLANG_PATH = pkgs.libclang.lib + "/lib/";

          nativeBuildInputs = with pkgs; [
            bashInteractive
            clang
            go
            gopls
            cmake
            llvmPackages_11.bintools
            llvmPackages_11.libclang
          ];
          buildInputs = with pkgs; [
          ];

        };
  });
}
