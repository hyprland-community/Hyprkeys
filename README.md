# Hyprkeys

A simple, and relatively fast application that 

> ⚠️ Hyprkeys is currently in early development and is not ready for use. Not like it can
break anything, but you will have wasted your time.

In case you still want to use it, instructions are below:

**(Looking for testers and contributors, Go is not a language I often work with.)**

## Installation & Usage

1. Download Go. You can find it [here](https://golang.org/dl/)
2. Clone this repository with `git clone https://github.com/notashelf/hyprkeys`
3. Build the application with `go build` and run it with `./hyprkeys`
   - There are multiple flags you can use, see `./hyprkeys --help` for more information

Alternatively, open this directory and run `go run .` to run without compiling.

## Project Roadmap

- [x] Format keybinds better, maybe with a proper table
  - [x]  Remove the `+` in the keybinds that don't have modifiers
  - [x]  Add an extra column to mouse keybinds to match table titles
- [x] Optionally (--variables) parse variables and replace them with their actual value
- [x] Account for bind flags, that may be passed in any random order
  - [x] Figure out a regex to match the flags
    - [x] Figure out why the regex doesn't work
- [x] Break code into multiple files, move command line parsing to a separate file
- [ ] Get more than just keybinds, try and get all config options separated by section
  - [ ] Potentially rename the project
- [ ] Command line options
  - [ ] Sort output by dispatcher
  - [ ] Account for multiple arguments being passed at once
  - [ ] formatting flags for json (maybe) and markdown
  - [ ] search by config section or section:value
- [ ] Account for line comments in rows
- [ ] Somehow account for keybinds can be set dynamically? (I don't know how to do this)
  - [ ] Add instructions for a pipe to `hyprkeys` to get the keybinds from (???)
- [ ] Convert sway keybinds to Hyprland keybinds with `--convert`
  - [ ] Possibly more wayland compositors, but sway is enough for now.
- [ ] Add a way to change keybinds after reading them (???)
- [ ] Packaging for AUR and maybe other distros (why would anyone want that)

### Current TODOs

- [ ] (hyprkeys.go) Trim "bind = " from keybinds before printing in markdown
- [ ] (hyprkeys.go) Switch regex mechanism
- [ ] (hyprkeys.go) `--help flag`

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contributing

If you want to contribute, feel free to open a pull request. I'll try to review it as soon as possible.

### Contributors

- [flick0])(https://github.com/flick0) - Awesome job with the structs and general configuration parsing

## Example Output
**Moved to [test/markdown.md](test/markdown.md)**
