{
  description = "Treman android build environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";

    flake-utils.url = "github:numtide/flake-utils";

    gio.url = "github:gioui/gio";
    gio.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils, gio }:
    flake-utils.lib.eachSystem
      [ "x86_64-linux" "x86_64-darwin" "aarch64-darwin" ]
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
            config = {
              permittedInsecurePackages = [
                "python-2.7.18.6"
              ];
            };
          };

          gio-shell = gio.devShells.${system}.default;

          gogio = pkgs.buildGoModule {
            inherit (gio-shell) ANDROID_SDK_ROOT JAVA_HOME nativeBuildInputs;
            name = "gogio";

            buildInputs = gio-shell.nativeBuildInputs;
            src = pkgs.fetchFromSourcehut {
              owner = "~eliasnaur";
              repo = "gio-cmd";
              rev = "0a86898b418418e80fba9c12e71dfdf764bb01d6";
              hash = "sha256-6e51+ZHflWaMiB0HPhVnRsLPbHV3zQBMLkpizu/rzLo=";
            };

            vendorHash = "sha256-2LQCFYyEletx+FswLV1Ui506qG62yHUKGr5vP5Y/b/s=";
            doCheck = false;
          };
        in
        {
          devShells.default = pkgs.mkShell
            {
              inherit (gio-shell) ANDROID_SDK_ROOT JAVA_HOME nativeBuildInputs;

              buildInputs = [ gogio ] ++ (with pkgs;
                [ go gum apksigner ]);

              shellHook = ''
                build () {
                  local V=$(gum input --prompt 'Version? ' --placeholder '7')

                  gum spin --spinner dot --title "Building..." -- \
                  gogio -target android -icon assets/meta/icon.png -version $V -minsdk 29 .

                  apksigner sign --ks ~/.android/sign.keystore treman.apk
                }

                stream () {
                  gum spin --spinner dot --title "Building..." -- \
                  gogio -target android -icon assets/meta/icon.png -minsdk 29 .
                  adb uninstall com.github.treman > /dev/null 2>&1
                  adb install treman.apk
                }
              '';
            };
        });
}
