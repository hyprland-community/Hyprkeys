package flags

import (
	"github.com/pborman/getopt"
)

type Flags struct {
	Help      bool
	Json      bool
	Markdown  bool
	Raw       bool
	Version   bool
	Variables bool
	AutoStart bool

	FilterBinds string
	Output      string
	ConfigPath  string
}

func ReadFlags() *Flags {
	optHelp := getopt.BoolLong("help", 'h', "Show this help menu")
	optVersion := getopt.BoolLong("version", 'V', "Show the version number")

	optPath := getopt.StringLong("config-file", 'f', "", "path to config file, default is $HOME/.config/hypr/hyprland.conf")

	optAutoStart := getopt.BoolLong("auto-start", 'a', "Show autostarting programs")
	optVariables := getopt.BoolLong("variables", 'v', "Show variables")
	optFilterBinds := getopt.StringLong("filter-binds", 'b', "", "get binding where command or dispatcher contains given string use * for all")

	optJson := getopt.BoolLong("json", 'j', "Return settigns as json")
	optMarkdown := getopt.BoolLong("markdown", 'm', "Print the binds as a markdown table")
	optRaw := getopt.BoolLong("raw", 'r', "Print text as is, without making it pretty")

	optOutput := getopt.StringLong("output-file", 'o', "", "File path to output file")

	getopt.Parse()

	return &Flags{
		Help:        *optHelp,
		Json:        *optJson,
		Markdown:    *optMarkdown,
		Raw:         *optRaw,
		Version:     *optVersion,
		Variables:   *optVariables,
		FilterBinds: *optFilterBinds,
		AutoStart:   *optAutoStart,
		ConfigPath:  *optPath,
		Output:      *optOutput,
	}
}
