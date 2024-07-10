package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/apelweb15/go_project_manager/internal"
	"os"
	"strings"
)

//go:embed template/*
var template embed.FS

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Module Name: ")
	moduleName, _ := reader.ReadString('\n')
	moduleName = strings.TrimSpace(moduleName)
	fmt.Print("Instalation directory: ")
	directory, _ := reader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	project := internal.Project{
		ModuleName: moduleName,
		Directory:  directory,
		Template:   template,
	}
	err := project.Start()
	if err != nil {
		panic(err)
	}
}
