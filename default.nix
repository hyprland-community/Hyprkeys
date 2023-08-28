{
  buildGoModule,
  nix-filter,
  lib,
  installShellFiles,
  ...
}:
buildGoModule rec {
  pname = "Hyprkeys";
  version = "1.0.3";
  src = nix-filter.lib {
    root = ./.;
    exclude =
      [
        ./README.md
        ./out.md
        ./out.json
        ./.github
      ]
      ++ lib.optional (!doCheck) [./test];
  };

  ldflags = ["-s" "-w" "-X main.version=v${version}"];
  vendorHash = "sha256-JFvC9V0xS8SZSdLsOtpyTrFzXjYAOaPQaJHdcnJzK3s=";

  doCheck = true;

  nativeBuildInputs = [installShellFiles];

  preFixup = ''
    mkdir -p completions
    $out/bin/hyprkeys completion zsh > completions/_hyprkeys
    $out/bin/hyprkeys completion bash > completions/hyprkeys.bash
    $out/bin/hyprkeys completion fish > completions/hyprkeys.fish

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
