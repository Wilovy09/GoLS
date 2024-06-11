# GoLS

We deserve a better and faster LS command.

> [!NOTE]
> PR is accepted to improve the code

> [!WARNING]
> This project is under construction, it has visual errors more than anything else.

## Preview

```sh
$ ls     
📁.
├──📄README.md
├──📄go.mod
└──📄main.go
$ ls -a
📁.
├──📁.git
├──📄README.md
├──📄go.mod
└──📄main.go
```

## Installation

> [!NOTE]
> You need to have [Go](https://go.dev/dl/) installed on your machine.

Its to simple to install, just open your terminal and run the following command.

```bash
go install github.com/Wilovy09/GoLS@latest
```

Find the `gols` binary in your `$GOPATH/bin` directory and create an alias in your `.bashrc` or `.zshrc` file.

```bash
# [ GoLS better ls ]
alias ls='~/go/bin/GoLS'
```

And you're done.

### Compile

Clone the repository and compile the code.

```bash
go build -o GoLS -buildvcs=false
```

The `gols` binary will be moved to the root directory `/home/USER`.

You need to create a `alias` in your `.bashrc` or `.zshrc` file.

```bash
# [ gols better ls ]
alias ls='~/gols'
```

## Flags

| Flag   | Description                 | Values  | Default |
|--------|-----------------------------|---------|-------- |
| -a     | Show hidden files           | boolean | false   |
| -depth | How deep the search will be | int     | 0       |
