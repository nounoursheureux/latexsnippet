package main

import (
	"log"
	"os"
	"os/exec"
	"io/ioutil"
)

var template_begin = `\documentclass[convert={density=800}]{standalone}
\begin{document}`

var template_end = `\end{document}`

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: " + os.Args[0] + " <LaTeX snippet>")
	}

	var cwd, err = os.Getwd()

	tmpdir, err := ioutil.TempDir("", "latex-snippet")
	handle(err)
	os.Chdir(tmpdir)

	defer os.RemoveAll(tmpdir)

	var latex_string = template_begin + "\n" + os.Args[1] + "\n" + template_end
	latex_file, err := os.Create("snippet.tex")
	handle(err)

	latex_file.WriteString(latex_string)

	var latex_cmd = exec.Command("latex", "-shell-escape", latex_file.Name())
	latex_cmd.Stdout = os.Stdout
	latex_cmd.Stderr = os.Stderr
	err = latex_cmd.Run()
	handle(err)
	err = latex_file.Close()
	handle(err)

	buf, err := ioutil.ReadFile("snippet.png")
	handle(err)

	err = os.Chdir(cwd)
	handle(err)

	err = ioutil.WriteFile("snippet.png", buf, 0666)
	handle(err)
}
