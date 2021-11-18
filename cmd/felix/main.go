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

	"github.com/alecthomas/kong"
	"github.com/felix-cli/felix/internal/builder"
	endErrors "github.com/felix-cli/felix/internal/end_errors"
)

type Context struct {
	TemplateURL string
}

type VersionCmd struct{}

type InitCmd struct{}

type NewCmd struct {
	Name string `arg:"" required:"" help:"the name of the new service and directory you want to create"`
}

var cli struct {
	Template string `short:"t" help:"github url for the template"`

	Version VersionCmd `cmd:"" help:"list the latest version"`

	Init InitCmd `cmd:"" help:"creates new service in current directory"`

	New NewCmd `cmd:"" help:"creates new service in new directory"`
}

var (
	Version = "DEV"
)

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{TemplateURL: cli.Template})
	ctx.FatalIfErrorf(err)
}

func (v *VersionCmd) Run() error {
	fmt.Println(Version)
	return nil
}

func (i *InitCmd) Run(ctx *Context) error {
	endErrorHandler := endErrors.GetInstance()
	tmp := builder.Template{
		URL: ctx.TemplateURL,
	}

	if err := builder.Init(&tmp); err != nil {
		endErrorHandler.AddErrorf("running felix init: %s", err.Error())
		endErrorHandler.PrintErrors()
		return err
	}
	fmt.Println("All done!")
	return nil
}

func (n *NewCmd) Run(ctx *Context) error {
	endErrorHandler := endErrors.GetInstance()
	tmp := builder.Template{
		Name: n.Name,
		URL:  ctx.TemplateURL,
	}

	if err := builder.Init(&tmp); err != nil {
		endErrorHandler.AddErrorf("running felix new: %s", err.Error())
		endErrorHandler.PrintErrors()

		return nil
	}
	fmt.Println("All done!")
	return nil
}
