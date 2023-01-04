package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// io/ioutil is deprecated, use io and os packages instead
)

type Settings []*Setting

type Setting struct {
	Name          string
	Settings      map[string]string
	SubCategories []Setting
}

type Keybind struct {
	BindType   string
	Bind       string
	Dispatcher string
	Command    string
	Comments   string
}

type ConfigValues struct {
	Settings      Settings
	KeyboardBinds []*Keybind
	MouseBinds    []*Keybind
}

// Read Hyprland configuration file and return lines that start with bind= and bindm=
func ReadHyprlandConfig(configPath string) (*ConfigValues, error) {
	settings := make(Settings, 0)
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var kbKeybinds []*Keybind
	var mKeybinds []*Keybind
	// var variables []string

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "bindm"):
			mKeybinds = append(mKeybinds, makeBind(line))
		case strings.HasPrefix(line, "bind"):
			kbKeybinds = append(kbKeybinds, makeBind(line))
		case strings.HasPrefix(line, "input"):
			settings = append(settings, handler("input", scanner))
		case strings.HasPrefix(line, "general"):
			settings = append(settings, handler("general", scanner))
		case strings.HasPrefix(line, "decoration"):
			settings = append(settings, handler("decoration", scanner))
		case strings.HasPrefix(line, "animations"):
			settings = append(settings, handler("animations", scanner))
		case strings.HasPrefix(line, "gestures"):
			settings = append(settings, handler("gestures", scanner))
		case strings.HasPrefix(line, "misc"):
			settings = append(settings, handler("misc", scanner))
		case strings.HasPrefix(line, "binds"):
			settings = append(settings, handler("binds", scanner))
		case strings.HasPrefix(line, "debug"):
			settings = append(settings, handler("debug", scanner))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	configValues := &ConfigValues{
		Settings:      settings,
		KeyboardBinds: kbKeybinds,
		MouseBinds:    mKeybinds,
	}
	return configValues, nil
}

func makeBind(bind string) *Keybind {
	split := strings.SplitN(bind, "=", 2)
	keyBind := &Keybind{
		BindType: strings.TrimSpace(split[0]),
	}
	bind = strings.TrimSpace(split[1])

	// Split "bind" into a slice of strings
	// based on the comma delimiter
	keybindSlice := strings.SplitN(bind, ",", 4)

	// if it is just a dispatcher then add blank command
	if len(keybindSlice) < 4 {
		keybindSlice = append(keybindSlice, "")
	}

	// Trim whitespace from keybindSlice[1] to keybindSlice[3]
	keybindSlice[1] = strings.TrimSpace(keybindSlice[1])
	keybindSlice[2] = strings.TrimSpace(keybindSlice[2])
	keybindSlice[3] = strings.TrimSpace(keybindSlice[3])

	keyBind.Dispatcher = keybindSlice[2]
	keyBind.Command = keybindSlice[3]

	// Check if keybindSlice is empty
	// Trim the whitespace and "+" if it is
	if keybindSlice[0] == "" {
		keybindSlice[1] = strings.TrimSpace(keybindSlice[1])
		keyBind.Bind = keybindSlice[1]
	} else {
		keyBind.Bind = fmt.Sprintf("%s %s", keybindSlice[0], keybindSlice[1])
	}

	lastString := keybindSlice[3]
	// comment handler
	if keybindSlice[3] == "" {
		lastString = keybindSlice[2]
	}
	comments := strings.SplitN(lastString, "#", 2)
	if len(comments) > 1 {
		lastString = comments[0]
		keyBind.Comments = strings.TrimSpace(comments[1])
	}
	if keybindSlice[3] == "" {
		keyBind.Dispatcher = lastString
	} else {
		keyBind.Command = lastString
	}
	return keyBind
}

func handler(name string, scanner *bufio.Scanner) *Setting {
	input := make(map[string]string)
	touchpad := make(map[string]string)
	touchdevice := make(map[string]string)
	settings := &Setting{
		Name: name,
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "}") {
			settings.Settings = input
			return settings
		}
		if strings.HasPrefix(strings.TrimSpace(line), "touchpad") {
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(strings.TrimSpace(line), "}") {
					break
				}
				trimmed := strings.TrimSpace(line)
				val := strings.Split(strings.ReplaceAll(trimmed, " ", ""), "=")
				if len(val) > 1 {
					if val[1] != "" {
						touchpad[val[0]] = val[1]
					}
				}
			}
			subcategory := &Setting{
				Name:     "touchpad",
				Settings: touchpad,
			}
			settings.SubCategories = append(settings.SubCategories, *subcategory)
		}
		if strings.HasPrefix(strings.TrimSpace(line), "touchdevice") {
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(strings.TrimSpace(line), "}") {
					break
				}
				trimmed := strings.TrimSpace(line)
				val := strings.Split(strings.ReplaceAll(trimmed, " ", ""), "=")
				if len(val) > 1 {
					if val[1] != "" {
						touchdevice[val[0]] = val[1]
					}
				}
			}
			subcategory := &Setting{
				Name:     "touchdevice",
				Settings: touchpad,
			}
			settings.SubCategories = append(settings.SubCategories, *subcategory)
		}
		trimmed := strings.TrimSpace(line)
		val := strings.Split(strings.ReplaceAll(trimmed, " ", ""), "=")
		if len(val) > 1 {
			if val[1] != "" {
				input[val[0]] = val[1]
			}
		}
	}
	return settings
}
