package main

import (
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"time"
)

var Spinner = spinner.New(spinner.CharSets[43], 100*time.Millisecond)
var Log = NewLogger()

func main() {

	var file string
	var path string
	var share bool

	var evaluate = func(playground IPlayground) {

		init := playground.Default()

		if len(path) > 0 {
		 	init = playground.Import(path)
		}

		editor := NewEditor(file, init)

		code, err := editor.Open()

		if err != nil {
			Log.Error(err)
			return
		}

		playground.Init(code)

		playground.Evaluate()

		if share {
			playground.Share()
		}
	}

	var cmd = &cobra.Command{
		Use:  "play",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	var goCmd = &cobra.Command{
		Use:  "go",
		Args: cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 0 {
				file = args[0]
			}

			evaluate(new(Go))
		},
	}

	var rustCmd = &cobra.Command{
		Use:  "rust",
		Args: cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 0 {
				file = args[0]
			}

			evaluate(new(Rust))
		},
	}

	goCmd.Flags().BoolVarP(&share, "share", "s", false, "share playground")
	rustCmd.Flags().BoolVarP(&share, "share", "s", false, "share playground")
	rustCmd.Flags().StringVarP(&path, "import", "i", "", "import playground")

	cmd.AddCommand(goCmd, rustCmd)

	_ = cmd.Execute()
}
