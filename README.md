# Introduction

A skeleton for writing `kubectl` [exec credential plugins](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins) in Go that includes:
1. Basic logging configuration.
1. Outputting an appropriately formatted JSON to `stdout`.
1. A `wrap.sh` script to directly run your code as a plugin during development.

To start developing ensure your `~/.kube/config`<sup>1</sup> file configures a context with a user as follows:

```yaml
apiVersion: v1
clusters:
- cluster:
    server: https://
  name: my-cluster
contexts:
- context:
    cluster: my-cluster
    user: my-user
  name: my-context
current-context: my-context
kind: Config
preferences: {}
users:
- name: my-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      args:
      - arg1
      - arg2
      command: /absolute/path/of/wrap.sh
      env:
      - name: FOO
        value: bar
      provideClusterInfo: true
```

and then run `kubectl get pods`. Modify [cmd/main.go](cmd/main.go) to have the plugin logic specific to your environment and clusters (i.e. one could return a token from [DefaultTokenSource](https://pkg.go.dev/golang.org/x/oauth2@v0.0.0-20210220000619-9bb904979d93/google#DefaultTokenSource) for Google Kubernetes Engine clusters).

<sup>1</sup> If the `KUBECONFIG` environment variable is set, the configuration should instead be added to one of the files listed in `KUBECONFIG`.
