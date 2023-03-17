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
	"notashelf.dev/hyprkeys/ctl"
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
	if flags.Help || len(os.Args) == 0 {
		getopt.Usage()
		return
	}

	if flags.ConfigPath == "" {
		flags.ConfigPath = filepath.Join(os.Getenv("HOME"), ".config/hypr/hyprland.conf")
	}

	binds := ctl.Binds{}
	if flags.Ctl {
		var err error
		binds, err = ctl.BindsFromCtl()
		if err != nil {
			panic(err)
		}
		if flags.Json {
			out, err := json.MarshalIndent(binds, "", " ")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(out))
			if flags.Output != "" {
				err := os.WriteFile(flags.Output, out, 0o644)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			return
		}
		if flags.Raw {
			fmt.Printf("%+v\n", binds)
			if flags.Output != "" {
				err := os.WriteFile(flags.Output, []byte(fmt.Sprintf("%+v", binds)), 0o644)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			return
		}
		if flags.Markdown {
			ctlToMarkDown(binds, flags)
			return
		}

	}

	configValues, err := reader.ReadHyprlandConfig(flags)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = outputData(configValues, flags)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func outputData(configValues *reader.ConfigValues, flags *flags.Flags) error {
	configValues.Binds = filterBinds(configValues, flags)
	if flags.Markdown {
		if flags.Binds {
			return markdownHandler(configValues, flags)
		} else {
			fmt.Println("Markdown is only supported for binds currently")
			return nil
		}
	}
	if flags.Raw {
		return rawHandler(configValues, flags)
	}
	if flags.Json {
		return jsonHandler(configValues, flags)
	}
	return fmt.Errorf("no output flag selected")
}

func filterBinds(configValues *reader.ConfigValues, flags *flags.Flags) []*reader.Keybind {
	matchedBinds := make([]*reader.Keybind, 0)
	for _, val := range configValues.Binds {
		if strings.Contains(val.Dispatcher, flags.FilterBinds) || strings.Contains(val.Command, flags.FilterBinds) {
			matchedBinds = append(matchedBinds, val)
		}
	}
	return matchedBinds
}

func markdownHandler(configValues *reader.ConfigValues, flags *flags.Flags) error {
	md := keybindsToMarkdown(configValues.Binds, flags)
	out := ""
	for _, val := range configValues.Keywords {
		out += fmt.Sprintf("#### $%s = %s", val.Name, val.Value)
	}
	out += "\n"
	if flags.Comments {
		out += "| Keybind | Dispatcher | Command | Comments |\n"
		out += "|---------|------------|---------|----------|\n"
	} else {
		out += "| Keybind | Dispatcher | Command |\n"
		out += "|---------|------------|---------|\n"
	}
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

func ctlToMarkDown(binds ctl.Binds, flags *flags.Flags) error {
	md := ctlBindsToMarkdown(binds)
	out := ""
	out += "\n"
	out += "| Keybind | Locked | Mouse | Release | Repeat | Submap | Dispatcher | Command |\n"
	out += "|---------|--------|-------|---------|--------|--------|------------|---------|\n"
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
	var out []byte
	var err error

	out, err = json.MarshalIndent(configValues, "", " ")
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
	if flags.Variables {
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
	}
	if flags.AutoStart {
		for _, val := range configValues.AutoStart {
			out += fmt.Sprintf("%s=%s\n", val.ExecType, val.Command)
		}
	}
	if flags.Keywords {
		for _, key := range configValues.Keywords {
			out += fmt.Sprintf("$%s = %s\n", key.Name, key.Value)
		}
	}
	if flags.Binds {
		for _, bind := range configValues.Binds {
			out += fmt.Sprintf("%s = %s %s %s", bind.BindType, bind.Bind, bind.Dispatcher, bind.Command)
			if flags.Comments {
				if bind.Comments != "" {
					out += fmt.Sprintf("#%s", bind.Comments)
				}
			}
			out += "\n"
		}
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
func keybindsToMarkdown(binds []*reader.Keybind, flags *flags.Flags) []string {
	var markdown []string
	for _, keybind := range binds {
		if flags.Comments {
			markdown = append(markdown, "| <kbd>"+keybind.Bind+"</kbd> | "+keybind.Dispatcher+" | "+strings.ReplaceAll(keybind.Command, "|", "\\|")+" | "+strings.ReplaceAll(keybind.Comments, "|", "\\|")+" |")
		} else {
			markdown = append(markdown, "| <kbd>"+keybind.Bind+"</kbd> | "+keybind.Dispatcher+" | "+strings.ReplaceAll(keybind.Command, "|", "\\|")+" |")
		}
	}
	return markdown
}

// Pass both kbKeybinds and mKeybinds to this function
func ctlBindsToMarkdown(binds ctl.Binds) []string {
	var markdown []string
	for _, keybind := range binds {
		markdown = append(markdown, "| <kbd>"+keybind.Mods+" "+keybind.Key+"</kbd> | "+
			fmt.Sprintf("%t", keybind.Locked)+" | "+
			fmt.Sprintf("%t", keybind.Mouse)+" | "+
			fmt.Sprintf("%t", keybind.Release)+" | "+
			fmt.Sprintf("%t", keybind.Repeat)+" | "+
			strings.ReplaceAll(keybind.Submap, "|", "\\|")+" |"+
			keybind.Dispatcher+" | "+
			strings.ReplaceAll(keybind.Arg, "|", "\\|")+" | ")
	}
	return markdown
}
