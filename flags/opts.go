package flags

import (
	"github.com/pborman/getopt"
)

type Flags struct {
	Help       bool
	Json       bool
	Markdown   bool
	Raw        bool
	Version    bool
	Variables  bool
	Blocks     bool
	GetBind    string
	Output     string
	ConfigPath string
}

func ReadFlags() *Flags {
	optHelp := getopt.BoolLong("help", 'h', "Show this help menu")
	optJson := getopt.BoolLong("json", 'j', "Return settigns as json")
	optMarkdown := getopt.BoolLong("markdown", 'm', "Print the binds as a markdown table")
	optRaw := getopt.BoolLong("raw", 'r', "Print text as is, without making it pretty")
	optVersion := getopt.BoolLong("version", 'V', "Show the version number")
	optVariables := getopt.BoolLong("variables", 'v', "Show variables")
	optBlocks := getopt.BoolLong("blocks", 'b', "Show blocks")
	optGetBind := getopt.StringLong("get-bind", 'g', "", "get binding where command or dispatcher contains given string")
	optOutput := getopt.StringLong("output-file", 'o', "", "File path to output file")
	optPath := getopt.StringLong("config-file", 'f', "", "path to config file, default is $HOME/.config/hypr/hyprland.conf")
	getopt.Parse()
	return &Flags{
		Help:       *optHelp,
		Json:       *optJson,
		Markdown:   *optMarkdown,
		Raw:        *optRaw,
		Version:    *optVersion,
		Variables:  *optVariables,
		Blocks:     *optBlocks,
		GetBind:    *optGetBind,
		ConfigPath: *optPath,
		Output:     *optOutput,
	}
}
