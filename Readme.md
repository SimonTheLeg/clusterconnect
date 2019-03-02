# Clusterconnect Package

Simple Package for connecting to a Kubernetes Cluster, because I am tired of copy-pasting the same lines over and over :)

## Usage

ConnectToCluster() returns a Kubernetes Clientset. You can specify whether to use the local kubeconfig file or an in-cluster-config by passing the `--local` flag when running your program.