package main

type IPlayground interface {
	Init(string)
	Evaluate()
	Share()
	Default() string
}
