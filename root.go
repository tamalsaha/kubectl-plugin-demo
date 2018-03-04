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
	var (
		enableAnalytics = true
	)
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
		flags.Set(clientcmd.FlagClusterName, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagClusterName, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagAuthInfoName, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagAuthInfoName, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagContext, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagContext, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagNamespace, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagNamespace, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagAPIServer, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagAPIServer, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagInsecure, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagInsecure, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagCertFile, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagCertFile, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagKeyFile, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagKeyFile, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagCAFile, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagCAFile, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagBearerToken, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagBearerToken, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagImpersonate, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagImpersonate, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagImpersonateGroup, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagImpersonateGroup, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagUsername, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagUsername, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagPassword, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagPassword, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
		flags.Set(clientcmd.FlagTimeout, os.Getenv(plugins.FlagToEnvName(clientcmd.FlagTimeout, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
	}
	// rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	rootCmd.PersistentFlags().BoolVar(&enableAnalytics, "analytics", enableAnalytics, "Send analytical events to Google Analytics")

	rootCmd.AddCommand(NewCmdListNodes(clientConfig))
	rootCmd.AddCommand(NewCmdEnv())
	rootCmd.AddCommand(NewCmdInstall(rootCmd))
	return rootCmd
}
