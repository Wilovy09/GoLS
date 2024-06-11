# GoLS

A better LS command.

> [!NOTE]
> PR is accepted to improve the code

## Preview

```sh
alias ls='~/gols'
```

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

Clone the repository and compile the code.

```bash
go build -o gols -buildvcs=false
```

The `gols` binary will be moved to the root directory `/home/USER`.

You need to create a `alias` in your `.bashrc` or `.zshrc` file.

```bash
# [ gols better ls ]
alias ls='~/gols'
```

> [!WARNING]
> This project is under construction, it has visual errors more than anything else.
