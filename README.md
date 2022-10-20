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

- [x] Format keybinds better, maybe with a proper table
  - [x]  Remove the `+` in the keybinds that don't have modifiers
  - [ ]  Add an extra column to mouse keybinds to match table titles
- [ ] Break code into multiple files, move command line parsing to a separate file
- [ ] Optionally (--con_vars) parse variables and replace them with their actual value
- [ ] Somehow account for keybinds can be set dynamically?
- [ ] Convert sway keybinds to Hyprland keybinds
- [ ] Add a way to change keybinds with Hyprkeys after reading them (???)
- [ ] Packaging for AUR and maybe other distros (why would anyone want that)

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contributing
If you want to contribute, feel free to open a pull request. I'll try to review it as soon as possible.

## Example Output (as of 20 OCT 2022)

**Moved to [test/example.md](test/markdown.md)**