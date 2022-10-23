package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
}

// Scenario 1 - turn a doc path into an indexed tree of documents to start the generation from

func CompileSourceFolder(path string) (DocumentationSourceFolder, error) {

	return DocumentationSourceFolder{}, errors.New("to implement")
}

type DocumentationSourceFolder struct {
	Path string
}

func (d DocumentationSourceFolder) GetPeopleFilePath() string {

	return "todo"
}

type DocumentationFolder struct {
	Path         string
	DocFiles     []DocumentationFile
	OpenApiFiles []OpenApiFile
}

func (d DocumentationFolder) GetFolderFilePath() string {

	return "todo"
}

type DocumentationFile struct {
	Path string
}

type OpenApiFile struct {
	Path string
}
