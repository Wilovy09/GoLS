# GoLS a common LS

We are a project that seeks to bring together different functionalities of different existing LS.

- We are not the fastest
- We are not the best
- We're just different.

We try to make all without external libreries.

> [!NOTE]
> PRs are welcome to improve the code.

> [!WARNING]
> This project is under construction.

## Preview

![GoLS](./assets/previews/preview-GoLS-v1.gif)

## Installation

> [!NOTE]
> You need to have [Go](https://go.dev/dl/) installed on your machine.

It's to simple to install, just open your terminal and run the following command.

```bash
go install github.com/Wilovy09/GoLS@latest
```

Find the `gols` binary in your `$GOPATH/bin` directory and create an alias in your `.bashrc` or `.zshrc` file.

```bash
# [ GoLS better ls ]
alias ls='~/go/bin/GoLS --tree --emoji'
```

And you're done.

## Download de binary

In [latest release](https://github.com/Wilovy09/GoLS/releases/latest) you can find the latest binary file used to run GoLS.

This binary must be moved to a directory in `%PATH`.

Create an alias in your `bashrc` or `.zshrc` file.

```bash
# [ GoLS better ls ]
alias ls='~/GoLS'
```

And you're done.

### Compile

Clone the repository and compile the code.

```bash
go build -o GoLS -buildvcs=false
```

The `gols` binary must be moved to the directory `$HOME`.

Create an `alias` in your `.bashrc` or `.zshrc` file.

```bash
# [ gols better ls ]
alias ls='~/GoLS'
```

## Flags

| Flag     | Description                              | Values  | Default |
| -------- | ---------------------------------------- | ------- | ------- |
| -a       | Show hidden files                        | boolean | false   |
| -depth   | How deep the search will be              | int     | 0       |
| -details | Show de file/dir details                 | bollean | false   |
| -tree    | Show the files in tree view              | boolean | false   |
| -emoji   | Show emojis in the view                  | boolean | false   |
| -d       | Show only directorys                     | boolean | false   |
| -f       | Show only files                          | boolean | false   |
| -n       | Show how many files each folder contains | boolean | false   |
| -h       | Show the help message                    | boolean | false   |
 
## Features

- [ ] Add a different flags to details (permissions, creator, date created)
- [x] Add how many files each folder contains `[n]Desktop/` `[n]📁Desktop` can enabled with flag
- [ ] Add colors
