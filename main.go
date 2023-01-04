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
	if !(len(os.Args) > 1) || flags.Help {
		getopt.Usage()
		return
	}

	if flags.ConfigPath == "" {
		flags.ConfigPath = filepath.Join(os.Getenv("HOME"), ".config/hypr/hyprland.conf")
	}

	configValues, err := reader.ReadHyprlandConfig(flags.ConfigPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if flags.GetBind != "" {
		getBindHandler(configValues, flags)
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

func getBindHandler(configValues *reader.ConfigValues, flags *flags.Flags) error {
	matchedBinds := make([]*reader.Keybind, 0)
	for _, val := range configValues.KeyboardBinds {
		if strings.Contains(val.Dispatcher, flags.GetBind) || strings.Contains(val.Command, flags.GetBind) {
			matchedBinds = append(matchedBinds, val)
		}
	}
	for _, val := range configValues.MouseBinds {
		if strings.Contains(val.Dispatcher, flags.GetBind) || strings.Contains(val.Command, flags.GetBind) {
			matchedBinds = append(matchedBinds, val)
		}
	}
	out, err := json.MarshalIndent(matchedBinds, "", " ")
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

func markdownHandler(configValues *reader.ConfigValues, flags *flags.Flags) error {
	md := keybindsToMarkdown(configValues.KeyboardBinds, configValues.MouseBinds)
	out := ""
	out += "| Keybind | Dispatcher | Command | Comments |\n"
	out += "|---------|------------|---------|----------|\n"
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
		out += fmt.Sprintf("%s = %s %s %s #%s", bind.BindType, bind.Bind, bind.Dispatcher, bind.Command, bind.Comments) + "\n"
	}
	for _, bind := range configValues.MouseBinds {
		out += fmt.Sprintf("%s = %s %s %s #%s", bind.BindType, bind.Bind, bind.Dispatcher, bind.Command, bind.Comments) + "\n"
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
		markdown = append(markdown, "| <kbd>"+keybind.Bind+"</kbd> | "+keybind.Dispatcher+" | "+strings.ReplaceAll(keybind.Command, "|", "\\|")+" | "+strings.ReplaceAll(keybind.Comments, "|", "\\|")+" |")
	}
	for _, keybind := range mKeybinds {
		markdown = append(markdown, "| <kbd>"+keybind.Bind+"</kbd> | "+keybind.Dispatcher+" | "+strings.ReplaceAll(keybind.Command, "|", "\\|")+" |")
	}
	return markdown
}
