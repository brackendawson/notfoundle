# notfoundle
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
