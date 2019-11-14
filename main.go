package main

import (
    "flag"
    "fmt"
    "github.com/golang/glog"
    "k8s.io/client-go/tools/clientcmd"
)

var (
    kubeconfig                       string
)


func main() {
    flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig.")

    flag.Parse()

    cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        glog.Fatalf("error building kubeconfig: %v", err)
    }

    fmt.Printf("get kubeconfig: %v", cfg)
}
