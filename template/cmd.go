package main

import (
	"fmt"

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
	var kubeconfig, namespace string
	var verbose bool
	var port int
	var command = &cobra.Command{
		Use:   "controller",
		Short: "start the controller",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			initLog(verbose)

			kube := newKube(kubeconfig)

			controller := newInformer(kube, port)
			controller.Run()
		},
	}
	command.Flags().StringVarP(
		&kubeconfig,
		"kube.config", "k", "",
		"only needed with running outside cluster, path to kube config")
	command.Flags().StringVarP(
		&namespace,
		"namespace", "n",
		"default", "ns where the service accounts and cluster role bindings is created",
	)
	command.Flags().IntVarP(&port, "http.port", "p", 8080, "port to expose service on")
	command.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	return command
}

func runCmd() {
	var rootCmd = &cobra.Command{Use: "k8s-metadata"}
	rootCmd.AddCommand(cmdVersion())
	rootCmd.AddCommand(cmdRunController())

	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err.Error())
	}
	return
}
