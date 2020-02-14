package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const GoHelloWorld = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, code-playground")
}
`

type GoCompile struct {
	Errors      string           `json:"Errors"`
	Events      []GoCompileEvent `json:"Events"`
	Status      int              `json:"Status"`
	IsTest      bool             `json:"IsTest"`
	TestsFailed int              `json:"TestsFailed"`
	VetOK       bool             `json:"VetOK"`
}
type GoCompileEvent struct {
	Message string `json:"Message"`
	Kind    string `json:"Kind"`
	Delay   int32  `json:"Delay"`
}

type Go struct {
	Code string
}

func (r *Go) Type() string {
	return ".go"
}

func (g *Go) Init(code string) {
	g.Code = code
}

func (g *Go) Default() string {
	return GoHelloWorld
}

func (g *Go) Evaluate() {

	Spinner.Start()
	defer Spinner.Stop()

	values := url.Values{
		"version": {"2"},
		"body":    {g.Code},
		"withVet": {"true"},
	}

	values.Add("version", "2")
	values.Add("body", g.Code)
	values.Add("withVet", "true")

	response, err := http.PostForm("https://play.golang.org/compile", values)

	if err != nil {
		Log.Error(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Log.Error(err)
		return
	}

	result := &GoCompile{}

	err = json.Unmarshal(body, result)

	if err != nil {
		Log.Error(err)
		return
	}

	result.Execute()

}

func (g *Go) Share() {

	Spinner.Start()
	defer Spinner.Stop()

	response, err := http.Post("https://play.golang.org/share", "text/plain", strings.NewReader(g.Code))

	if err != nil {
		Log.Error(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Log.Error(err)
		return
	}

	var result = string(body)

	Log.Info("https://play.golang.org/p/" + result)

}

func (g *Go) Import(path string) string {

	Spinner.Start()
	defer Spinner.Stop()

	if !strings.HasPrefix(path, "https://play.golang.org/p/") {
		Log.Error("Import path is not valid")
		return ""
	}

	response, err := http.Get(path)

	if err != nil {
		Log.Error(err)
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Log.Error(err)
		return ""
	}

	result := string(body)

	tokens := strings.Split(result, "id=\"code\"")

	tokens = strings.Split(tokens[1], ">")
	tokens = strings.Split(tokens[1], "<")

	result = html.UnescapeString(tokens[0])

	Spinner.Stop()

	return result
}

func (g *GoCompile) Execute() {

	Spinner.Stop()

	fmt.Println("-------------------------------------------------")

	if g.VetOK {
		for _, event := range g.Events {
			time.Sleep(time.Duration(event.Delay))
			fmt.Print(event.Message)
		}
	} else {
		Log.Error(g.Errors)
	}

	fmt.Println("-------------------------------------------------")

}
