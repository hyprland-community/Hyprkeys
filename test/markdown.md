# Sample Markdown

This is an example output for `hyprkeys --markdown`. Syntax is last updated on 07/11/2022.
May be subjected to change.

## Known Issues
The "bind =" prefix from the struct currently gets picked up by the parser.

## Keys

| Keybind | Dispatcher | Command |
|---------|------------|---------|
| <kbd>bind = $mainMod + Q</kbd> | | exec, kitty |
| <kbd>bind = $mainMod + C</kbd> | | killactive, |
| <kbd>bind = $mainMod + M</kbd> | | exit, |
| <kbd>bind = $mainMod + E</kbd> | | exec, dolphin |
| <kbd>bind = $mainMod + V</kbd> | | togglefloating, |
| <kbd>bind = $mainMod + R</kbd> | | exec, wofi --show drun |
| <kbd>bind = $mainMod + P</kbd> | | pseudo, # dwindle |
| <kbd>bind = $mainMod + J</kbd> | | togglesplit, # dwindle |
| <kbd>bind = $mainMod + left</kbd> | | movefocus, l |
| <kbd>bind = $mainMod + right</kbd> | | movefocus, r |
| <kbd>bind = $mainMod + up</kbd> | | movefocus, u |
| <kbd>bind = $mainMod + down</kbd> | | movefocus, d |
| <kbd>bind = $mainMod + 1</kbd> | | workspace, 1 |
| <kbd>bind = $mainMod + 2</kbd> | | workspace, 2 |
| <kbd>bind = $mainMod + 3</kbd> | | workspace, 3 |
| <kbd>bind = $mainMod + 4</kbd> | | workspace, 4 |
| <kbd>bind = $mainMod + 5</kbd> | | workspace, 5 |
| <kbd>bind = $mainMod + 6</kbd> | | workspace, 6 |
| <kbd>bind = $mainMod + 7</kbd> | | workspace, 7 |
| <kbd>bind = $mainMod + 8</kbd> | | workspace, 8 |
| <kbd>bind = $mainMod + 9</kbd> | | workspace, 9 |
| <kbd>bind = $mainMod + 0</kbd> | | workspace, 10 |
| <kbd>bind = $mainMod SHIFT + 1</kbd> | | movetoworkspace, 1 |
| <kbd>bind = $mainMod SHIFT + 2</kbd> | | movetoworkspace, 2 |
| <kbd>bind = $mainMod SHIFT + 3</kbd> | | movetoworkspace, 3 |
| <kbd>bind = $mainMod SHIFT + 4</kbd> | | movetoworkspace, 4 |
| <kbd>bind = $mainMod SHIFT + 5</kbd> | | movetoworkspace, 5 |
| <kbd>bind = $mainMod SHIFT + 6</kbd> | | movetoworkspace, 6 |
| <kbd>bind = $mainMod SHIFT + 7</kbd> | | movetoworkspace, 7 |
| <kbd>bind = $mainMod SHIFT + 8</kbd> | | movetoworkspace, 8 |
| <kbd>bind = $mainMod SHIFT + 9</kbd> | | movetoworkspace, 9 |
| <kbd>bind = $mainMod SHIFT + 0</kbd> | | movetoworkspace, 10 |
| <kbd>bind = $mainMod + mouse_down</kbd> | | workspace, e+1 |
| <kbd>bind = $mainMod + mouse_up</kbd> | | workspace, e-1 |
| <kbd>bindm = $mainMod + mouse:272</kbd> | | movewindow |
| <kbd>bindm = $mainMod + mouse:273</kbd> | | resizewindow |