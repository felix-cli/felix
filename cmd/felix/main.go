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
	fixitCommand   = app.Command("fixit", "list the latest version")
	org            = fixitCommand.Flag("org", "Set your org name").Short('o').String()
	proj           = fixitCommand.Flag("project", "Set your project name").Short('p').String()
)

func main() {
	versionCommand.Action(getVersion)
	fixitCommand.Action(fixit)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func getVersion(c *kingpin.ParseContext) error {
	fmt.Println(currentVersion)
	return nil
}

func fixit(c *kingpin.ParseContext) error {
	tmp := builder.Template{
		Org:  *org,
		Proj: *proj,
	}

	if err := builder.Init(&tmp); err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())

		return err
	}
	fmt.Println("All done!")
	return nil
}
