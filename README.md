# ![ ](whatever_your_browser_does_for_image_not_found.jpeg) notfoundle
A command not found handler like wordle:
```
% probet
Did you mean: ⬜🟨🟨🟨🟨🟩
```

It's not deterministic so don't try to solve them or you will lose your mind.

## Installation
You need [Go](https://go.dev) and `$GOPATH/bin` (or `~/go/bin`) in your PATH.
```
go install github.com/brackendawson/notfoundle@latest
```

### Zsh
```
echo 'command_not_found_handler() { notfoundle "$1" }' >>~/.zshrc
```

### fish shell
```
echo -e "function fish_command_not_found\n    notfoundle \$argv[1]\nend" >> ~/.config/fish/functions/default/not_found.fish
```
