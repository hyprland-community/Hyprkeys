# Hyprkeys
A simple application to quickly get your Hyprand keybinds

For now the software is in really early alpha quality. It is not recommended to use it yet.
Not like it can break anything, but you'll have wasted your time

In case you want to use it, instructions are below:

## Installation & Usage
1. Download Go. You can find it [here](https://golang.org/dl/)
2. Clone this repository with `git clone https://github.com/notashelf/hyprkeys`
3. Build the application with `go build` and run it with `./hyprkeys`

Alternatively, open this directory and run `go run .` to run without compiling.

## TODO
- [ ] Format keybinds better, maybe with a proper table
- [ ] Add a way to change keybinds with Hyprkeys after reading them

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contributing
If you want to contribute, feel free to open a pull request. I'll try to review it as soon as possible.

## Example Output (as of 2022-10-17)

| Mod | Key | Dispatcher | Command |
| --- | --- | --- | --- |
| <xkb> SUPER | Return | exec | $term </xkb> |
| <xkb> SUPERSHIFT | return | exec | $floating_term </xkb> |
| <xkb> SUPERSHIFT | Y | exec | $floating_term --class NVIM -e nvim </xkb> |
| <xkb> SUPER | D | exec | wofi --show run </xkb> |
| <xkb> SUPERSHIFT | n | exec | swaync-client -t -sw </xkb> |
| <xkb> SUPER | F1 | exec | $browser </xkb> |
| <xkb> SUPER | F2 | exec | $filemanager </xkb> |
| <xkb> SUPER | F3 | exec | $discord </xkb> |
| <xkb> SUPER | F4 | exec | $spotify </xkb> |
| <xkb> SUPER | F5 | exec | $office </xkb> |
| <xkb> SUPER | F6 | exec | $mailer </xkb> |
| <xkb> SUPER | F7 | exec | $videcon </xkb> |
| <xkb> SUPER | F8 | exec | $calendar </xkb> |
