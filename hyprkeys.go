package main

import (
	"bufio"
	"fmt"
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
func readHyprlandConfig() ([]string, []string, []string, map[string]string) {
	// If --test is passed as an argument, read the test file
	//file, err := os.Open(os.Getenv("HOME") + "/.config/hypr/hyprland.conf")
	file, err := os.Open("test/hyprland.conf") // testing config
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[string]string)

	var kbKeybinds []string
	var mKeybinds []string
	var variables []string
	var variableMap = m

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "bind=") {
			kbKeybinds = append(kbKeybinds, line)
		} else if strings.HasPrefix(line, "bindm=") {
			mKeybinds = append(mKeybinds, line)
		} else if strings.HasPrefix(line, "$") {
			// Probably not the best way to do this, but can't think of another occasion where a line would start with "$"
			// and include "=", yet still not be a variable
			if strings.Contains(line, "=") {
				// Store variables and their values in a map
				// This will be used to replace variables in the markdown table
				// with their values
				variable := strings.SplitN(line, "=", 2)
				variableMap[variable[0]] = variable[1]
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return kbKeybinds, mKeybinds, variables, variableMap
}

// func parseVariables(variables []string) []string {
// 	var parsedVariables []string
// 	for _, variable := range variables {
// 		variable = strings.TrimPrefix(variable, "$")
// 		variableSlice := strings.SplitN(variable, "=", 2)
// 		variableSlice[0] = strings.TrimSpace(variableSlice[0])
// 		variableSlice[1] = strings.TrimSpace(variableSlice[1])
// 		parsedVariables = append(parsedVariables, variableSlice[0]+" = "+variableSlice[1])
// 	}

// 	return parsedVariables
// }

// Return each keybind as a markdown table row
// like this: | <kbd>SUPER + L</kbd> | firefox | , firefox
// we also account for no MOD key.

// Pass both kbKeybinds and mKeybinds to this function
func keybindsToMarkdown(kbKeybinds, mKeybinds []string) []string {
	var markdown []string
	for _, keybind := range kbKeybinds {
		keybind = strings.TrimPrefix(keybind, "bind=")

		// Split "keybind" into a slice of strings
		// based on the comma delimiter
		keybindSlice := strings.SplitN(keybind, ",", 4)

		// Trim whitespace from keybindSlice[1] to keybindSlice[3]
		keybindSlice[1] = strings.TrimSpace(keybindSlice[1])
		keybindSlice[2] = strings.TrimSpace(keybindSlice[2])
		keybindSlice[3] = strings.TrimSpace(keybindSlice[3])

		// Print the keybind as a markdown table row

		// Check if keybindSlice is empty
		// Trim the whitespace and "+" if it is
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

		// Check if keybindSlice[0] is null
		// Trim the whitespace and "+" if it is
		if keybindSlice[0] == "" {
			markdown = append(markdown, "| <kbd>"+keybindSlice[1]+"</kbd> | | "+keybindSlice[2]+" |")
		} else {
			// put "| |" inbetween the keybindSlice[0] and keybindSlice[1]
			markdown = append(markdown, "| <kbd>"+keybindSlice[0]+" + "+keybindSlice[1]+"</kbd> | | "+keybindSlice[2]+" |")
		}

	}

	return markdown
}

func main() {
	kbKeybinds, mKeybinds, variables, variableMap := readHyprlandConfig()

	// If --verbose is passed as an argument, print the keybinds
	// to the terminal
	if len(os.Args) > 1 && os.Args[1] == "--verbose" {
		for _, keybind := range kbKeybinds {
			fmt.Println(keybind)

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

	if len(os.Args) > 1 && os.Args[1] == "--variables" {
		for _, variable := range variables {
			println(variable)
		}

		// Now we replace the variables in the markdown table with their values
		// and print the table if --markdown is also passed as an argument
		markdown := keybindsToMarkdown(kbKeybinds, mKeybinds)
		println("| Keybind | Dispatcher | Command |")
		println("|---------|------------|---------|")
		for _, row := range markdown {
			for key, value := range variableMap {

				row = strings.ReplaceAll(row, key, value)
			}
			println(row)
		}
	}
}
