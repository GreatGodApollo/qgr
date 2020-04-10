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
	"github.com/spf13/cobra"
	"os"
	"text/template"
	"time"
)

var project projInfo
var doInit bool

type projInfo struct {
	AuthorName         string
	AuthorUsername     string
	AuthorEmail        string
	ProjectName        string
	ProjectDescription string
	Year               string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qgr",
	Short: "A quick and easy way to start your git repos",
	Long: `This command line utility gives you a simple place
to start with your git repository. It starts you out
with a fancy shmancy README.md and the MIT LICENSE.`,
	Run:     run,
	Version: "1.1.0",
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

	if doInit {
		_, err = git.PlainInit("./", false)
		if err != nil {
			fmt.Println("Could not initialize git repo: ", err)
		}
	}

	fmt.Println("Successfully generated repository!")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&project.AuthorName, "author", "a", "", "author's name")
	rootCmd.Flags().StringVarP(&project.AuthorUsername, "authorUsername", "u", "", "author's username")
	rootCmd.Flags().StringVarP(&project.ProjectName, "name", "n", "", "project's name")
	rootCmd.Flags().StringVarP(&project.ProjectDescription, "description", "d", "", "project's description")

	rootCmd.Flags().BoolVarP(&doInit, "init", "i", false, "initialize git repo")

	_ = rootCmd.MarkFlagRequired("author")
	_ = rootCmd.MarkFlagRequired("authorUsername")
	_ = rootCmd.MarkFlagRequired("name")
	_ = rootCmd.MarkFlagRequired("description")

	project.Year = time.Now().Format("2006")
}
