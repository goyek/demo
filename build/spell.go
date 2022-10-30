package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/workflow"
)

var spell = goyek.Define(goyek.Task{
	Name:  "spell",
	Usage: "misspell",
	Action: func(tf *goyek.TF) {
		if !workflow.Exec(tf, "go install github.com/client9/misspell/cmd/misspell") {
			return
		}
		mdFiles := find(tf, ".md")
		if len(mdFiles) == 0 {
			tf.Skip("no .md files")
		}
		workflow.Exec(tf, "misspell -error -locale=US -w "+strings.Join(mdFiles, " "))
	},
})

func find(tf *goyek.TF, ext string) []string {
	tf.Helper()
	var files []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ext {
			files = append(files, filepath.ToSlash(path))
		}
		return nil
	})
	if err != nil {
		tf.Fatal(err)
	}
	return files
}
