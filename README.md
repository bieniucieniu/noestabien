###run:

```bash
templ generate && go run \*.go
```

Start at [http://localhost:3000](http://localhost:3000).

###requirements

ensure installed templ cli

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

and `$HOME/go/bin` added to PATH

```bash
export PATH="$PATH:$HOME/go/bin"
```

in `.zshrc`, `.bashrc` or similar.

[go templ](https://templ.guide/)
