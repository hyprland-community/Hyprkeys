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
	Keywords  bool
	Comments  bool
	Ctl       bool
	Binds     bool

	FilterBinds string
	Output      string
	ConfigPath  string
}

func ReadFlags() *Flags {
	optHelp := getopt.BoolLong("help", 'h', "Show this help menu")
	optVersion := getopt.BoolLong("version", 'V', "Show the version number")

	optPath := getopt.StringLong("config-file", 'c', "", "path to config file, default is $HOME/.config/hypr/hyprland.conf")

	optBinds := getopt.BoolLong("binds", 'b', "output binds")
	optCtl := getopt.BoolLong("from-ctl", 't', "get binds from ctl")
	optAutoStart := getopt.BoolLong("auto-start", 'a', "Show autostarting programs")
	optVariables := getopt.BoolLong("variables", 'v', "Show variables")
	optKeywords := getopt.BoolLong("keywords", 'k', "Show keywords")
	optComments := getopt.BoolLong("comments", 'l', "Show comments in output")
	optFilterBinds := getopt.StringLong("filter-binds", 'f', "", "get binding where command or dispatcher contains given string use * for all")

	optJson := getopt.BoolLong("json", 'j', "Return settigns as json")
	optMarkdown := getopt.BoolLong("markdown", 'm', "Print the binds as a markdown table")
	optRaw := getopt.BoolLong("raw", 'r', "Print text as is, without making it pretty")

	optOutput := getopt.StringLong("output-file", 'o', "", "File path to output file")

	getopt.Parse()

	return &Flags{
		Help:        *optHelp,
		Ctl:         *optCtl,
		Json:        *optJson,
		Markdown:    *optMarkdown,
		Raw:         *optRaw,
		Comments:    *optComments,
		Version:     *optVersion,
		Variables:   *optVariables,
		Binds:       *optBinds,
		FilterBinds: *optFilterBinds,
		AutoStart:   *optAutoStart,
		ConfigPath:  *optPath,
		Keywords:    *optKeywords,
		Output:      *optOutput,
	}
}
