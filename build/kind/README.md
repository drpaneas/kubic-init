# Using `kind` with `kubic-init`

Usage:

1) (optional) for a 2-nodes configuration, create a YAML file like this one: 

```yaml
    # multi-node.yaml example
    # technically valid, minimal config file with two nodes
    kind: Config
    apiVersion: kind.sigs.k8s.io/v1alpha2
    nodes:
    - role: control-plane
    - role: worker
```

2) assuming you have a checkout of the kubernetes sources in 
`GOPATH/k8s.io/kubernetes`, you can build the images with:
    ```sh
    $ make kind-images
    ```
   
    but you can use your own checkout of kubernetes with:
    ```sh
    $ make kind-images NODE_BUILD_ARGS="--kube-root $GOPATH/src/k8s.io/kubernetes"
    ```
    

3) create the cluster with

    ```sh
    $ kind create cluster --loglevel debug --image kindest/node:kubic
    ```

   (adding `--config multi-node.yaml` for using the 2-nodes config)
