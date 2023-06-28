package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"notashelf.dev/hyprkeys/src/config"
	// io/ioutil is deprecated, use io and os packages instead
)

type ConfigCategory string

type Settings []*Setting

type Setting struct {
	Name          string
	Settings      map[string]string
	SubCategories []Setting
}

type Keyword struct {
	Name  string
	Value string
}

type Keybind struct {
	BindType   string
	Bind       string
	Dispatcher string
	Command    string
	Comments   string
}

type Exec struct {
	ExecType string
	Command  string
}

type ConfigValues struct {
	Settings  Settings   `json:",omitempty"`
	AutoStart []*Exec    `json:",omitempty"`
	Binds     []*Keybind `json:",omitempty"`
	Keywords  []*Keyword `json:",omitempty"`
}

// Read Hyprland configuration file and return lines that start with bind= and bindm=
func ReadHyprlandConfig(flags config.Flags) (*ConfigValues, error) {
	categories := []string{
		"general",
		"input",
		"decoration",
		"animation",
		"gestures",
		"misc",
		"binds",
		"debug",
		"device",
	}

	subcategories := []string{
		"touchdevice",
		"touchpad",
	}

	settings := make(Settings, 0)
	file, err := os.Open(flags.ConfigPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var binds []*Keybind
	var autostart []*Exec
	var keywords []*Keyword

	for scanner.Scan() {
		line := scanner.Text()
		category, isCategory := getCategory(line, categories)
		if flags.Variables {
			if isCategory {
				settings = append(settings, handler(category, scanner, subcategories))
			}
		}
		switch {
		case strings.HasPrefix(line, "bind") && !strings.HasPrefix(line, "binds"):
			if flags.Binds {
				binds = append(binds, makeBind(line, flags))
			}
		case strings.HasPrefix(line, "$"):
			if flags.Keywords {
				keywords = append(keywords, makeKeyword(line))
			}
		case strings.HasPrefix(line, "exec"):
			if flags.AutoStart {
				autostart = append(autostart, makeExec(line))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	configValues := &ConfigValues{
		Settings:  settings,
		Binds:     binds,
		AutoStart: autostart,
		Keywords:  keywords,
	}
	return configValues, nil
}

func getCategory(check string, categories []string) (string, bool) {
	for _, cat := range categories {
		if strings.HasPrefix(strings.ToLower(check), cat) {
			categoryName := strings.TrimSpace(strings.TrimSuffix(check, "{"))
			return categoryName, true
		}
	}
	return "", false
}

func getSubCategory(check string, categories []string) (string, bool) {
	for _, cat := range categories {
		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(check)), cat) {
			return cat, true
		}
	}
	return "", false
}

func makeExec(line string) *Exec {
	split := strings.SplitN(line, "=", 2)
	exec := &Exec{
		ExecType: strings.TrimSpace(split[0]),
		Command:  strings.TrimSpace(split[1]),
	}
	return exec
}

func makeKeyword(line string) *Keyword {
	split := strings.SplitN(line, "=", 2)
	keyword := &Keyword{
		Name:  strings.TrimSpace(strings.TrimPrefix(split[0], "$")),
		Value: strings.TrimSpace(split[1]),
	}
	return keyword
}

func makeBind(bind string, flags config.Flags) *Keybind {
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
		if flags.Comments {
			keyBind.Comments = strings.TrimSpace(comments[1])
		}
	}
	if keybindSlice[3] == "" {
		keyBind.Dispatcher = lastString
	} else {
		keyBind.Command = lastString
	}
	return keyBind
}

func handler(name string, scanner *bufio.Scanner, subcategories []string) *Setting {
	input := make(map[string]string)
	settings := &Setting{
		Name: name,
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "}") {
			settings.Settings = input
			return settings
		}
		subCategory, isSubCategory := getSubCategory(line, subcategories)
		if isSubCategory {
			subSettings := make(map[string]string)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(strings.TrimSpace(line), "}") {
					break
				}
				trimmed := strings.TrimSpace(line)
				val := strings.Split(strings.ReplaceAll(trimmed, " ", ""), "=")
				if len(val) > 1 {
					if val[1] != "" {
						subSettings[val[0]] = val[1]
					}
				}
			}
			subcategory := &Setting{
				Name:     subCategory,
				Settings: subSettings,
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
