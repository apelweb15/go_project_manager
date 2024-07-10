package main

import (
	"bufio"
	"fmt"
	"gitlab.com/apelweb/go_project_manager/internal"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//Meminta input nama
	fmt.Print("Module Name: ")
	moduleName, _ := reader.ReadString('\n')
	moduleName = strings.TrimSpace(moduleName)

	//// Meminta input alamat
	fmt.Print("Instalation directory: ")
	directory, _ := reader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	project := internal.Project{
		ModuleName: moduleName,
		Directory:  directory,
	}
	err := project.Start()
	if err != nil {
		panic(err)
	}
}
