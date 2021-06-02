/*
Copyright © 2019 Scott Crawford scottcrawford03@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/felix-cli/felix/internal/builder"
)

var (
	currentVersion = "1.0.0-beta"

	app = kingpin.New("felix", "Golang template tool")

	versionCommand = app.Command("version", "list the latest version")
	initCommand    = app.Command("init", "creates new service in current directory")
	newCommand     = app.Command("new", "creates new service in new directory")
	name           = newCommand.Arg("name", "the name of the new service and directory you want to create").Required().String()

	templateURL = app.Flag("template", "github url for the template").Short('t').String()
)

func main() {
	versionCommand.Action(getVersion)
	initCommand.Action(felixInit)
	newCommand.Action(felixNew)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func getVersion(c *kingpin.ParseContext) error {
	fmt.Println(currentVersion)
	return nil
}

func felixInit(c *kingpin.ParseContext) error {
	tmp := builder.Template{
		URL: *templateURL,
	}

	if err := builder.Init(&tmp); err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())

		return err
	}
	fmt.Println("All done!")
	return nil
}

func felixNew(c *kingpin.ParseContext) error {
	tmp := builder.Template{
		Name: *name,
		URL:  *templateURL,
	}

	if err := builder.Init(&tmp); err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())

		return err
	}
	fmt.Println("All done!")
	return nil
}
