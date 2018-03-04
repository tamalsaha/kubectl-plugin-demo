package main

import (
	"fmt"

	"github.com/appscode/go/log"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewCmdListNodes(clientConfig clientcmd.ClientConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list-nodes",
		Short:             "List nodes",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			namespace, _, err := clientConfig.Namespace()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("namespace = ", namespace)

			config, err := clientConfig.ClientConfig()
			if err != nil {
				log.Fatalln(err)
			}
			client := kubernetes.NewForConfigOrDie(config)
			nodes, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			for _, node := range nodes.Items {
				fmt.Println(node.Name)
			}
		},
	}
	return cmd
}
