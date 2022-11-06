package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var spell = goyek.Define(goyek.Task{
	Name:  "spell",
	Usage: "misspell",
	Action: func(a *goyek.A) {
		if !cmd.Exec(a, "go install github.com/client9/misspell/cmd/misspell") {
			return
		}
		mdFiles := find(a, ".md")
		if len(mdFiles) == 0 {
			a.Skip("no .md files")
		}
		cmd.Exec(a, "misspell -error -locale=US -w "+strings.Join(mdFiles, " "))
	},
})

func find(a *goyek.A, ext string) []string {
	a.Helper()
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
		a.Fatal(err)
	}
	return files
}
