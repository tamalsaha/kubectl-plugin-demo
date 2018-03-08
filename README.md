# kubectl-plugin-demo

```
$ go get -u github.com/tamalsaha/kubectl-plugin-demo && kubectl-plugin-demo install
```

This will copy the plugin binary and write a plugin.yaml file in the `~/.kube/plugins/kubectl-plugin-demo` directory.

Now you can use it like below:

```
$ kubectl plugin kubectl-plugin-demo
```
