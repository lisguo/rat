/*
Generates LaTeX resume components given a .cls template file
*/

package lib

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteHeader(f *os.File) {
	text := fmt.Sprintf(`
	\documentclass{resume} % Use the custom resume.cls style

	\usepackage[left=0.75in,top=0.6in,right=0.75in,bottom=0.6in]{geometry} % Document margins
	\newcommand{\tab}[1]{\hspace{.2667\textwidth}\rlap{#1}}
	\newcommand{\itab}[1]{\hspace{0em}\rlap{#1}}
	`)
	fmt.Fprintln(f, text)
}

func WriteBeginDocument(f *os.File) {
	fmt.Fprintln(f, "\\begin{document}")
}

func WriteEndDocument(f *os.File) {
	fmt.Fprintln(f, "\\end{document}")
}

func WriteBeginSection(f *os.File, section string) {
	text := fmt.Sprintf(`
	\begin{rSection}{%s}
	`, section)
	fmt.Fprintln(f, text)
}

func WriteEndSection(f *os.File) {
	fmt.Fprintln(f, "\\end{rSection}")
}

func WriteContactInfo(f *os.File, name string, address string,
	phone string, email string) {
	text := fmt.Sprintf(`
	\name{%s}
	\address{%s}
	\address{%s \\ %s}
	`, name, address, phone, email)
	fmt.Fprintln(f, text)
}

func WriteEducation(f *os.File, institution string, startDate string,
	endDate string, degree string, department string, gpa string) {
	text := fmt.Sprintf(`
	{\bf %s} \hfill {\em %s - %s}
	\\ %s
	\\ %s\\
	`, institution, startDate, endDate, degree, department)
	fmt.Fprintln(f, text)

	if gpa != "" {
		gpaText := fmt.Sprintf(`
		\hfill {GPA: %s}
		`, gpa)
		fmt.Fprintln(f, gpaText)
	}
}
