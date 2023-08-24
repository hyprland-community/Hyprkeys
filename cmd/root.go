package cmd

import (
	"fmt"
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
	RunE:  run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&conf.Json, "json", "j", false, "Output in json")
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

	rootCmd.Flags().StringVarP(&conf.FilterBinds,"filter-binds", "f", "", "Filter binds where command or dispatcher contains given string")
}

func run(cmd *cobra.Command, args []string) error {
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
		err = writer.OutputCtl(binds, conf)
		if err != nil {
			return err
		}
	}

	configValues, err := reader.ReadHyprlandConfig(conf)
	if err != nil {
		return err
	}

	err = writer.OutputConfig(configValues, conf)
	if err != nil {
		return err
	}
	return nil
}
