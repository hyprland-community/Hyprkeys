package properties

// config sections

// bool values need to be parsed from 0 or 1

type S_general struct {
	S_sensitivity             float64
	S_apply_sens_to_raw       bool
	S_main_mod                string
	S_border_size             int64
	S_no_border_on_floating   bool
	S_gaps_in                 int64
	S_gaps_out                int64
	S_col__active_border      string
	S_col__inactive_border    string
	S_cursor_inactive_timeout int64
	S_no_cursor_warps         bool
	S_layout                  string
}

type S_misc struct {
	S_disable_hyprland_logo     bool
	S_disable_splash_rendering  bool
	S_no_vfr                    bool
	S_damage_entire_on_snapshot bool
	S_mouse_move_enables_dpms   bool
	S_always_follow_on_dnd      bool
	S_layers_hog_keyboard_focus bool
	S_animate_manual_resizes    bool
	S_disable_autoreload        bool
	S_enable_swallow            bool
	S_swallow_regex             string
	S_focus_on_activate         bool
	S_no_direct_scanout         bool
}

type S_debug struct {
	S_log_damage      bool
	S_overlay         bool
	S_damage_blink    bool
	S_disable_logs    bool
	S_disable_time    bool
	S_damage_tracking int64
}

type S_decoration struct {
	S_rounding               int64
	S_blur                   bool
	S_blur_size              int64
	S_blur_passes            int64
	S_blur_ignore_opacity    bool
	S_blur_new_optimizations bool
	S_active_opacity         float64
	S_inactive_opacity       float64
	S_fullscreen_opacity     float64
	S_multisample_edges      bool
	S_no_blur_on_oversized   bool
	S_drop_shadow            bool
	S_shadow_range           int64
	S_shadow_render_power    int64
	S_shadow_ignore_window   bool
	S_shadow_offset          [2]float64
	S_col__shadow            string
	S_col__shadow_inactive   string
	S_dim_inactive           bool
	S_dim_strength           float64
}

type S_dwindle struct {
	S_pseudotile               bool
	S_col__group_border        string
	S_col__group_border_active string
	S_force_split              int64
	S_preserve_split           bool
	S_special_scale_factor     float64
	S_split_width_multiplier   float64
	S_no_gaps_when_only        bool
	S_use_active_for_splits    bool
}

type S_master struct {
	S_special_scale_factor float64
	S_new_is_master        bool
	S_new_on_top           bool
	S_no_gaps_when_only    bool
}

type S_animations struct {
	S_enabled                bool
	S_use_resize_transitions bool
}

type S_touchpad struct {
	S_natural_scroll          bool
	S_disable_while_typing    bool
	S_clickfinger_behavior    bool
	S_middle_button_emulation bool
	S_tap_to_click            bool //say vaxry to rename this shit to use underscores instead of dashes
	S_drag_lock               bool
	S_scroll_factor           float64
	S_transform               int64
	S_output                  string
}

type S_touchdevice struct {
	S_transform int64
	S_output    string
}

type S_input struct {
	S_sensitivity                 float64
	S_accel_profile               string
	S_kb_file                     string
	S_kb_layout                   string
	S_kb_variant                  string
	S_kb_options                  string
	S_kb_rules                    string
	S_kb_model                    string
	S_repeat_rate                 int64
	S_repeat_delay                int64
	S_natural_scroll              bool
	S_numlock_by_default          bool
	S_force_no_accel              bool
	S_float_switch_override_focus int64
	S_left_handed                 bool
	S_scroll_method               string
	S_touchpad                    *S_touchpad
	S_touchdevice                 *S_touchdevice
	S_follow_mouse                int64
	// TODO: per device config
}

type S_binds struct {
	S_pass_mouse_when_bound    bool
	S_scroll_event_delay       int64
	S_workspace_back_and_forth bool
	S_allow_workspace_cycles   bool
}

type S_gestures struct {
	S_workspace_swipe                    bool
	S_workspace_swipe_fingers            int64
	S_workspace_swipe_distance           int64
	S_workspace_swipe_invert             bool
	S_workspace_swipe_min_speed_to_force int64
	S_workspace_swipe_cancel_ratio       float64
	S_workspace_swipe_create_new         bool
	S_workspace_swipe_forever            bool
}

// config

type Config struct {
	Global     string
	General    *S_general
	Misc       *S_misc
	Debug      *S_debug
	Decoration *S_decoration
	Dwindle    *S_dwindle
	Master     *S_master
	Animations *S_animations
	Input      *S_input
	Binds      *S_binds
	Gestures   *S_gestures
}

func NewConf() Config {
	return Config{
		Global: "",
		General: &S_general{
			S_sensitivity:             1.0,
			S_apply_sens_to_raw:       false,
			S_main_mod:                "SUPER",
			S_border_size:             1,
			S_no_border_on_floating:   false,
			S_gaps_in:                 5,
			S_gaps_out:                20,
			S_col__active_border:      "0xffffffff",
			S_col__inactive_border:    "0xff444444",
			S_cursor_inactive_timeout: 0,
			S_no_cursor_warps:         false,
			S_layout:                  "dwindle",
		},
		Misc: &S_misc{
			S_disable_hyprland_logo:     false,
			S_disable_splash_rendering:  false,
			S_no_vfr:                    true,
			S_damage_entire_on_snapshot: false,
			S_mouse_move_enables_dpms:   false,
			S_always_follow_on_dnd:      true,
			S_layers_hog_keyboard_focus: true,
			S_animate_manual_resizes:    false,
			S_disable_autoreload:        false,
			S_enable_swallow:            false,
			S_swallow_regex:             "",
			S_focus_on_activate:         true,
			S_no_direct_scanout:         false,
		},
		Debug: &S_debug{
			S_overlay:         false,
			S_damage_blink:    false,
			S_disable_logs:    false,
			S_disable_time:    true,
			S_damage_tracking: 2,
		},
		Decoration: &S_decoration{
			S_rounding:               0,
			S_multisample_edges:      true,
			S_active_opacity:         1.0,
			S_inactive_opacity:       1.0,
			S_fullscreen_opacity:     1.0,
			S_blur:                   true,
			S_blur_size:              8,
			S_blur_passes:            1,
			S_blur_ignore_opacity:    false,
			S_blur_new_optimizations: false,
			S_drop_shadow:            true,
			S_shadow_range:           4,
			S_shadow_render_power:    3,
			S_shadow_ignore_window:   true,
			S_col__shadow:            "0xee1a1a1a",
			S_col__shadow_inactive:   "0xee1a1a1a",
			S_shadow_offset:          [2]float64{0, 0},
			S_dim_inactive:           false,
			S_dim_strength:           0.5,
		},
		Dwindle: &S_dwindle{
			S_pseudotile:               false,
			S_col__group_border:        "0x66777700",
			S_col__group_border_active: "0x66ffff00",
			S_force_split:              0,
			S_preserve_split:           false,
			S_special_scale_factor:     0.8,
			S_split_width_multiplier:   1.0,
			S_no_gaps_when_only:        false,
			S_use_active_for_splits:    true,
		},
		Master: &S_master{
			S_special_scale_factor: 0.8,
			S_new_is_master:        false,
			S_new_on_top:           false,
			S_no_gaps_when_only:    false,
		},
		Animations: &S_animations{
			S_enabled:                true,
			S_use_resize_transitions: false,
		},
		Input: &S_input{
			S_kb_layout:                   "us",
			S_kb_variant:                  "",
			S_kb_model:                    "",
			S_kb_options:                  "",
			S_kb_rules:                    "",
			S_kb_file:                     "",
			S_follow_mouse:                1,
			S_float_switch_override_focus: 1,
			S_repeat_rate:                 25,
			S_repeat_delay:                600,
			S_natural_scroll:              false,
			S_numlock_by_default:          false,
			S_force_no_accel:              false,
			S_sensitivity:                 0.0,
			S_left_handed:                 false,
			S_accel_profile:               "",
			S_scroll_method:               "",
			S_touchpad: &S_touchpad{
				S_disable_while_typing:    true,
				S_natural_scroll:          false,
				S_clickfinger_behavior:    false,
				S_middle_button_emulation: false,
				S_tap_to_click:            true,
				S_drag_lock:               false,
				S_scroll_factor:           1.0,
			},
			S_touchdevice: &S_touchdevice{
				S_transform: 0,
				S_output:    "",
			},
		},
		Binds: &S_binds{
			S_pass_mouse_when_bound:    false,
			S_scroll_event_delay:       300,
			S_workspace_back_and_forth: false,
			S_allow_workspace_cycles:   false,
		},
		Gestures: &S_gestures{
			S_workspace_swipe:                    false,
			S_workspace_swipe_fingers:            3,
			S_workspace_swipe_distance:           300,
			S_workspace_swipe_invert:             true,
			S_workspace_swipe_min_speed_to_force: 30,
			S_workspace_swipe_cancel_ratio:       0.5,
			S_workspace_swipe_create_new:         true,
			S_workspace_swipe_forever:            false,
		},
	}
}
