package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/appscode/go/ioutil"
	"github.com/appscode/go/log"
	"github.com/appscode/go/runtime"
	"github.com/ghodss/yaml"
	"github.com/kardianos/osext"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/kubectl/plugins"
)

func NewCmdInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "install",
		Short:             "Install as kubectl plugin",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			var e []string
			for _, pair := range os.Environ() {
				if strings.HasPrefix(pair, "KUBECTL_") {
					e = append(e, pair)
				}
			}
			sort.Strings(e)
			for _, v := range e {
				fmt.Println(v)
			}

			dir := runtime.GOPath() + "/src/github.com/tamalsaha/kubectl-plugin-demo/demo"
			os.MkdirAll(dir, 0755)

			p, err := osext.Executable()
			if err != nil {
				log.Fatal(err)
			}
			p = filepath.Clean(p)
			ioutil.CopyFile(filepath.Join(dir, filepath.Base(p)), p, 0755)

			var traverse func(cmd *cobra.Command, p *plugins.Plugin)
			traverse = func(cmd *cobra.Command, p *plugins.Plugin) {
				p.Name = cmd.Name()
				p.ShortDesc = cmd.Short
				p.LongDesc = cmd.Long
				p.Example = cmd.Example
				p.Command = "./" + strings.TrimSpace(cmd.CommandPath())

				cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
					if flag.Hidden {
						return
					}
					p.Flags = append(p.Flags, plugins.Flag{
						Name:      flag.Name,
						Shorthand: flag.Shorthand,
						Desc:      flag.Usage,
						DefValue:  flag.DefValue,
					})
				})

				for _, cc := range cmd.Commands() {
					cp := &plugins.Plugin{}
					traverse(cc, cp)
					p.Tree = append(p.Tree, cp)
				}
			}

			plugin := &plugins.Plugin{}
			traverse(NewRootCmd(), plugin)

			data, err := yaml.Marshal(plugin)
			if err != nil {
				log.Fatal(err)
			}
			ioutil.WriteFile(filepath.Join(dir, "plugins.yaml"), bytes.NewBuffer(data), 0755)
		},
	}
	return cmd
}
