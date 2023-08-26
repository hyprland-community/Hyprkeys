{
  buildGoModule,
  nix-filter,
  lib,
  ...
}:
buildGoModule {
  pname = "Hyprkeys";
  version = "1.0.2";
  src = nix-filter.lib {
    root = ./.;
    exclude = [
      ./README.md
      ./out.md
      ./out.json
      ./.github
    ];
  };

  ldflags = ["-s" "-w"];
  vendorHash = "sha256-JFvC9V0xS8SZSdLsOtpyTrFzXjYAOaPQaJHdcnJzK3s=";

  doCheck = true;

  preFixup = ''
    mkdir -p completions
    $out/bin/hyprkeys completions zsh > completions/hyprkeys.zsh
    $out/bin/hyprkeys completions bash > completions/hyprkeys.bash
    $out/bin/hyprkeys completions fish > completions/hyprkeys.fish

    installShellCompletion completions/*
  '';

  meta = with lib; {
    description = "A simple, scriptable keybind retrieval utility for Hyprland.";
    homepage = "https://github.com/hyprland-community/Hyprkeys";
    maintainers = with maintainers; [NotAShelf];
    mainProgram = "hyprkeys";
    platforms = ["x86_64-linux" "aarch64-linux"];
  };
}
