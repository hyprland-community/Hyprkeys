{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ../flake.lock)).nodes) nixpkgs gomod2nix;
    in
      import (fetchTree nixpkgs.locked) {
        overlays = [
          (import "${fetchTree gomod2nix.locked}/overlay.nix")
        ];
      }
  ),
  lib,
}:
pkgs.buildGoApplication {
  pname = "Hyprkeys";
  version = "1.0.1";
  pwd = ../.;
  src = ../.;
  modules = ./gomod2nix.toml;

  meta = with lib; {
    mainProgram = "hyprkeys";
    license = licenses.mit;
    description = "A simple, scriptable keybind retrieval utility for Hyprland";
    homepage = "https://github.com/hyprland-community/${pname}";
    changelog = "https://github.com/hyprland-community/${pname}/releases/tag/v${version}";
  };
}
