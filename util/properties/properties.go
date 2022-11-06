package properties

type SConfigValue struct {
	int64_t  int64
	float    float64
	strValue string
	vecValue [2]int // is this how u make a vector, idk
	set      bool   // used for device configs
}

type SMonitorRule struct {
	name             string
	resolution       [2]int16
	offset           [2]int16
	scale            float32
	refreshrate      float32
	defaultworkspace string
	disabled         bool
	mirrorOf         string
	enable10bit      bool
}

type SMonitorAdditionalReservedArea struct {
	top    int16
	bottom int16
	left   int16
	right  int16
}

type SWindowRule struct {
	szRule    string
	szValue   string
	v2        bool
	szTitle   string
	szClass   string
	bX11      int
	bFloating int
}

type SAnimationPropertyConfig struct {
	overriden       bool
	internalBezier  string
	internalStyle   string
	internalSpeed   float64
	internalEnabled int

	// TODO: AnimationPropertyConfig* pValues = nullptr;
	// TODO: SAnimationPropertyConfig* pParentAnimation = nullptr;
}

// config sections

// bool values need to be parsed from 0 or 1

type General struct {
	max_fps                 int
	sensitivity             float64
	apply_sens_to_raw       bool
	main_mod                string
	main_mod_internal       int // prolly dont need this
	border_size             int
	no_border_on_floating   bool
	gaps_in                 int
	gaps_out                int
	col__active_border      string // hex how
	col__inactive_border    string
	cursor_inactive_timeout int
	no_cursor_warps         bool
	layout                  string // dwindle or master

}

type Misc struct {
	disable_hyprland_logo     bool
	disable_splash_rendering  bool
	no_vfr                    bool
	damage_entire_on_snapshot bool
	mouse_move_enables_dpms   bool
	always_follow_on_dnd      bool
	layers_hog_keyboard_focus bool
	animate_manual_resizes    bool
	disable_autoreload        bool
	enable_swallow            bool
	swallow_regex             string
	focus_on_activate         bool
}

type Debug struct {
	_int            int // remove underscore before int
	log_damage      bool
	overlay         bool
	damage_blink    bool
	disable_logs    bool
	disable_time    bool
	damage_tracking int
}

type Decoration struct {
	rounding               int
	blur                   bool
	blur_size              int
	blur_passes            int
	blur_ignore_opacity    bool
	blur_new_optimizations bool
	active_opacity         float64
	inactive_opacity       float64
	fullscreen_opacity     float64
	multisample_edges      bool
	no_blur_on_oversized   bool
	drop_shadow            bool
	shadow_range           int
	shadow_render_power    int
	shadow_ignore_window   int
	shadow_offset          [2]int32
	col__shadow            string
	col__shadow_inactive   string
	dim_inactive           bool
	dim_strength           float64
}

type Dwindle struct {
	pseudotile               bool
	col__group_border        string
	col__group_border_active string
	force_split              bool
	preserve_split           bool
	special_scale_factor     float64
	split_width_multiplier   float64
	no_gaps_when_only        bool
	use_active_for_splits    bool
}

type Master struct {
	special_scale_factor float64
	new_is_master        bool
	new_on_top           bool
	no_gaps_when_only    bool
}

type Animations struct {
	enabled bool
	speed   float64
	curve   string // curve name

	// move these into a diff structs?
	windows_style string
	windows_curve string
	windows_speed float64
	windows       bool
	//
	borders_style string
	borders_curve string
	borders_speed float64
	borders       bool
	//
	fadein_style string
	fadein_curve string
	fadein_speed float64
	fadein       bool
	//
	workspaces_style string
	workspaces_curve string
	workspaces_speed float64
	workspaces       bool
}

type Touchpad struct {
	natural_scroll          bool
	disable_while_typing    bool
	clickfinger_behavior    bool
	middle_button_emulation bool
	// `tap-to-click` say vaxry to rename this shit to use _
	drag_lock     bool
	scroll_factor float64
	transform     int
	output        string
}

type Input struct {
	sensitivity                 float64
	accel_profile               string
	kb_file                     string
	kb_layout                   string
	kb_variant                  string
	kb_options                  string
	kb_rules                    string
	kb_model                    string
	repeat_rate                 int
	repeat_delay                int
	natural_scroll              bool
	numlock_by_default          bool
	force_no_accel              bool
	float_switch_override_focus bool
	left_handed                 bool
	scroll_method               string
	touchpad                    *Touchpad
	follow_mouse                bool
}

type Binds struct {
	pass_mouse_when_bound    bool
	scroll_event_delay       int
	workspace_back_and_forth bool
	allow_workspace_cycles   bool
}

type Gestures struct {
	workspace_swipe                    bool
	workspace_swipe_fingers            int
	workspace_swipe_distance           int
	workspace_swipe_invert             bool
	workspace_swipe_min_speed_to_force bool
	workspace_swipe_cancel_ratio       float64
	workspace_swipe_create_new         bool
	workspace_swipe_forever            bool
}

// config

type Config struct {
	autogenerated bool
	general       *General
	misc          *Misc
	debug         *Debug
	decoration    *Decoration
	dwindle       *Dwindle
	master        *Master
	animations    *Animations
	input         *Input
	binds         *Binds
	gestures      *Gestures
}
