# kubectl-plugin-demo

```
$ go get -u github.com/tamalsaha/kubectl-plugin-demo && kubectl-plugin-demo install
```

This will copy the plugin binary and write a plugin.yaml file in the `~/.kube/plugins/kubectl-plugin-demo` directory.

Now you can use it like below:

```
$ kubectl plugin kubectl-plugin-demo
```

#### Reading List
- https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/
- https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/plugins
- https://github.com/kubernetes/kubernetes/tree/master/test/fixtures/pkg/kubectl/plugins

#### Known Issues
https://github.com/kubernetes/kubectl/issues/created_by/tamalsaha
