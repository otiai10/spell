# spell

[![Go](https://github.com/otiai10/spell/actions/workflows/go.yml/badge.svg)](https://github.com/otiai10/spell/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/otiai10/spell/branch/master/graph/badge.svg)](https://codecov.io/gh/otiai10/spell)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/spell)](https://goreportcard.com/report/github.com/otiai10/spell)

Command line string parser for Go.

Given

```
bash -c 'for i in 0 1 2 3; do echo $i && sleep 1; done' && echo "foo baa"
```

Want

```sh
# []string{}
bash
-c
'for i in 0 1 2 3; do echo $i && sleep 1; done'
&&
echo
"foo baa"
```

# Example

```go
s := `bash -c 'for i in 0 1 2 3; do echo $i && sleep 1; done' && echo "foo baa"`
tokens := spell.Parse(s)
```
