package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const DefaultEditor = "vim"

type Editor struct {
	File    string
	Content string
}

func NewEditor(file, content string) *Editor {
	return &Editor{
		File:    file,
		Content: content,
	}
}

func (e *Editor) Capture() error {
	editor := os.Getenv("PLAY_EDITOR")

	if editor == "" {
		editor = DefaultEditor
	}

	executables := strings.Split(editor, " ")

	executables = append(executables, e.File)

	editor = executables[0]

	executableEditor, err := exec.LookPath(editor)

	if err != nil {
		return err
	}

	cmd := exec.Command(executableEditor, executables[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (e *Editor) Open() (string, error) {

	var file *os.File
	var err error

	if len(e.File) > 0 {

		e.File = e.File + ".playground"

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

		file, err = ioutil.TempFile(os.TempDir(), "play-*.playground")

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
