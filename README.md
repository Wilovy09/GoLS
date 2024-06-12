# GoLS

We deserve a better and faster LS command.

> [!NOTE]
> PR is accepted to improve the code

> [!WARNING]
> This project is under construction, it has visual errors more than anything else.

## Preview

![GoLS](./assets/previews/preview-GoLS-v1.gif)

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

## Download de binary

In the [release](https://github.com/Wilovy09/GoLS/releases/tag/release) section you can find the binary file used to run GoLS.

Thins binary will be moved to yours root directory.

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

The `gols` binary will be moved to the root directory `/home/USER`.

You need to create a `alias` in your `.bashrc` or `.zshrc` file.

```bash
# [ gols better ls ]
alias ls='~/GoLS'
```

## Flags

| Flag     | Description                 | Values  | Default |
|----------|-----------------------------|---------|-------- |
| -a       | Show hidden files           | boolean | false   |
| -depth   | How deep the search will be | int     | 0       |
| -details | Show de file/dir details    | bollean | false   |
| -tree    | Show the files in tree view | boolean | false   |
| -emoji   | Show emojis in the view     | boolean | false   |
