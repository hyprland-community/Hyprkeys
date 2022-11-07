package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oleiade/reflections"
	props "notashelf.dev/hyprkeys/util/properties"
)

// check if word is a valid label
func IsLabel(word string, s []string, i int, content string) bool {
	for _, w := range s {
		if strings.Trim(strings.Trim(w, " "), "\n") == strings.Trim(strings.Trim(word, " "), "\n") {
			lastspace := strings.Split(content[i-len(word):], "\n")
			lastline := strings.Split(lastspace[0], " ")
			lastword := lastline[len(lastline)-1]
			if strings.HasSuffix(lastword, "{") {
				return true
			}
		}
	}
	return false
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// gets the word to the left of index i from content
func GetLabel(i int, content string) string {
	lines := strings.Split(content[:i], "\n")
	part := strings.Split(lines[len(lines)-1], " ")
	for i := range part {
		word := part[len(part)-1-i]
		if len(word) > 0 {
			return strings.TrimSpace(word)
		}
	}
	return ""
}

func TrimBlock(block string) string {
	out := ""
	for _, line := range strings.Split(block, "\n") {
		if len(strings.Split(line, " ")) > 0 {
			out += strings.Trim(line, " \n") + "\n"
		}
	}
	return out
}

func ParseBlocks(content string) map[string]string {
	blocks := make(map[string]string)
	depth := []string{""} // keeps track of the blocks currently in scope and their depths(index)

	labels := []string{}
	depth_label := []string{}
	label_buffer := "" // probably a variable to store the previous word before a block starts (to identify labels)

	firstpass := false

	//get all the labels
	for i, letter := range content {
		letter := string(letter)
		if letter == "{" {
			label := GetLabel(i, content)
			labels = append(labels, label)
		}
	}

	for i, letter := range content {
		letter := string(letter)
		if IsLabel(label_buffer, labels, i, content) == true {
			firstpass = true // without this the global settings before first block dont show up properly :/
			// removes the label names from the prev block contents
			if firstpass {
				con := depth[len(depth)-1]
				depth[len(depth)-1] = con[:len(con)-len(label_buffer)]
			}
		}
		if letter == "{" {
			depth_label = append(depth_label, GetLabel(i, content))
			label_buffer = ""
			depth = append(depth, "")
		} else if letter == "}" {
			// move block from `depth` to `blocks` as it goes out of scope now
			blocks[depth_label[len(depth_label)-1]] = TrimBlock(depth[len(depth)-1])
			depth_label = RemoveIndex(depth_label, len(depth_label)-1)
			depth = RemoveIndex(depth, len(depth)-1)

		} else {
			if letter == " " || letter == "\n" {
				label_buffer = ""
			} else {
				label_buffer = label_buffer + letter
			}
			// append content to the currently focused block in scope
			depth[len(depth)-1] += letter
		}

	}

	// set the global block
	blocks["global"] = TrimBlock(depth[0])
	return blocks
}

func ParseComments(content string) string {
	out := ""
	for _, line := range strings.Split(content, "\n") {
		for _, i := range line {
			if i == '#' {
				break
			} else {
				out += string(i)
			}
		}
		out += "\n"
	}
	return out
}

func ParseConfig(blocks map[string]string) props.Config {
	defaults := props.NewConf()
	for rawlabel, block := range blocks {
		label := strings.ToUpper(string(rawlabel[0])) + rawlabel[1:]
		if label == "Global" {
			reflections.SetField(&defaults, label, block)
			continue
		}
		var section interface{}
		var err error
		if label == "Touchpad" || label == "Touchdevice" {
			section, err = reflections.GetField(defaults.Input, label)
		} else {
			section, err = reflections.GetField(&defaults, label)
		}
		if err != nil {
			fmt.Println("error parsing a label: "+label, err)
			continue
		}
		lines := strings.Split(block, "\n")
		keyval := make(map[string]string)
		for _, i := range lines {
			pairs := strings.Split(i, "=")
			pairs[0] = strings.Trim(pairs[0], " \n")
			if len(pairs) == 2 {
				pairs[1] = strings.Trim(pairs[1], " \n")
				keyval[pairs[0]] = pairs[1]
			} else {
				keyval[pairs[0]] = ""
			}
		}
		for key, val := range keyval {
			key = strings.Replace(key, ".", "__", 1)
			key = "S_" + key
			val = strings.Trim(val, " \n")
			val = strings.ReplaceAll(val, "yes", "true")
			val = strings.ReplaceAll(val, "no", "false")
			val = strings.ReplaceAll(val, "on", "true")
			val = strings.ReplaceAll(val, "off", "false")
			fieldt, err := reflections.GetFieldType(section, key)
			if err != nil {
				fmt.Println("error parsing a field: "+key, err)
				continue
			}
			if fieldt == "bool" {
				parsed_val, err := strconv.ParseBool(val)
				if err != nil {
					fmt.Println("error parsing a type: "+fieldt+"  with data: "+val, err)
					continue
				}
				err = reflections.SetField(section, key, parsed_val)
				if err != nil {
					fmt.Println("failed setting field: "+key+"  value: "+val+"||", err)
					continue
				}
			} else if fieldt == "int64" {

				parsed_val, err := strconv.ParseInt(val, 6, 64)
				if err != nil {
					fmt.Println("error parsing a type: "+fieldt+"  with data: "+val, err)
					continue
				}
				err = reflections.SetField(section, key, parsed_val)
				if err != nil {
					fmt.Println("failed setting field: "+key+"  value: "+val+"||", err)
					continue
				}
			} else if fieldt == "float64" {

				parsed_val, err := strconv.ParseFloat(val, 64)
				if err != nil {
					fmt.Println("error parsing a type: "+fieldt+"  with data: "+val, err)
					continue
				}
				err = reflections.SetField(section, key, parsed_val)
				if err != nil {
					fmt.Println("failed setting field: "+key+"  value: "+val+"||", err)
					continue
				}
			} else if fieldt == "[2]float64" {
				vec := strings.Split(val, " ")
				parsed_val := [2]int{0, 0}
				for i, v := range vec[:2] {
					parsed_v, err := strconv.ParseFloat(v, 64)
					if err != nil {
						fmt.Println("error parsing a type: "+fieldt+"  with data: "+val, err)
						break
					}
					parsed_val[i] = int(parsed_v)
				}

				err = reflections.SetField(section, key, parsed_val)
				if err != nil {
					fmt.Println("failed setting field: "+key+"  value: "+val+"||", err)
					continue
				}
			} else {
				err = reflections.SetField(section, key, val)
				if err != nil {
					fmt.Println("failed setting field(final else): "+key+"  with data: "+val, err)
					continue
				}
			}
		}
	}
	return defaults
}

func Parse(content string) props.Config {
	content = ParseComments(content)
	blocks := ParseBlocks(content)
	return ParseConfig(blocks)
}
