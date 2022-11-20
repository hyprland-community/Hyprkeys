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
		line = strings.Trim(line, " \n")
		if len(line) > 0 {
			out += line + "\n"
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
	return TrimBlock(out)
}

func ParseGlobal(content string) *props.S_global {
	lines := strings.Split(content, "\n")
	global := props.NewGlobal()
	for _, line := range lines {
		line = strings.Trim(line, " ")
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			parts[0] = strings.Trim(parts[0], " ")
			parts[1] = strings.Trim(parts[1], " ")
			if strings.HasPrefix(line, "$") {
				global.S_variables[parts[0]] = parts[1]
			} else if strings.HasPrefix(line, "bind") {
				global.S_binds = append(global.S_binds, map[string][]string{parts[0]: strings.Split(parts[1], ",")})
			}
		} else {
			global.S_raw += line + "\n"
		}
	}
	return global
}

func ParseConfig(blocks map[string]string) props.Config {
	defaults := props.NewConf()
	for rawlabel, block := range blocks {
		label := strings.ToUpper(string(rawlabel[0])) + rawlabel[1:]
		if label == "Global" {
			reflections.SetField(&defaults, label, ParseGlobal(block))
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

func BuildGlobal(glob props.S_global) string {
	var out string
	for key, val := range glob.S_variables {
		out += key + " = " + val + "\n"
	}
	out += "\n"
	for _, binds := range glob.S_binds {
		for key, val := range binds {
			out += key + " = " + strings.Join(val, ",") + "\n"
		}
	}
	out += glob.S_raw
	return out
}

func BuildConf(conf props.Config) string {
	output := "#-----------------------------#\n#    Generated By HyprKeys    #\n#-----------------------------#\n\n"
	fields, err := reflections.Fields(conf)
	if err != nil {
		fmt.Println("error getting fields", err)
	}
	for _, field := range fields {
		block, err := reflections.GetField(conf, field)
		if err != nil {
			fmt.Println("error getting block", err)
			continue
		}
		field = strings.Replace(field, "S_", "", 1)
		field = strings.ToLower(field)
		if field == "global" {
			output += BuildGlobal(*conf.Global)
			continue
		}
		output += field
		output += " {\n"
		block_fields, err := reflections.Fields(block)
		if err != nil {
			fmt.Println("error getting block fields", err)
			continue
		}
		for _, block_field := range block_fields {
			val, err := reflections.GetField(block, block_field)
			if err != nil {
				fmt.Println("error getting block field", err)
				continue
			}
			block_field = strings.Replace(block_field, "__", ".", 1)
			block_field = strings.TrimPrefix(block_field, "S_")
			val_fields, err := reflections.Fields(val)
			if val_fields == nil || err != nil {
				fmt.Println(val_fields, err)
				output += "    " + block_field + " = " + fmt.Sprint(val) + "\n"
			} else {
				fmt.Println("got field: ", val_fields)
				output += "    " + block_field + " {\n"
				for _, val_field := range val_fields {
					val_val, err := reflections.GetField(val, val_field)
					if err != nil {
						fmt.Println("error getting block field", err)
						continue
					}
					val_field = strings.Replace(val_field, "__", ".", 1)
					val_field = strings.TrimPrefix(val_field, "S_")
					output += "        " + val_field + " = " + fmt.Sprint(val_val) + "\n"
				}
				output += "    }\n"
			}
		}
		output += "}\n\n"
	}
	return output
}
