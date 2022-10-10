package main

import (
	"fmt"
	"os"
	"strings"

	createproj "github.com/mizumoto-cn/gcp-go-tut/create_proj"
	"github.com/mizumoto-cn/gcp-go-tut/create_proj/project"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Create GCloud Project",
		Usage: "Semi-automatic creating project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Usage:   "Project name",
				Aliases: []string{"n"},
			},
			&cli.StringFlag{
				Name:    "project-id",
				Usage:   "Project ID",
				Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:    "labels",
				Usage:   "Project labels, pls write in format of key1=value1,key2=value2,...",
				Aliases: []string{"ls"},
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			id := c.String("project-id")
			l := c.String("labels")
			var labels map[string]string
			if l != "" {
				labels = make(map[string]string)
				for _, v := range strings.Split(l, ",") {
					// split by "="
					kv := strings.Split(v, "=")
					if len(kv) == 2 {
						labels[kv[0]] = kv[1]
					}
				}
			}
			if name == "" || id == "" || len(labels) == 0 {
				fmt.Println("All flags are blank, will be named by default: Default-Proj-Name || Default-Proj-Id")
			}
			b := project.NewProjectBuilder()
			p, err := b.SetName(name).SetProjectId(id).SetLabels(labels).Build()
			if err != nil {
				return err
			}
			err = createproj.CreateProj(p)
			if err != nil {
				return err
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
