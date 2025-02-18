// Copyright Contributors to the Open Cluster Management project

package main

import (
	"os"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"open-cluster-management.io/clusteradm/pkg/cmd"
)

func main() {
	streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	root, _ := cmd.NewRootCmd("clusteradm", streams)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
