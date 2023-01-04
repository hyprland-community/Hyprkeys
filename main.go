package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"

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
	if !(len(os.Args) > 1) || flags.Help {
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
		rawHandler(configValues, flags)
	}

	if flags.Json {
		jsonHandler(configValues, flags)
	}

	if flags.Markdown {
		markdownHandler(configValues, flags)
	}
}

func markdownHandler(configValues *reader.ConfigValues, flags *flags.Flags) error {
	md := keybindsToMarkdown(configValues.KeyboardBinds, configValues.MouseBinds)
	out := ""
	out += "| Keybind | Dispatcher | Command |\n"
	out += "|---------|------------|---------|\n"
	for _, row := range md {
		out += row + "\n"
	}
	fmt.Println(out)
	if flags.Output != "" {
		err := os.WriteFile(flags.Output, []byte(out), 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func jsonHandler(configValues *reader.ConfigValues, flags *flags.Flags) error {
	out, err := json.MarshalIndent(configValues, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	if flags.Output != "" {
		err := os.WriteFile(flags.Output, out, 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func rawHandler(configValues *reader.ConfigValues, flags *flags.Flags) error {
	out := ""
	for _, bind := range configValues.KeyboardBinds {
		out += fmt.Sprintf("%s = %s %s %s", bind.BindType, bind.Bind, bind.Dispatcher, bind.Command) + "\n"
	}
	for _, bind := range configValues.MouseBinds {
		out += fmt.Sprintf("%s = %s %s %s", bind.BindType, bind.Bind, bind.Dispatcher, bind.Command) + "\n"
	}
	for _, val := range configValues.Settings {
		out += val.Name + " {" + "\n"
		for setting, value := range val.Settings {
			out += "\t" + setting + " = " + value + "\n"
		}
		for _, set := range val.SubCategories {
			out += "\t" + set.Name + " {\n"
			for setting, value := range set.Settings {
				out += "\t\t" + setting + " = " + value + "\n"
			}
			out += "\t}\n"
		}
		out += "}\n"
	}
	fmt.Print(out)
	if flags.Output != "" {
		err := os.WriteFile(flags.Output, []byte(out), 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

// Pass both kbKeybinds and mKeybinds to this function
func keybindsToMarkdown(kbKeybinds, mKeybinds []*reader.Keybind) []string {
	var markdown []string
	for _, keybind := range kbKeybinds {
		markdown = append(markdown, "| <kbd>"+keybind.Bind+"</kbd> | "+keybind.Dispatcher+" | "+keybind.Command+" |")
	}
	for _, keybind := range mKeybinds {
		markdown = append(markdown, "| <kbd>"+keybind.Bind+"</kbd> | "+keybind.Dispatcher+" | "+keybind.Command+" |")
	}
	return markdown
}
