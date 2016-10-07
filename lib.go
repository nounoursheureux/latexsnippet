package latexsnippet

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

func RenderSnippet(snippet string, outpath string) error {
	var cwd, err = os.Getwd()
	if err != nil {
		return err
	}

	tmpdir, err := ioutil.TempDir("", "latexsnippet")
	if err != nil {
		return err
	}
	os.Chdir(tmpdir)

	defer os.RemoveAll(tmpdir)

	var latex_string = template_begin + "\n" + snippet + "\n" + template_end
	latex_file, err := os.Create("snippet.tex")
	if err != nil {
		return err
	}
	defer latex_file.Close()

	latex_file.WriteString(latex_string)

	var latex_cmd = exec.Command("latex", "-shell-escape", latex_file.Name())
	latex_cmd.Stdout = os.Stdout
	latex_cmd.Stderr = os.Stderr
	err = latex_cmd.Run()
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadFile("snippet.png")
	if err != nil {
		return err
	}

	err = os.Chdir(cwd)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outpath, buf, 0666)
	if err != nil {
		return err
	}

	return nil
}
