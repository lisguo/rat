/*
Generates LaTeX resume components given a .cls template file
*/

package lib

import (
	"fmt"
)

func WriteDocumentClass(documentClass string){
	fmt.Println("Writing documentclass!")
	strTemplate := "\\documentclass{%s}"
	fmt.Sprintf(strTemplate, documentClass)
}