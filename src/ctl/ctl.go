package ctl

import (
	"encoding/json"
	"os/exec"
	"strings"
)

var MOD_VALS map[int]string

func init() {
	MOD_VALS = make(map[int]string)
	MOD_VALS[1] = "SHIFT"
	MOD_VALS[2] = "CAPS"
	MOD_VALS[4] = "CTRL"
	MOD_VALS[8] = "ALT"
	MOD_VALS[16] = "MOD2"
	MOD_VALS[32] = "MOD3"
	MOD_VALS[64] = "SUPER"
	MOD_VALS[128] = "MOD5"
}

func modMaskToString(modMask int) string {
	curVal := 7
	mods := []string{}
	for modMask > 0 {
		modVal := 1 << curVal
		if modMask >= modVal {
			modMask -= modVal
			mods = append(mods, MOD_VALS[1<<curVal])
		}
		curVal--
	}
	return strings.Join(mods, " ")
}

type Binds []struct {
	Locked     bool   `json:"locked"`
	Mouse      bool   `json:"mouse"`
	Release    bool   `json:"release"`
	Repeat     bool   `json:"repeat"`
	Modmask    int    `json:"modmask"`
	Mods       string `json:"mods"`
	Submap     string `json:"submap"`
	Key        string `json:"key"`
	Keycode    int    `json:"keycode"`
	Dispatcher string `json:"dispatcher"`
	Arg        string `json:"arg"`
}

func BindsFromCtl() (Binds, error) {
	out, err := exec.Command("hyprctl", "binds", "-j").Output()
	if err != nil {
		return nil, err
	}
	binds := Binds{}
	err = json.Unmarshal(out, &binds)
	if err != nil {
		return nil, err
	}
	for idx, bind := range binds {
		binds[idx].Mods = modMaskToString(bind.Modmask)
	}
	return binds, err
}
