package main

import (
	"notashelf.dev/hyprkeys/cmd"
)

var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}
