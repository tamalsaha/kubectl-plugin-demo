package main

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

func NewCmdList(clientConfig clientcmd.ClientConfig) *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "list",
		Short:             `list`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdListNodes(clientConfig))
	return cmd
}
