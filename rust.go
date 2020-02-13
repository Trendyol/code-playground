package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const RustHelloWorld = `
fn main() {
    println!("Hello, code-playground");
}
`

type RustCompileRequest struct {
	Version  string `json:"version"`
	Optimize string `json:"optimize"`
	Code     string `json:"code"`
	Edition  string `json:"edition"`
}

type RustCompile struct {
	Result string  `json:"result"`
	Error  *string `json:"error"`
}

type RustShareRequest struct {
	Code string `json:"code"`
}

type RustShare struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Code string `json:"code"`
}

type Rust struct {
	Code string
}

func (r *Rust) Init(code string) {
	r.Code = code
}

func (r *Rust) Default() string {
	return RustHelloWorld
}

func (r *Rust) Evaluate() {

	Spinner.Start()
	defer Spinner.Stop()

	request := RustCompileRequest{
		Version:  "stable",
		Optimize: "0",
		Code:     r.Code,
		Edition:  "2018",
	}

	req, err := json.Marshal(request)

	if err != nil {
		Log.Error(err)
		return
	}

	response, err := http.Post("https://play.rust-lang.org/evaluate.json", "application/json", bytes.NewBuffer(req))

	if err != nil {
		Log.Error(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Log.Error(err)
		return
	}

	result := &RustCompile{}

	err = json.Unmarshal(body, result)

	if err != nil {
		Log.Error(err)
		return
	}

	result.Execute()

}

func (r *Rust) Share() {

	Spinner.Start()
	defer Spinner.Stop()

	request := RustShareRequest{
		Code: r.Code,
	}

	req, err := json.Marshal(request)

	if err != nil {
		Log.Error(err)
		return
	}

	response, err := http.Post("https://play.rust-lang.org/meta/gist/", "application/json", bytes.NewBuffer(req))

	if err != nil {
		Log.Error(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Log.Error(err)
		return
	}

	result := &RustShare{}

	err = json.Unmarshal(body, result)

	if err != nil {
		Log.Error(err)
		return
	}

	Log.Info("https://play.rust-lang.org/?gist=" + result.ID)

}

func (r *Rust) Import(path string) string {

	Spinner.Start()
	defer Spinner.Stop()

	if !strings.HasPrefix(path, "https://play.rust-lang.org") {
		Log.Error("Import path is not valid")
		return ""
	}

	url, err := url.Parse(path)

	if err != nil {
		Log.Error(err)
		return ""
	}

	gist := url.Query().Get("gist")

	if len(gist) == 0 {
		Log.Error("gist not found")
		return ""
	}

	response, err := http.Get("https://play.rust-lang.org/meta/gist/" + gist)

	if err != nil {
		Log.Error(err)
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Log.Error(err)
		return ""
	}

	result := &RustShare{}

	err = json.Unmarshal(body, result)

	if err != nil {
		Log.Error(err)
		return ""
	}

	r.Code = result.Code

	Spinner.Stop()

	return result.Code

}

func (r *RustCompile) Execute() {

	Spinner.Stop()

	fmt.Println("-------------------------------------------------")

	if r.Error == nil {
		fmt.Print(r.Result)
	} else {
		Log.Error(*r.Error)
	}

	fmt.Println("-------------------------------------------------")

}
