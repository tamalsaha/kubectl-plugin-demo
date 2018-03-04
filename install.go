package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/appscode/go/runtime"
	"github.com/kardianos/osext"
	"github.com/spf13/cobra"
	"path/filepath"
	"github.com/appscode/go/log"
	"github.com/appscode/go/ioutil"
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
		},
	}
	return cmd
}
