/*
Copyright © 2020 Brett Bender

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
	"os"
	"text/template"
	"time"
)

var project projInfo

type projInfo struct {
	AuthorName			string
	AuthorUsername		string
	AuthorEmail			string
	ProjectName			string
	ProjectDescription	string
	Year				string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qgr",
	Short: "A quick and easy way to generate a git repository",
	Long: `This command line utility gives you a simple
place to start with your git repository. It starts you
out with a fancy shmancy README.md and the MIT LICENSE.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: run,
}

func run(cmd *cobra.Command, args []string) {

	lTmpl, err := template.New("license").Parse(licenseTemplate())

	if err != nil {
		fmt.Printf("Error parsing license generation template\n%v\n", err)
		os.Exit(1)
	}

	rTmpl, err := template.New("readme").Parse(readmeTemplate())
	if err != nil {
		fmt.Printf("Error parsing license generation template\n%v\n", err)
		os.Exit(1)
	}

	f, err := os.Create("./LICENSE")
	if err != nil {
		fmt.Println("Could not create LICENSE: ", err)
	} else {
		err = lTmpl.Execute(f, project)
		if err != nil {
			fmt.Println("Could not create LICENSE: ", err)
		}
	}

	f, err = os.Create("./README.md")
	if err != nil {
		fmt.Println("Could not create README.md: ", err)
	} else {
		err = rTmpl.Execute(f, project)
		if err != nil {
			fmt.Println("Could not create README.md: ", err)
		}
	}


	fmt.Println("Successfully generated repository!")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&project.AuthorName, "author", "a", "", "Author's name")
	rootCmd.Flags().StringVarP(&project.AuthorUsername, "authorUsername", "u", "", "Author's Username")
	rootCmd.Flags().StringVarP(&project.ProjectName, "name", "n", "", "Project's name")
	rootCmd.Flags().StringVarP(&project.ProjectDescription, "description", "d", "", "Project's description")

	_ = rootCmd.MarkFlagRequired("author")
	_ = rootCmd.MarkFlagRequired("authorUsername")
	_ = rootCmd.MarkFlagRequired("name")
	_ = rootCmd.MarkFlagRequired("description")

	project.Year = string(time.Now().Year())
}

func licenseTemplate() string {
	var tmpl =
		`MIT License

Copyright (c) {{.Year}} {{.AuthorName}}

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`

	return tmpl
}

func readmeTemplate() string {
	var tmpl =
		`<h1 align="center">{{.ProjectName}}</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/{{.AuthorUsername}}">@{{.AuthorUsername}}</a></i></p>

{{.ProjectDescription}}

## Built With
> Used technologies here

## Usage
> How to use this software

## Licensing

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)

## Authors

* [{{.AuthorName}}](https://github.com/{{.AuthorUsername}})`

	return tmpl
}