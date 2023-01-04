package main

import (
	"testing"

	// io/ioutil is deprecated, use io and os packages instead
	"notashelf.dev/hyprkeys/flags"
	"notashelf.dev/hyprkeys/reader"
)

func TestMarkdown(t *testing.T) {
	flags := &flags.Flags{}
	flags.ConfigPath = "test/hyprland.conf"
	flags.Output = "test/markdown.md"
	flags.Markdown = true
	configValues, err := reader.ReadHyprlandConfig(flags.ConfigPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = markdownHandler(configValues, flags)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestJson(t *testing.T) {
	flags := &flags.Flags{}
	flags.ConfigPath = "test/hyprland.conf"
	flags.Output = "test/out.json"
	flags.Json = true
	configValues, err := reader.ReadHyprlandConfig(flags.ConfigPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = jsonHandler(configValues, flags)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRaw(t *testing.T) {
	flags := &flags.Flags{}
	flags.ConfigPath = "test/hyprland.conf"
	flags.Output = "test/out"
	flags.Json = true
	configValues, err := reader.ReadHyprlandConfig(flags.ConfigPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = rawHandler(configValues, flags)
	if err != nil {
		t.Errorf(err.Error())
	}
}
