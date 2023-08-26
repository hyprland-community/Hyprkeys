{
  description = "Hyprkeys - a keybind inspection utility for Hyprland";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    nix-filter.url = "github:numtide/nix-filter";
  };

  outputs = {
    self,
    nixpkgs,
    nix-filter,
  }: let
    systems = ["x86_64-linux" "aarch64-linux"]; # we assume Hyprland only works on those two, for now
    forEachSystem = nixpkgs.lib.genAttrs systems;

    pkgsForEach = nixpkgs.legacyPackages;
  in {
    packages = forEachSystem (system: rec {
      hyprkeys = pkgsForEach.${system}.callPackage ./default.nix {inherit nix-filter;};
      default = hyprkeys;
    });

    devShells = forEachSystem (system: {
      default = pkgsForEach.${system}.mkShell {
        name = "hyprkeys-dev";
        packages = with pkgsForEach.${system}; [go];
      };
    });

    formatter = forEachSystem (system: nixpkgs.legacyPackages.${system}.alejandra);
  };
}
