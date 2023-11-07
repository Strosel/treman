{
  description = "Build a dioxus project";
  # This flake is an emalgam of the trunk template from crane as well as the source of buildTrunkPackage

  inputs = {
    # Pinned version of nixpkgs
    # Has correct version of both wasm-bindgen-cli and dioxus-cli
    nixpkgs.url = "github:NixOS/nixpkgs/85f1ba3e51676fa8cc604a3d863d729026a6b8eb";

    crane = {
      url = "github:ipetkov/crane";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    flake-utils.url = "github:numtide/flake-utils";

    rust-overlay = {
      url = "github:oxalica/rust-overlay";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };

  outputs = { self, nixpkgs, crane, flake-utils, rust-overlay, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ (import rust-overlay) ];
        };

        inherit (pkgs) lib;

        rustToolchain = pkgs.rust-bin.stable.latest.default.override {
          # Set the build targets supported by the toolchain,
          # wasm32-unknown-unknown is required for trunk
          targets = [ "wasm32-unknown-unknown" ];
        };
        craneLib = (crane.mkLib pkgs).overrideToolchain rustToolchain;

        # When filtering sources, we want to allow assets other than .rs files
        src = lib.cleanSourceWith {
          src = ./.; # The original, unfiltered source
          filter = path: type:
            (lib.hasSuffix "\.html" path) ||
            (lib.hasSuffix "\.css" path) ||
            (lib.hasSuffix "\.js" path) ||
            # Example of a folder for images, icons, etc
            (lib.hasInfix "assets/" path) ||
            (lib.hasInfix "public/" path) ||
            # Default filter from crane (allow .rs files)
            (craneLib.filterCargoSources path type)
          ;
        };

        # Common arguments can be set here to avoid repeating them later
        commonArgs = {
          inherit src;
          strictDeps = true;
          # We must force the target, otherwise cargo will attempt to use your native target
          CARGO_BUILD_TARGET = "wasm32-unknown-unknown";

          buildInputs = [
            # Add additional build inputs here
          ] ++ lib.optionals pkgs.stdenv.isDarwin [
            # Additional darwin specific inputs can be set here
            pkgs.libiconv
          ];
        };

        # Build *just* the cargo dependencies, so we can reuse
        # all of that work (e.g. via cachix) when running in CI
        cargoArtifacts = craneLib.buildDepsOnly (commonArgs // {
          # You cannot run cargo test on a wasm build
          doCheck = false;
        });

        # Build the actual crate itself, reusing the dependency
        # artifacts from above.
        # This derivation is a directory you can put on a webserver.
        inherit (pkgs) binaryen tailwindcss;
        my-app = craneLib.mkCargoDerivation (commonArgs // {
          inherit cargoArtifacts;

          # Force dioxus to not download dependencies, but use the provided version 
          preConfigure = ''
            HOME=$(pwd)
            XDG_DATA_HOME=$HOME/.local/share

            mkdir -p $HOME/.local/share/dioxus \
              $HOME/Library/Application\ Support

            ln -sv $HOME/.local/share/dioxus \
              $HOME/Library/Application\ Support

            ln -s ${binaryen} $HOME/.local/share/dioxus/binaryen
            ln -s ${tailwindcss} $HOME/.local/share/dioxus/tailwindcss
          '';

          buildPhaseCargoCommand = ''
            local profileArgs=""
            if [[ "$CARGO_PROFILE" == "release" ]]; then
              profileArgs="--release"
            fi

            tailwindcss -i ./input.css -o ./public/tailwind.css
            dx build $profileArgs
          '';

          installPhaseCommand = ''
            cp -r dist $out
          '';

          # Installing artifacts on a distributable dir does not make much sense
          doInstallCargoArtifacts = false;

          nativeBuildInputs = [
            pkgs.dioxus-cli
            pkgs.wasm-bindgen-cli
            binaryen
            pkgs.nodejs
            tailwindcss
          ];
        });
      in
      {
        checks = {
          # Build the crate as part of `nix flake check` for convenience
          inherit my-app;

          # Run clippy (and deny all warnings) on the crate source,
          # again, reusing the dependency artifacts from above.
          #
          # Note that this is done as a separate derivation so that
          # we can block the CI if there are issues here, but not
          # prevent downstream consumers from building our crate by itself.
          my-app-clippy = craneLib.cargoClippy (commonArgs // {
            inherit cargoArtifacts;
            cargoClippyExtraArgs = "--all-targets -- --deny warnings";
          });

          # Check formatting
          my-app-fmt = craneLib.cargoFmt {
            inherit src;
          };
        };

        packages.default = my-app;
      });
}
