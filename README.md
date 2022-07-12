# Game of life

Golang implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) which can be used in combination with
[i3lock](https://github.com/i3/i3lock) and [Alacritty](https://github.com/alacritty/alacritty) as lockscreen.

*Note: Other setups are possible, the game itself will resize to fit the terminal*

This specific setup uses two scripts.

### lock.sh

---

1. Run [Alacritty](https://github.com/alacritty/alacritty) in fullscreen and execute `run.sh` in background (save PID)
2. Lock the screen with [i3lock](https://github.com/i3/i3lock) and set the backgorund to transparent, also don't fork the process
so the rest of the script runs after unlock
3. After unclock kill the background terminal

