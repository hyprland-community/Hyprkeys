# Hyprkeys

A simple, and relatively fast application that returns your Hyprland keys.

*No, not just keybinds. We now return config sections as well*


> ⚠️ Hyprkeys is currently in early development and is not ready for use. Not like it can
break anything, but you will have wasted your time.

In case you still want to use it, instructions are below:

**(Looking for testers and contributors, Go is not a language I often work with.)**

## Installation & Usage

1. Download Go. You can find it [here](https://golang.org/dl/)
2. Clone this repository with `git clone https://github.com/notashelf/hyprkeys`
3. Build the application with `go build` and run it with `./hyprkeys`. This will display a list of command line options.

**Alternatively, open this directory and run `go run .` to run without compiling.**

## Project Roadmap

- [x] Format keybinds better, maybe with a proper table
  - [x]  Remove the `+` in the keybinds that don't have modifiers
  - [x]  Add an extra column to mouse keybinds to match table titles
- [x] Optionally (--variables) parse variables and replace them with their actual value
- [x] Account for bind flags, that may be passed in any random order
  - [x] Figure out a regex to match the flags
    - [x] Figure out why the regex doesn't work
- [ ] Account for line comments in rows
- [ ] Break code into multiple files
    - [ ] move command line parsing to a separate file
- [ ] Command line options
  - [ ] Sort output by dispatcher
  - [ ] Account for multiple arguments being passed at once
- [ ] Somehow account for keybinds can be set dynamically? (I don't know how to do this)
  - [ ] Add instructions for a pipe to `hyprkeys` to get the keybinds from
- [ ] Convert sway keybinds to Hyprland keybinds with `--convert`
  - [ ] Possibly more wayland compositors, sway is enough for now.
- [ ] Get more than just keybinds, try and get all config options seperated by section
  - [ ] Potentially rename the project
- [ ] Add a way to change keybinds with Hyprkeys after reading them (???)
- [ ] Packaging for AUR and maybe other distros (why would anyone want that)


### Current TODOs

- [ ] (hyprkeys.go) Trim "bind = " from keybinds before printing in markdown
- [ ] (hyprkeys.go) Switch regex mechanism
- [x] (hyprkeys.go) `--help flag`
Search for `TODO` in the code for more information.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contributing

If you want to contribute, feel free to open a pull request. I'll try to review it as soon as possible.

## Example Output (as of 20 NOV 2022)

**Moved to [test/markdown.md](test/markdown.md)**
