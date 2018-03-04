package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
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

		},
	}
	return cmd
}
