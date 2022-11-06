package parser

import (
	"fmt"
	"strings"
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
			if firstpass {
				// removes the label names from the prev block contents
				con := depth[len(depth)-1]
				depth[len(depth)-1] = con[:len(con)-len(label_buffer)]
				con = depth[0]
				depth[0] = con[:len(con)-len(label_buffer)]
			}
		}
		if letter == "{" {
			firstpass = true // without this the global settings before first block dont show up properly :/
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

	for label, block := range blocks {
		fmt.Println(label + "::\n" + block)
	}
	return blocks
}
