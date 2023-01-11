
| Keybind | Dispatcher | Command | Comments |
|---------|------------|---------|----------|
| <kbd>SUPER D</kbd> | exec | rofi -modi combi -combi-modi drun,window,ssh -show combi |  |
| <kbd>SUPER SHIFT V</kbd> | exec | clipman pick -t rofi |  |
| <kbd>SUPER SHIFT K</kbd> | exec | rofi -show keybinds -modi keybinds:/home/abs3nt/Scripts/keybinds.sh |  |
| <kbd>SUPER Return</kbd> | exec | kitty |  |
| <kbd>SUPER BACKSPACE</kbd> | exec | firefox |  |
| <kbd>SUPER H</kbd> | exec | pcmanfm  | WOOO |
| <kbd>SUPER CONTROL ALT P</kbd> | exec | gospt toggleplay && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Toggle Play" "$OUTPUT"; done |  |
| <kbd>SUPER CONTROL ALT Y</kbd> | exec | gospt link \| wl-copy && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Copied Song Link" "$OUTPUT"; done |  |
| <kbd>SUPER CONTROL ALT N</kbd> | exec | gospt skip && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Now Playing" "$OUTPUT"; done |  |
| <kbd>SUPER CONTROL ALT H</kbd> | exec | gospt previous && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Now Playing" "$OUTPUT"; done |  |
| <kbd>SUPER CONTROL ALT U</kbd> | exec | gospt unlike && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Unliked Song" "$OUTPUT"; done |  |
| <kbd>SUPER CONTROL ALT L</kbd> | exec | gospt like && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Liked Song" "$OUTPUT"; done |  |
| <kbd>SUPER CONTROL ALT M</kbd> | exec | gospt radio &&  gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 "Radio Started" "$OUTPUT"; done |  |
| <kbd>SUPER SHIFT S</kbd> | exec | slurp -d \| grim -g - - \| wl-copy |  |
| <kbd>SUPER CONTROL ALT S</kbd> | exec | grim -o DP-1 - \| wl-copy |  |
| <kbd>SUPER SHIFT CONTROL S</kbd> | exec | wf-recorder -g "$(slurp)" --audio -f ~/Videos/Screenrecord/record.mp4 |  |
| <kbd>SUPER SHIFT CONTROL A</kbd> | exec | killall -s SIGINT wf-recorder |  |
| <kbd>SUPER SHIFT W</kbd> | exec | ~/Scripts/set_wallpaper.sh |  |
| <kbd>SUPER ALT L</kbd> | exec | ~/Scripts/lockscreen.sh |  |
| <kbd>SUPER V</kbd> | togglefloating |  |  |
| <kbd>SUPER F</kbd> | fullscreen |  |  |
| <kbd>SUPER SHIFT F</kbd> | fakefullscreen |  |  |
| <kbd>SUPER P</kbd> | pseudo |  | dwindle |
| <kbd>SUPER J</kbd> | togglesplit |  | dwindle |
| <kbd>SUPER SHIFT Q</kbd> | killactive |  |  |
| <kbd>SUPER left</kbd> | movefocus | l |  |
| <kbd>SUPER right</kbd> | movefocus | r |  |
| <kbd>SUPER up</kbd> | movefocus | u |  |
| <kbd>SUPER down</kbd> | movefocus | d |  |
| <kbd>SUPER 1</kbd> | workspace | 1 |  |
| <kbd>SUPER 2</kbd> | workspace | 2 |  |
| <kbd>SUPER 3</kbd> | workspace | 3 |  |
| <kbd>SUPER 4</kbd> | workspace | 4 |  |
| <kbd>SUPER 5</kbd> | workspace | 5 |  |
| <kbd>SUPER 6</kbd> | workspace | 6 |  |
| <kbd>SUPER 7</kbd> | workspace | 7 |  |
| <kbd>SUPER 8</kbd> | workspace | 8 |  |
| <kbd>SUPER 9</kbd> | workspace | 9 |  |
| <kbd>SUPER 0</kbd> | workspace | 10 |  |
| <kbd>SUPER SHIFT 1</kbd> | movetoworkspacesilent | 1 |  |
| <kbd>SUPER SHIFT 2</kbd> | movetoworkspacesilent | 2 |  |
| <kbd>SUPER SHIFT 3</kbd> | movetoworkspacesilent | 3 |  |
| <kbd>SUPER SHIFT 4</kbd> | movetoworkspacesilent | 4 |  |
| <kbd>SUPER SHIFT 5</kbd> | movetoworkspacesilent | 5 |  |
| <kbd>SUPER SHIFT 6</kbd> | movetoworkspacesilent | 6 |  |
| <kbd>SUPER SHIFT 7</kbd> | movetoworkspacesilent | 7 |  |
| <kbd>SUPER SHIFT 8</kbd> | movetoworkspacesilent | 8 |  |
| <kbd>SUPER SHIFT 9</kbd> | movetoworkspacesilent | 9 |  |
| <kbd>SUPER SHIFT 0</kbd> | movetoworkspacesilent | 10 |  |
| <kbd>SUPER mouse_down</kbd> | workspace | e+1 |  |
| <kbd>SUPER mouse_up</kbd> | workspace | e-1 |  |
| <kbd>SUPER mouse:272</kbd> | movewindow |  |  |
| <kbd>SUPER mouse:273</kbd> | resizewindow |  |  |
| <kbd>SUPER SHIFT M</kbd> | exit |  |  |
