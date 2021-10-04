package main

import (
	"github.com/google/shlex"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const DefaultEditor = "vim"

type Editor struct {
	File    string
	Type    string
	Content string
}

func NewEditor(file, t, content string) *Editor {
	return &Editor{
		File:    file,
		Type:    t,
		Content: content,
	}
}

func (e *Editor) Capture() error {
	editor := os.Getenv("PLAY_EDITOR")
	if editor == "" {
		editor = DefaultEditor

		if runtime.GOOS == "windows" {
			editor = "notepad"
		}
	}

	executables, err := shlex.Split(editor)
	if err != nil {
		executables = []string{strings.Split(editor, " ")[0]}
	}
	executables = append(executables, e.File)

	cmd := exec.Command(executables[0], executables[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (e *Editor) Open() (string, error) {

	var file *os.File
	var err error

	if len(e.File) > 0 {

		e.File = e.File + ".playground" + e.Type

		file, err = os.Open(e.File)

		if os.IsNotExist(err) {

			file, err = os.Create(e.File)

			if err != nil {
				return "", err
			}
		}

		f, err := file.Stat()

		if err != nil {
			return "", err
		}

		if f.Size() == 0 {

			_, err = file.WriteString(e.Content)

			if err != nil {
				return "", err
			}
		}

	} else {

		file, err = ioutil.TempFile(os.TempDir(), "play-*.playground"+e.Type)

		if err != nil {
			return "", err
		}

		_, err = file.Write([]byte(e.Content))

		if err != nil {
			return "", err
		}

		defer os.Remove(e.File)

	}

	e.File = file.Name()

	if err = file.Close(); err != nil {
		return "", err
	}

	if err = e.Capture(); err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(e.File)

	if err != nil {
		return "", err
	}

	result := string(bytes)

	return result, nil
}
