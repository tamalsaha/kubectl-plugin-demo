package main

import (
	"flag"
	stdlog "log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCmd() *cobra.Command {
	var (
		enableAnalytics = true
	)
	var rootCmd = &cobra.Command{
		Use:               "tamal",
		Short:             `Tamal's kubectl plugin'`,
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

	rootCmd.AddCommand(NewCmdListNodes())
	rootCmd.AddCommand(NewCmdEnv())
	rootCmd.AddCommand(NewCmdInstall(rootCmd))
	return rootCmd
}
