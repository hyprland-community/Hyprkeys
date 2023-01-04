package reader

import (
	"bufio"
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

type ConfigValues struct {
	Settings      Settings
	KeyboardBinds []string
	MouseBinds    []string
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

	var kbKeybinds []string
	var mKeybinds []string
	// var variables []string

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "bindm"):
			mKeybinds = append(mKeybinds, line)
		case strings.HasPrefix(line, "bind"):
			kbKeybinds = append(kbKeybinds, line)
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

func generalHandler(name string, scanner *bufio.Scanner) *Setting {
	settings := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "}") {
			return &Setting{
				Name:     name,
				Settings: settings,
			}
		}
		trimmed := strings.TrimSpace(line)
		val := strings.Split(strings.ReplaceAll(trimmed, " ", ""), "=")
		if len(val) > 1 {
			if val[1] != "" {
				settings[val[0]] = val[1]
			}
		}
	}
	return &Setting{
		Name:     name,
		Settings: settings,
	}
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
