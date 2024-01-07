[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/words)
[![codecov](https://codecov.io/gh/go-corelibs/words/graph/badge.svg?token=99HKVL3Y3f)](https://codecov.io/gh/go-corelibs/words)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/words)](https://goreportcard.com/report/github.com/go-corelibs/words)

# words - word metrics from mixed-locale content

words is a package for counting the numbers of words and providing simple
metrics like estimated reading times.

# Installation

``` shell
> go get github.com/go-corelibs/words@latest
```

# Examples

## Count

``` go
func main() {
    text := "さらに「やり遂げる」ためのEnjin"
    count := words.Count(text)
    // count == 12
    fmt.Printf("There are %d words in %q\n", count, text)
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
