# Kubic-init User Guide:

- [Installation](#installation)
- [Configuration Reference](#configuration-reference)

# Installation

This section contains instructions for installing kubic-init.

# Terraform based deployment:

Have a look [here](https://github.com/kubic-project/kubic-init/tree/master/deployments)

# Manual Installation

## Install a cluster with kubic-init with manual workflow.

This installation will guide you to install a minimal setup of a cluster, having 1 seeder and 1 node.
This is in sake of simplicity. You can have other more complex setups, cluster configurations.

#### Configuring the Kubernetes bootstrap

You can read about the concept of [bootstrap](/design/design-bootstrap.md) with kubic-init.

#### Requirements for Nodes (seeder and others)

Install following pkgs on your Linux Distro

* kubelet
* kubectl

You also need to satisfy `kubeadm` [Requirements](https://kubernetes.io/docs/setup/independent/install-kubeadm/#before-you-begin)

**NOTE**: Although this guide is based on [kubic](https://en.opensuse.org/Portal:Kubic), it should also work with any other GNU/Linux distribution.

#### Installation

##### Seeder Node and others Nodes configuration:

* Copy configuration files needed on seeder and nodes.
 ( https://github.com/kubic-project/kubic-init/tree/master/init)

```bash
curl https://raw.githubusercontent.com/kubic-project/kubic-init/master/init/kubic-init.systemd.conf > /etc/systemd/system/kubic-init.service
curl https://raw.githubusercontent.com/kubic-project/kubic-init/master/init/kubelet-sysctl.conf >  /etc/systctl.d/99-kubernetes-cri.conf
curl https://raw.githubusercontent.com/kubic-project/kubic-init/master/init/kubelet-drop-in.conf >  /etc/systemd/system/kubelet.service.d/kubelet.conf
curl https://raw.githubusercontent.com/kubic-project/kubic-init/master/init/kubic-init.sysconfig >  /etc/sysconfig/kubic-init"

```

* Container runtime

You will need on the `seeder` node,  a container runtime installed. (podman,docker, etc).
The contaner-runtime will execute the kubic-init container.

On All node pull the `kubic-init`, e.g : 
`docker pull registry.opensuse.org/devel/kubic/containers/container/kubic/init:0.0`

* kubic-init.yml configuration (seeder only)

Write the  configuration for kubic-init. into `/etc/kubic/kubic-init.yaml`

Here is an example of minimal configuration.

```yaml
 apiVersion: kubic.suse.com/v1alpha1
 kind: KubicInitConfiguration
 network:
   cni:
     driver: flannel
```

##### enable kubic-init ( on seeder and nodes).

`curl https://raw.githubusercontent.com/kubic-project/kubic-init/master/init/kubic-init.systemd.conf > /etc/systemd/system/kubic-init.service`
`systemctl enable --now kubic-init`


For more elaborate configuration have a look at  - [Configuration Reference](#configuration-reference)

* Node Join

To join a node to cluster you will need to use the `token` provived by kubic-init

___


# Configuration reference

There is an example [kubic-init.yaml](../config/kubic-init.yaml.sample) that has all the possible configurations

# Features explained in details

TODO: this need to be reworked.

#### Bootstrap section

The bootstrap section is reserved to the configuration that is needed in actions that take part before initializing kubic-init

##### Registries

Inside we have the section registries: this section will let you configure mirrors for registries

* multiple mirrors can be set

* each prefix can have multiple mirrors addresses

* each mirror can have certificates configured. This will be used for security to validate the origin of the server.

> certificates: In this configuration you must add the Certificate content, the Fingerprint and the Hash Algorithm that was used.

This is useful for air-gapped environments, where it is not possible to access upstream registries and you have configured a local mirror.

In this scenario, you have to configure the daemon.json file for the container engine to be able to find the initial images, otherwise Kubic would not start if this was not configured.

Example:
 
```yaml
bootstrap:
  registries:
    - prefix: https://registry.suse.com
      mirrors:
        - url: https://airgapped.suse.com
        - url: https://airgapped2.suse.com
          certificate: "-----BEGIN CERTIFICATE-----
  MIIGJzCCBA+gAwIBAgIBATANBgkqhkiG9w0BAQUFADCBsjELMAkGA1UEBhMCRlIx
  DzANBgNVBAgMBkFsc2FjZTETMBEGA1UEBwwKU3RyYXNib3VyZzEYMBYGA1UECgwP
  hnx8SB3sVJZHeer8f/UQQwqbAO+Kdy70NmbSaqaVtp8jOxLiidWkwSyRTsuU6D8i
  DiH5uEqBXExjrj0FslxcVKdVj5glVcSmkLwZKbEU1OKwleT/iXFhvooWhQ==
  -----END CERTIFICATE-----"
          fingerprint: "E8:73:0C:C5:84:B1:EB:17:2D:71:54:4D:89:13:EE:47:36:43:8D:BF:5D:3C:0F:5B:FC:75:7E:72:28:A9:7F:73"
          hashalgorithm: "SHA256"
    - prefix: https://registry.io
      mirrors:
        - url: https://airgapped.suse.com
        - url: https://airgapped2.suse.com
          certificate: "-----BEGIN CERTIFICATE-----
  MIIGJzCCBA+gAwIBAgIBATANBgkqhkiG9w0BAQUFADCBsjELMAkGA1UEBhMCRlIx
  DzANBgNVBAgMBkFsc2FjZTETMBEGA1UEBwwKU3RyYXNib3VyZzEYMBYGA1UECgwP
  hnx8SB3sVJZHeer8f/UQQwqbAO+Kdy70NmbSaqaVtp8jOxLiidWkwSyRTsuU6D8i
  DiH5uEqBXExjrj0FslxcVKdVj5glVcSmkLwZKbEU1OKwleT/iXFhvooWhQ==<
  -----END CERTIFICATE-----"
          fingerprint: "E8:73:0C:C5:84:B1:EB:17:2D:71:54:4D:89:13:EE:47:36:43:8D:BF:5D:3C:0F:5B:FC:75:7E:72:28:A9:7F:73"
          hashalgorithm: "SHA256"
```
