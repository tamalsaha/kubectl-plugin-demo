package main

import (
	"flag"
	//stdlog "log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	utilflag "k8s.io/apiserver/pkg/util/flag"
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
			//c.Flags().VisitAll(func(flag *pflag.Flag) {
			//	stdlog.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			//})
		},
	}
	util.NewFactory(nil)

	flags := rootCmd.PersistentFlags()
	flags.AddGoFlagSet(flag.CommandLine)
	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	clientConfig := util.DefaultClientConfig(flags)
	if plugin {
		processKubectlPlugin(rootCmd.PersistentFlags())
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

func processKubectlPlugin(flags *pflag.FlagSet) {
	loadFromEnv(flags, "kubeconfig")
	loadFromEnv(flags, clientcmd.FlagClusterName)
	loadFromEnv(flags, clientcmd.FlagAuthInfoName)
	loadFromEnv(flags, clientcmd.FlagContext)
	loadFromEnv(flags, clientcmd.FlagNamespace)
	loadFromEnv(flags, clientcmd.FlagAPIServer)
	loadFromEnv(flags, clientcmd.FlagInsecure)
	loadFromEnv(flags, clientcmd.FlagCertFile)
	loadFromEnv(flags, clientcmd.FlagKeyFile)
	loadFromEnv(flags, clientcmd.FlagCAFile)
	loadFromEnv(flags, clientcmd.FlagBearerToken)
	loadFromEnv(flags, clientcmd.FlagImpersonate)
	loadFromEnv(flags, clientcmd.FlagImpersonateGroup)
	loadFromEnv(flags, clientcmd.FlagUsername)
	loadFromEnv(flags, clientcmd.FlagPassword)
	loadFromEnv(flags, clientcmd.FlagTimeout)

	loadFromEnv(flags, "alsologtostderr")
	loadFromEnv(flags, "log-backtrace-at")
	loadFromEnv(flags, "log-dir")
	loadFromEnv(flags, "logtostderr")
	loadFromEnv(flags, "stderrthreshold")
	loadFromEnv(flags, "v")
	loadFromEnv(flags, "vmodule")
}

func loadFromEnv(flags *pflag.FlagSet, name string) {
	flags.Set(name, os.Getenv(plugins.FlagToEnvName(name, "KUBECTL_PLUGINS_GLOBAL_FLAG_")))
	flags.MarkHidden(name)
}

func hideKubectlPlugin(flags *pflag.FlagSet) {
	return

	flags.MarkHidden("kubeconfig")
	flags.MarkHidden(clientcmd.FlagClusterName)
	flags.MarkHidden(clientcmd.FlagAuthInfoName)
	flags.MarkHidden(clientcmd.FlagContext)
	flags.MarkHidden(clientcmd.FlagNamespace)
	flags.MarkHidden(clientcmd.FlagAPIServer)
	flags.MarkHidden(clientcmd.FlagInsecure)
	flags.MarkHidden(clientcmd.FlagCertFile)
	flags.MarkHidden(clientcmd.FlagKeyFile)
	flags.MarkHidden(clientcmd.FlagCAFile)
	flags.MarkHidden(clientcmd.FlagBearerToken)
	flags.MarkHidden(clientcmd.FlagImpersonate)
	flags.MarkHidden(clientcmd.FlagImpersonateGroup)
	flags.MarkHidden(clientcmd.FlagUsername)
	flags.MarkHidden(clientcmd.FlagPassword)
	flags.MarkHidden(clientcmd.FlagTimeout)

	flags.MarkHidden("alsologtostderr")
	flags.MarkHidden("log-backtrace-at")
	flags.MarkHidden("log-dir")
	flags.MarkHidden("logtostderr")
	flags.MarkHidden("stderrthreshold")
	flags.MarkHidden("v")
	flags.MarkHidden("vmodule")
}
