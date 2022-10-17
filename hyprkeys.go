package main

import (
	"bufio"
	"os"
	"strings"
)

// Read Hyprland configuration file and return lines that start with bind= and bindm=
// Hyprland configuration file is stored in $HOME/.config/hypr/hyprland.conf
// A keyboard bind looks like this: bind=MOD,KEY,exec,COMMAND
// A mouse bind looks like this: bindm=MOD,KEY,exec,COMMAND
// We want to return the keys like this:
// Keybind = | <kbd>SUPER + L</kbd> | firefox | , firefox
// and put them in a markdown table
func readHyprlandConfig() []string {
	file, err := os.Open(os.Getenv("HOME") + "/.config/hypr/hyprland.conf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var keybinds []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "bind=") {
			keybinds = append(keybinds, line)
		} else if strings.HasPrefix(line, "bindm=") {
			keybinds = append(keybinds, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return keybinds
}

// Return each keybind as a markdown table row
// like this: | <kbd>SUPER + L</kbd> | firefox | , firefox

func keybindsToMarkdown(keybinds []string) []string {
	var markdown []string
	for _, keybind := range keybinds {
		keybind = strings.TrimPrefix(keybind, "bind=")
		keybind = strings.TrimPrefix(keybind, "bindm=")
		keybind = strings.Replace(keybind, ",", " | ", 2)
		keybind = strings.Replace(keybind, ",", " | ", 1)
		keybind = "| <xkb> " + keybind + " </xkb> |"
		markdown = append(markdown, keybind)
	}
	return markdown
}

func main() {
	keybinds := readHyprlandConfig()
	// Return each keybind on a new line before converting to markdown
	for _, keybind := range keybinds {
		println(keybind)
	}
	markdown := keybindsToMarkdown(keybinds)
	// Return each keybind on a new line after converting to markdown
	for _, keybind := range markdown {
		println(keybind)
	}
}
