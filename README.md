# CodePlayground

CodePlayground is a playground tool for go and rust language.

[![Go Report Card](https://goreportcard.com/badge/github.com/Trendyol/code-playground)](https://goreportcard.com/report/github.com/Trendyol/code-playground)
[![Go Doc Card](https://godoc.org/github.com/Trendyol/code-playground?status.svg)](https://godoc.org/github.com/Trendyol/code-playground)

## Installation

Use homebrews to install code-playground.

```bash
brew tap trendyol/trendyol-tap
brew install code-playground
```

## Usage

Commands open default (vim) editor

```bash
code-playground go

code-playground rust
```

If you want to use another editor for code-playground, export ``PLAYGROUND ENVIRONMENT`` value. 

e.g. ``export PLAYGROUND_ENVIRONMENT="subl -w"`` (use sublime editor)

e.g. ``export PLAYGROUND_ENVIRONMENT="code -w"`` (use vscode editor)

## Flags

### Share (-s)
This flag generate to playground link.

```bash
code-playground go -s

code-playground rust -s
```

### Import (-i)
This flag import to shared playground link.

```bash
code-playground go -i "https://play.golang.org/p/9geTEmeOzJO"

code-playground rust -i "https://play.rust-lang.org/?gist=62a1c0b6f2aee0f7ebd78cfbddaae0e4"
```

Output

```bash

>>> code-playground rust -s

-------------------------------------------------
Hello, code-playground
-------------------------------------------------
⇨ https://play.rust-lang.org/?gist=7f28f8a1a7f35c35a903f983fa95b9e0
```
