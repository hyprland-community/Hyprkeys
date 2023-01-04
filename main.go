package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/pborman/getopt"

	// io/ioutil is deprecated, use io and os packages instead
	"notashelf.dev/hyprkeys/flags"
	"notashelf.dev/hyprkeys/reader"
)

func main() {
	flags := flags.ReadFlags()

	if flags.Version {
		version := "unknown"
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
		fmt.Println("version:", version)
		return
	}

	if flags.Help {
		getopt.Usage()
		return
	}

	if flags.ConfigPath == "" {
		flags.ConfigPath = filepath.Join(os.Getenv("HOME"), ".config/hypr/hyprland.conf")
	}
	if flags.Test {
		flags.ConfigPath = "test/hyprland.conf"
	}

	configValues, err := reader.ReadHyprlandConfig(flags.ConfigPath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if flags.Raw {
		for _, bind := range configValues.KeyboardBinds {
			fmt.Println(bind)
		}
		for _, bind := range configValues.MouseBinds {
			fmt.Println(bind)
		}
		for _, val := range configValues.Settings {
			fmt.Println(val.Name, "{")
			for setting, value := range val.Settings {
				fmt.Println("\t", setting, "=", value)
			}
			for _, set := range val.SubCategories {
				fmt.Println("\t", set.Name, "{")
				for setting, value := range set.Settings {
					fmt.Println("\t\t", setting, "=", value)
				}
				fmt.Println("\t", "}")
			}
			fmt.Println("}")
		}
	}
	if flags.Json {
		out, err := json.MarshalIndent(configValues, "", " ")
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(string(out))
	}
	if flags.Markdown {
		md := keybindsToMarkdown(configValues.KeyboardBinds, configValues.MouseBinds)
		println("| Keybind | Dispatcher | Command |")
		println("|---------|------------|---------|")
		for _, row := range md {
			println(row)
		}
	}
}

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
