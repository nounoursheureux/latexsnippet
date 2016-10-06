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
	// TODO: use ioutils.TempDir
	/* err = os.MkdirAll("/tmp/latex-snippet", 0777)
	handle(err)
	err = os.Chdir("/tmp/latex-snippet")
	handle(err) */

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

	/* png_tmp, err := os.Open("snippet.png")
	err = png_tmp.Sync()
	handle(err)
	var buf []byte
	_, err = png_tmp.Read(buf)
	handle(err)
	err = png_tmp.Close()
	handle(err)
	println(len(buf)) */

	err = os.Chdir(cwd)
	handle(err)

	err = ioutil.WriteFile("snippet.png", buf, 0666)
	handle(err)

	/* png_dest, err := os.Create("snippet.png")
	handle(err)
	_, err = png_dest.Write(buf)
	handle(err)
	err = png_dest.Close()
	handle(err) */
}
