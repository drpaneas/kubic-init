![alpha](https://img.shields.io/badge/stability%3F-beta-yellow.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubic-project/kubic-init)](https://goreportcard.com/report/github.com/kubic-project/kubic-init)  [![GoDoc](https://godoc.org/github.com/kubic-project/kubic-init?status.svg)](https://godoc.org/github.com/kubic-project/kubic-init) [![CircleCI](https://circleci.com/gh/kubic-project/kubic-init/tree/master.svg?style=svg)](https://circleci.com/gh/kubic-project/kubic-init/tree/master)

# Kubic-init:

The Kubic-init repository is a project geared at simplifying the process of creating and maintaining kubernetes clusters. Kubic-init is being actively developed/tested on top of the [Kubic project](https://en.opensuse.org/Portal:Kubic).

- [User-Guide-documentation](docs/user-guide.md)
- [Design And Architecture](docs/design.md)
- [Devel](CONTRIBUTING.md)
- [Roadmap](docs/roadmap.md)

___
#### Kubic operators:

On top of a cluster bootstrapped with kubic-init, you can add following operators:

* [External authentication](https://github.com/kubic-project/dex-operator/blob/master/README.md)
* [Adding private registries](https://github.com/kubic-project/registries-operator/blob/master/README.md)
