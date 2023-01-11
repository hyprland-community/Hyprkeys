
| Keybind | Locked | Mouse | Release | Repeat | Submap | Dispatcher | Command |
|---------|--------|-------|---------|--------|--------|------------|---------|
| <kbd>SUPER D</kbd> | false | false | false | false |  |exec | rofi -modi combi -combi-modi drun,window,ssh -show combi | 
| <kbd>SUPER SHIFT V</kbd> | false | false | false | false |  |exec | clipman pick -t rofi | 
| <kbd>SUPER SHIFT K</kbd> | false | false | false | false |  |exec | rofi -show keybinds -modi keybinds:/home/abs3nt/Scripts/keybinds.sh | 
| <kbd>SUPER Return</kbd> | false | false | false | false |  |exec | kitty | 
| <kbd>SUPER BACKSPACE</kbd> | false | false | false | false |  |exec | firefox-nightly | 
| <kbd>SUPER H</kbd> | false | false | false | false |  |exec | pcmanfm | 
| <kbd>SUPER ALT CTRL P</kbd> | false | false | false | false |  |exec | gospt toggleplay && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Toggle Play' '$OUTPUT'; done | 
| <kbd>SUPER ALT CTRL Y</kbd> | false | false | false | false |  |exec | gospt link \| wl-copy && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Copied Song Link' '$OUTPUT'; done | 
| <kbd>SUPER ALT CTRL N</kbd> | false | false | false | false |  |exec | gospt skip && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Now Playing' '$OUTPUT'; done | 
| <kbd>SUPER ALT CTRL H</kbd> | false | false | false | false |  |exec | gospt previous && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Now Playing' '$OUTPUT'; done | 
| <kbd>SUPER ALT CTRL U</kbd> | false | false | false | false |  |exec | gospt unlike && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Unliked Song' '$OUTPUT'; done | 
| <kbd>SUPER ALT CTRL L</kbd> | false | false | false | false |  |exec | gospt like && gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Liked Song' '$OUTPUT'; done | 
| <kbd>SUPER ALT CTRL M</kbd> | false | false | false | false |  |exec | gospt radio &&  gospt nowplaying \| while read OUTPUT; do notify-send -t 4000 'Radio Started' '$OUTPUT'; done | 
| <kbd>SUPER SHIFT S</kbd> | false | false | false | false |  |exec | slurp -d \| grim -g - - \| wl-copy | 
| <kbd>SUPER ALT CTRL S</kbd> | false | false | false | false |  |exec | grim -o DP-1 - \| wl-copy | 
| <kbd>SUPER CTRL SHIFT S</kbd> | false | false | false | false |  |exec | wf-recorder -g '$(slurp)' --audio -f ~/Videos/Screenrecord/record.mp4 | 
| <kbd>SUPER CTRL SHIFT A</kbd> | false | false | false | false |  |exec | killall -s SIGINT wf-recorder | 
| <kbd>SUPER SHIFT W</kbd> | false | false | false | false |  |exec | ~/Scripts/set_wallpaper.sh | 
| <kbd>SUPER ALT L</kbd> | false | false | false | false |  |exec | ~/Scripts/lockscreen.sh | 
| <kbd>SUPER V</kbd> | false | false | false | false |  |togglefloating |  | 
| <kbd>SUPER F</kbd> | false | false | false | false |  |fullscreen |  | 
| <kbd>SUPER SHIFT F</kbd> | false | false | false | false |  |fakefullscreen |  | 
| <kbd>SUPER P</kbd> | false | false | false | false |  |pseudo |  | 
| <kbd>SUPER J</kbd> | false | false | false | false |  |togglesplit |  | 
| <kbd>SUPER SHIFT Q</kbd> | false | false | false | false |  |killactive |  | 
| <kbd>SUPER left</kbd> | false | false | false | false |  |movefocus | l | 
| <kbd>SUPER right</kbd> | false | false | false | false |  |movefocus | r | 
| <kbd>SUPER up</kbd> | false | false | false | false |  |movefocus | u | 
| <kbd>SUPER down</kbd> | false | false | false | false |  |movefocus | d | 
| <kbd>SUPER 1</kbd> | false | false | false | false |  |workspace | 1 | 
| <kbd>SUPER 2</kbd> | false | false | false | false |  |workspace | 2 | 
| <kbd>SUPER 3</kbd> | false | false | false | false |  |workspace | 3 | 
| <kbd>SUPER 4</kbd> | false | false | false | false |  |workspace | 4 | 
| <kbd>SUPER 5</kbd> | false | false | false | false |  |workspace | 5 | 
| <kbd>SUPER 6</kbd> | false | false | false | false |  |workspace | 6 | 
| <kbd>SUPER 7</kbd> | false | false | false | false |  |workspace | 7 | 
| <kbd>SUPER 8</kbd> | false | false | false | false |  |workspace | 8 | 
| <kbd>SUPER 9</kbd> | false | false | false | false |  |workspace | 9 | 
| <kbd>SUPER 0</kbd> | false | false | false | false |  |workspace | 10 | 
| <kbd>SUPER SHIFT 1</kbd> | false | false | false | false |  |movetoworkspacesilent | 1 | 
| <kbd>SUPER SHIFT 2</kbd> | false | false | false | false |  |movetoworkspacesilent | 2 | 
| <kbd>SUPER SHIFT 3</kbd> | false | false | false | false |  |movetoworkspacesilent | 3 | 
| <kbd>SUPER SHIFT 4</kbd> | false | false | false | false |  |movetoworkspacesilent | 4 | 
| <kbd>SUPER SHIFT 5</kbd> | false | false | false | false |  |movetoworkspacesilent | 5 | 
| <kbd>SUPER SHIFT 6</kbd> | false | false | false | false |  |movetoworkspacesilent | 6 | 
| <kbd>SUPER SHIFT 7</kbd> | false | false | false | false |  |movetoworkspacesilent | 7 | 
| <kbd>SUPER SHIFT 8</kbd> | false | false | false | false |  |movetoworkspacesilent | 8 | 
| <kbd>SUPER SHIFT 9</kbd> | false | false | false | false |  |movetoworkspacesilent | 9 | 
| <kbd>SUPER SHIFT 0</kbd> | false | false | false | false |  |movetoworkspacesilent | 10 | 
| <kbd>SUPER mouse_down</kbd> | false | false | false | false |  |workspace | e+1 | 
| <kbd>SUPER mouse_up</kbd> | false | false | false | false |  |workspace | e-1 | 
| <kbd>SUPER mouse:272</kbd> | false | true | false | false |  |mouse | movewindow | 
| <kbd>SUPER mouse:273</kbd> | false | true | false | false |  |mouse | resizewindow | 
| <kbd>SUPER SHIFT M</kbd> | false | false | false | false |  |exit |  | 
