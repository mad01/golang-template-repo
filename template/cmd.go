package main

import (
	"fmt"

	"time"

	"github.com/spf13/cobra"
)

func cmdVersion() *cobra.Command {
	var command = &cobra.Command{
		Use:   "version",
		Short: "get version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getVersion())
		},
	}
	return command
}

func cmdRunController() *cobra.Command {
	var kubeconfig string
	var verbose bool
	var interval time.Duration
	var command = &cobra.Command{
		Use:   "controller",
		Short: "run the controller",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			spewInit()
			initLog(verbose)

			kube := newKube()

			stopCh := make(chan struct{})
			defer close(stopCh)

			controller := newController(kube, interval)
			controller.Run(stopCh)
		},
	}
	command.Flags().StringVarP(&kubeconfig, "kube.config", "k", "", "outside cluster path to kube config")
	command.Flags().DurationVarP(&interval, "interval.controller", "i", 10*time.Second, "controller update interaval for internal k8s caches")
	command.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	return command
}

func runCmd() error {
	var rootCmd = &cobra.Command{Use: "{{GITHUB_REPO}}"}
	rootCmd.AddCommand(cmdVersion())
	rootCmd.AddCommand(cmdRunController())

	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	return nil
}
