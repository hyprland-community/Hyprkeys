package config

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
