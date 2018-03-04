package main

import "os"

type GlobalOptions struct {
	Namespace string
}

func Load() GlobalOptions {
	var opts GlobalOptions


	opts.Namespace = os.Getenv("KUBECTL_PLUGINS_CURRENT_NAMESPACE")
}
