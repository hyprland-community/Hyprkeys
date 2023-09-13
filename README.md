# ⌨️ Hyprkeys

A simple, fast and scriptable keybind inspection utility for [Hyprland](https://github.com/hyprwm/Hyprland)

## Installation & Usage

### From source

1. Download Go. You can obtain it from your distro's package manager. It is named "go" under most distros.
2. Clone this repository with `git clone https://github.com/hyprland-community/hyprkeys`
3. Install the application with `make build` then `sudo make install`
4. You can run the application with `hyprkeys`

### Using the Nix Flake

```nix
# flake.nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    hyprkeys = {
        url = "github:hyprland-community/hyprkeys";
        inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, hyprkeys }: let
  in {
    nixosConfigurations.yourHostName = nixpkgs.lib.nixosSystem {
      # ...

      system.packages = [ hyprkeys.packages.${system}.hyprkeys ];

      # ...
    };
  };
}
```
### Arch AUR Package

Arch User Repository [package](https://aur.archlinux.org/packages/hyprkeys) is maintained by [bloominstrong](https://github.com/bloominstrong)
```
git clone https://aur.archlinux.org/hyprkeys.git
cd hyprkeys
# Get needed dependencies, make and install hyprkeys
makepkg -si
```

### Installing Prebuild Binaries

Prebuilt binaries are distributed for each tagged release. You may find them under [releases](https://github.com/hyprland-community/Hyprkeys/releases)

## Usage

See `hyprkeys -v` for a full list of commands and flags.

## Project Demo

https://user-images.githubusercontent.com/86447830/211897915-778e9b24-061d-4d97-bc5e-444224610566.mp4

### Rofi script using hyprctl:

![OrCEzxZ - Imgur](https://user-images.githubusercontent.com/86447830/211898056-3bdb2f11-f7f5-4854-871f-4eabaa5b898f.png)

## Example Outputs (as of 11 JAN 2023)

**[MARKDOWN](test/out.md)**

**[JSON](test/out.json)**

**[RAW](test/out.txt)**

## Project Roadmap

See [TODO](todo.md).

## License

This project is licensed under the GPLv3 License. See the [LICENSE](LICENSE) file for more details.

## Contributing

Contributions, in form of Pull Requests and Issues are always welcome. If you wish to make changes, open a pull request and I will
guide you through it.
