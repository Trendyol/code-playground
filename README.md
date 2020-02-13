# CodePlayground

CodePlayground is a playground tool for go and rust language.

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

If you want use another editor for code-playground, export to ``PLAYGROUND ENVIRONMENT`` value. 

e.g. ``export PLAYGROUND_ENVIRONMENT="subl -w"`` (use sublime editor)

e.g. ``export PLAYGROUND_ENVIRONMENT="code -w"`` (use vscode editor)

## Flags

### Share (-s)
This flag generate to playground link.

```bash
code-playground go -s

code-playground rust -s
```

Output

```bash

>>> code-playground rust -s

-------------------------------------------------
Hello, code-playground
-------------------------------------------------
⇨ https://play.rust-lang.org/?gist=7f28f8a1a7f35c35a903f983fa95b9e0
```
