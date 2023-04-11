package main

import (
	"testing"

	// io/ioutil is deprecated, use io and os packages instead
	"notashelf.dev/hyprkeys/src/config"
	"notashelf.dev/hyprkeys/src/reader"
	"notashelf.dev/hyprkeys/src/writer"
)

func TestMarkdown(t *testing.T) {
	flags := config.Flags{}
	flags.ConfigPath = "test/hyprland.conf"
	flags.Binds = true
	flags.Output = "test/markdown.md"
	flags.Markdown = true
	configValues, err := reader.ReadHyprlandConfig(flags)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = writer.OutputConfig(configValues, flags)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestJson(t *testing.T) {
	flags := config.Flags{}
	flags.ConfigPath = "test/hyprland.conf"
	flags.Binds = true
	flags.Output = "test/out.json"
	flags.Json = true
	configValues, err := reader.ReadHyprlandConfig(flags)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = writer.OutputConfig(configValues, flags)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRaw(t *testing.T) {
	flags := config.Flags{}
	flags.ConfigPath = "test/hyprland.conf"
	flags.Binds = true
	flags.Output = "test/out"
	flags.Raw = true
	configValues, err := reader.ReadHyprlandConfig(flags)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = writer.OutputConfig(configValues, flags)
	if err != nil {
		t.Errorf(err.Error())
	}
}
