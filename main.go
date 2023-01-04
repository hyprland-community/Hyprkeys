package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/pborman/getopt"

	// io/ioutil is deprecated, use io and os packages instead
	"notashelf.dev/hyprkeys/flags"
	"notashelf.dev/hyprkeys/reader"
)

func main() {
	flags := flags.ReadFlags()

	if flags.Version {
		version := "unknown"
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
		fmt.Println("version:", version)
		return
	}

	if flags.Help {
		getopt.Usage()
		return
	}

	if flags.ConfigPath == "" {
		flags.ConfigPath = filepath.Join(os.Getenv("HOME"), ".config/hypr/hyprland.conf")
	}
	if flags.Test {
		flags.ConfigPath = "test/hyprland.conf"
	}

	configValues, err := reader.ReadHyprlandConfig(flags.ConfigPath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if flags.Raw {
		for _, bind := range configValues.KeyboardBinds {
			fmt.Println(bind)
		}
		for _, bind := range configValues.MouseBinds {
			fmt.Println(bind)
		}
		for _, val := range configValues.Settings {
			fmt.Println(val.Name, "{")
			for setting, value := range val.Settings {
				fmt.Println("\t", setting, "=", value)
			}
			for _, set := range val.SubCategories {
				fmt.Println("\t", set.Name, "{")
				for setting, value := range set.Settings {
					fmt.Println("\t\t", setting, "=", value)
				}
				fmt.Println("\t", "}")
			}
			fmt.Println("}")
		}
	}
	if flags.Json {

		out, err := json.MarshalIndent(configValues, "", " ")
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(string(out))
	}
}
