package main

import (
	"flag"
	"fmt"
	stdlog "log"
	"os"
	"sort"
	"strings"

	"github.com/appscode/go/log"
	logs "github.com/appscode/go/log/golog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewRootCmd() *cobra.Command {
	var (
		enableAnalytics = true
	)
	var rootCmd = &cobra.Command{
		Short:             `Log Demo`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.Flags().VisitAll(func(flag *pflag.Flag) {
				stdlog.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			})
		},
	}
	// rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	rootCmd.PersistentFlags().BoolVar(&enableAnalytics, "analytics", enableAnalytics, "Send analytical events to Google Analytics")

	rootCmd.AddCommand(NewCmdCheck())
	rootCmd.AddCommand(NewCmdInstall())
	return rootCmd
}

func NewCmdCheck() *cobra.Command {
	var (
		masterURL      string
		kubeconfigFile string
	)
	cmd := &cobra.Command{
		Use:               "check",
		Short:             "Check restic backup",
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

			config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigFile)
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
	cmd.Flags().StringVar(&masterURL, "master", masterURL, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	cmd.Flags().StringVar(&kubeconfigFile, "kubeconfig", kubeconfigFile, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	return cmd
}

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := NewRootCmd().Execute(); err != nil {
		log.Fatalln("Error in Stash Main:", err)
	}
}
