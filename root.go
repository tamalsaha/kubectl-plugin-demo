package main

import (
	"flag"
	stdlog "log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/plugins"
)

func NewRootCmd(plugin bool) *cobra.Command {
	//var (
	//	enableAnalytics = true
	//)
	var rootCmd = &cobra.Command{
		Use:               "kubectl-plugin-demo",
		Short:             `Tamal's kubectl plugin'`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.Flags().VisitAll(func(flag *pflag.Flag) {
				stdlog.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			})
		},
	}
	flags := rootCmd.PersistentFlags()
	clientConfig := util.DefaultClientConfig(flags)
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	if plugin {
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagClusterName)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagAuthInfoName)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagContext)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagNamespace)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagAPIServer)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagInsecure)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagCertFile)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagKeyFile)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagCAFile)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagBearerToken)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagImpersonate)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagImpersonateGroup)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagUsername)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagPassword)
		processKubectlFlag(rootCmd.PersistentFlags(), clientcmd.FlagTimeout)
	}
	// rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	// rootCmd.PersistentFlags().BoolVar(&enableAnalytics, "analytics", enableAnalytics, "Send analytical events to Google Analytics")

	rootCmd.AddCommand(NewCmdListNodes(clientConfig))
	rootCmd.AddCommand(NewCmdEnv())
	rootCmd.AddCommand(NewCmdInstall(rootCmd))
	return rootCmd
}

func processKubectlFlag(flags *pflag.FlagSet, name string) {
	flags.Set(name, os.Getenv(plugins.FlagToEnvName(name, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
	flags.MarkHidden(name)
}
