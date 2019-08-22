# homebrew-go

[![GoDoc][godoc-badge]][godoc]

A Go client package for Homebrew.

## Usage

```
import "github.com/micnncim/homebrew-go"
```

```go
func main() {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	h := homebrew.NewHomebrew(
		homebrew.WithStdout(stdout),
		homebrew.WithStderr(stderr),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	if err := h.Install(ctx, "jq", homebrew.InstallOptionVerbose, homebrew.InstallOptionForce); err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
}
```

## Supported commands

Currently only *essential commands*.

- `brew install [options] formula`
- `brew uninstall, rm, remove [options] formula`
- `brew list, ls [options] [formula]`
- `brew search [options] [text|/text/]`

## Dependency

Need to install [Homebrew](https://docs.brew.sh/) to use this package.

## License

MIT

<!-- badge links -->

[godoc]: https://godoc.org/github.com/micnncim/homebrew-go 

[godoc-badge]: https://img.shields.io/badge/godoc.org-reference-blue.svg 
