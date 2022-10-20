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
func readHyprlandConfig() ([]string, []string) {
	// If --test is passed as an argument, read the test file
	//file, err := os.Open(os.Getenv("HOME") + "/.config/hypr/hyprland.conf")
	file, err := os.Open("test/hyprland.conf") // testing config
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var kbKeybinds []string
	var mKeybinds []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "bind=") {
			kbKeybinds = append(kbKeybinds, line)
		} else if strings.HasPrefix(line, "bindm=") {
			mKeybinds = append(mKeybinds, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return kbKeybinds, mKeybinds
}

// Return each keybind as a markdown table row
// like this: | <kbd>SUPER + L</kbd> | firefox | , firefox

// Pass both kbKeybinds and mKeybinds to this function
func keybindsToMarkdown(kbKeybinds, mKeybinds []string) []string {
	var markdown []string
	for _, keybind := range kbKeybinds {
		keybind = strings.TrimPrefix(keybind, "bind=")
		keybind = strings.TrimPrefix(keybind, "bindm=")

		// Split "keybind" into a slice of strings
		// based on the comma delimiter
		keybindSlice := strings.SplitN(keybind, ",", 4)

		// Trim whitespace from keybindSlice[1] to keybindSlice[3]
		keybindSlice[1] = strings.TrimSpace(keybindSlice[1])
		keybindSlice[2] = strings.TrimSpace(keybindSlice[2])
		keybindSlice[3] = strings.TrimSpace(keybindSlice[3])

		// Print the keybind as a markdown table row
		//markdown = append(markdown, "| <kbd>"+keybindSlice[0]+" + "+keybindSlice[1]+"</kbd> | "+keybindSlice[2]+" | "+keybindSlice[3]+" |")

		// Check if keybindSlice is null
		if keybindSlice[0] == "" {
			keybindSlice[1] = strings.TrimSpace(keybindSlice[1])
			markdown = append(markdown, "| <kbd>"+keybindSlice[1]+"</kbd> | "+keybindSlice[2]+" | "+keybindSlice[3]+" |")

		} else {
			markdown = append(markdown, "| <kbd>"+keybindSlice[0]+" + "+keybindSlice[1]+"</kbd> | "+keybindSlice[2]+" | "+keybindSlice[3]+" |")
		}
	}

	for _, keybind := range mKeybinds {
		keybind = strings.TrimPrefix(keybind, "bindm=")

		// Split "keybind" into a slice of strings
		// based on the comma delimiter
		keybindSlice := strings.SplitN(keybind, ",", 3)

		// Trim whitespace from keybindSlice[1] to keybindSlice[2]
		keybindSlice[1] = strings.TrimSpace(keybindSlice[1])
		keybindSlice[2] = strings.TrimSpace(keybindSlice[2])

		// Print the keybind as a markdown table row
		//markdown = append(markdown, "| <kbd>"+keybindSlice[0]+" + "+keybindSlice[1]+"</kbd> | "+keybindSlice[2]+" |")

		// Check if keybindSlice[0] is null
		if keybindSlice[0] == "" {
			markdown = append(markdown, "| <kbd>"+keybindSlice[1]+"</kbd> | | "+keybindSlice[2]+" |")
		} else {
			// put an "| |" inbetween the keybindSlice[0] and keybindSlice[1]
			markdown = append(markdown, "| <kbd>"+keybindSlice[0]+" + "+keybindSlice[1]+"</kbd> | | "+keybindSlice[2]+" |")

		}

	}
	return markdown
}

func main() {
	kbKeybinds, mKeybinds := readHyprlandConfig()
	// If --verbose is passed as an argument, print the keybinds
	// to the terminal
	if len(os.Args) > 1 && os.Args[1] == "--verbose" {
		for _, keybind := range kbKeybinds {
			println(keybind)
		}
		for _, keybind := range mKeybinds {
			println(keybind)
		}
	}

	// If --markdown is passed as an argument, print the keybinds
	// as a markdown table
	if len(os.Args) > 1 && os.Args[1] == "--markdown" {
		markdown := keybindsToMarkdown(kbKeybinds, mKeybinds)
		println("| Keybind | Dispatcher | Command |")
		println("|---------|------------|---------|")
		for _, row := range markdown {
			println(row)
		}
	}

}
