package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"notashelf.dev/hyprkeys/src/config"
	"notashelf.dev/hyprkeys/src/ctl"
	"notashelf.dev/hyprkeys/src/reader"
)

func OutputConfig(configValues *reader.ConfigValues, conf config.Flags) error {
	configValues.Binds = filterBinds(configValues, conf)
	if conf.Raw {
		return rawHandler(configValues, conf)
	}
	if conf.Markdown {
		return markdownHandler(configValues, conf)
	}
	if conf.Json {
		return jsonHandler(configValues, conf)
	}
	return fmt.Errorf("no output flag selected")
}

func OutputCtl(ctlValues ctl.Binds, conf config.Flags) error {
	if conf.Raw {
		return ctlToRaw(ctlValues, conf)
	}
	if conf.Markdown {
		return ctlToMarkDown(ctlValues, conf)
	}
	if conf.Json {
		return ctlToJson(ctlValues, conf)
	}
	return fmt.Errorf("no output flag selected")
}

func ctlToRaw(binds ctl.Binds, conf config.Flags) error {
	fmt.Printf("%+v\n", binds)
	if conf.Output != "" {
		err := os.WriteFile(conf.Output, []byte(fmt.Sprintf("%+v", binds)), 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func ctlToJson(binds ctl.Binds, conf config.Flags) error {
	out, err := json.MarshalIndent(binds, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	if conf.Output != "" {
		err := os.WriteFile(conf.Output, out, 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func filterBinds(configValues *reader.ConfigValues, conf config.Flags) []*reader.Keybind {
	matchedBinds := make([]*reader.Keybind, 0)
	for _, val := range configValues.Binds {
		if strings.Contains(val.Dispatcher, conf.FilterBinds) || strings.Contains(val.Command, conf.FilterBinds) {
			matchedBinds = append(matchedBinds, val)
		}
	}
	return matchedBinds
}

func markdownHandler(configValues *reader.ConfigValues, conf config.Flags) error {
	md := keybindsToMarkdown(configValues.Binds, conf)
	out := ""
	for _, val := range configValues.Keywords {
		out += fmt.Sprintf("#### $%s = %s", val.Name, val.Value)
	}
	out += "\n"
	if conf.Comments {
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
	if conf.Output != "" {
		err := os.WriteFile(conf.Output, []byte(out), 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func ctlToMarkDown(binds ctl.Binds, conf config.Flags) error {
	md := ctlBindsToMarkdown(binds)
	out := ""
	out += "\n"
	out += "| Keybind | Locked | Mouse | Release | Repeat | Submap | Dispatcher | Command |\n"
	out += "|---------|--------|-------|---------|--------|--------|------------|---------|\n"
	for _, row := range md {
		out += row + "\n"
	}
	fmt.Println(out)
	if conf.Output != "" {
		err := os.WriteFile(conf.Output, []byte(out), 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func jsonHandler(configValues *reader.ConfigValues, conf config.Flags) error {
	var out []byte
	var err error

	out, err = json.MarshalIndent(configValues, "", " ")
	if err != nil {
		return err
	}

	fmt.Println(string(out))
	if conf.Output != "" {
		err := os.WriteFile(conf.Output, out, 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func rawHandler(configValues *reader.ConfigValues, conf config.Flags) error {
	out := ""
	if conf.Variables {
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
	if conf.AutoStart {
		for _, val := range configValues.AutoStart {
			out += fmt.Sprintf("%s=%s\n", val.ExecType, val.Command)
		}
	}
	if conf.Keywords {
		for _, key := range configValues.Keywords {
			out += fmt.Sprintf("$%s = %s\n", key.Name, key.Value)
		}
	}
	if conf.Binds {
		for _, bind := range configValues.Binds {
			out += fmt.Sprintf("%s = %s %s %s", bind.BindType, bind.Bind, bind.Dispatcher, bind.Command)
			if conf.Comments {
				if bind.Comments != "" {
					out += fmt.Sprintf("#%s", bind.Comments)
				}
			}
			out += "\n"
		}
	}
	fmt.Print(out)
	if conf.Output != "" {
		err := os.WriteFile(conf.Output, []byte(out), 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

// Pass both kbKeybinds and mKeybinds to this function
func keybindsToMarkdown(binds []*reader.Keybind, conf config.Flags) []string {
	var markdown []string
	for _, keybind := range binds {
		if conf.Comments {
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
