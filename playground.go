package main

type IPlayground interface {
	Init(string)
	Import(string) string
	Evaluate()
	Share()
	Default() string
	Type() string
}
