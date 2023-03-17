package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"notashelf.dev/hyprkeys/src/config"
	"notashelf.dev/hyprkeys/src/ctl"
	"notashelf.dev/hyprkeys/src/reader"
	"notashelf.dev/hyprkeys/src/writer"
)

var conf config.Flags

var rootCmd = &cobra.Command{
	Use:   "hyprkeys",
	Short: "A simple, scriptable keybind retrieval utility for Hyprand",
	Run:   run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&conf.Json, "json", "j", true, "Output in json")
	rootCmd.Flags().BoolVarP(&conf.Markdown, "markdown", "m", false, "Output in markdown")
	rootCmd.Flags().BoolVarP(&conf.Raw, "raw", "r", false, "Output in plain text")
	rootCmd.MarkFlagsMutuallyExclusive("json", "markdown", "raw")

	rootCmd.Flags().StringVarP(&conf.ConfigPath, "config-file", "c", "", "path to your hyprland config")
	rootCmd.Flags().StringVarP(&conf.Output, "output", "o", "", "path to output file")

	rootCmd.Flags().BoolVarP(&conf.Binds, "binds", "b", false, "output binds")
	rootCmd.Flags().BoolVarP(&conf.Ctl, "from-ctl", "t", false, "get binds from ctl")
	rootCmd.Flags().BoolVarP(&conf.AutoStart, "auto-start", "a", false, "Show autostarting programs")
	rootCmd.Flags().BoolVarP(&conf.Variables, "variables", "v", false, "Show variables")
	rootCmd.Flags().BoolVarP(&conf.Keywords, "keywords", "k", false, "Show keywords")
	rootCmd.Flags().BoolVarP(&conf.Comments, "comments", "l", false, "Show comments in output")

	rootCmd.Flags().StringP("filter-binds", "f", "", "get binding where command or dispatcher contains given string")
}

func run(cmd *cobra.Command, args []string) {
	if conf.ConfigPath == "" {
		confDir, err := os.UserConfigDir()
		if err != nil {
			homeDir, _ := os.UserHomeDir()
			confDir = filepath.Join(homeDir, ".config")
		}
		conf.ConfigPath = filepath.Join(confDir, "/hypr/hyprland.conf")
	}

	if conf.Ctl {
		var err error
		binds, err := ctl.BindsFromCtl()
		if err != nil {
			panic(err)
		}
		writer.OutputCtl(binds, conf)
	}

	configValues, err := reader.ReadHyprlandConfig(conf)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = writer.OutputConfig(configValues, conf)
	if err != nil {
		fmt.Println(err.Error())
	}
}
