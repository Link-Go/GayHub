## bcicen/ctop

Top-like interface for container metrics
容器指标终端监控

git: https://github.com/bcicen/ctop

指标展示类似 docker stats 指令，ctop 对其做了汇总操作

在指标看板上，通过键盘操作，与容器进行交互操作
有些指令并不是所有容器都可以操作。如：e。还是得通过 docker exec 实现操作

### Keybindings

|           Key            | Action                                                     |
| :----------------------: | ---------------------------------------------------------- |
| <kbd>&lt;ENTER&gt;</kbd> | Open container menu                                        |
|       <kbd>a</kbd>       | Toggle display of all (running and non-running) containers |
|       <kbd>f</kbd>       | Filter displayed containers (`esc` to clear when open)     |
|       <kbd>H</kbd>       | Toggle ctop header                                         |
|       <kbd>h</kbd>       | Open help dialog                                           |
|       <kbd>s</kbd>       | Select container sort field                                |
|       <kbd>r</kbd>       | Reverse container sort order                               |
|       <kbd>o</kbd>       | Open single view                                           |
|       <kbd>l</kbd>       | View container logs (`t` to toggle timestamp when open)    |
|       <kbd>e</kbd>       | Exec Shell                                                 |
|       <kbd>c</kbd>       | Configure columns                                          |
|       <kbd>S</kbd>       | Save current configuration to file                         |
|       <kbd>q</kbd>       | Quit ctop                                                  |