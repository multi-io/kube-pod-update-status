package main

import (
    "flag"
    "fmt"
    "github.com/golang/glog"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "os"
)

var (
    kubeconfigPath string
)


func main() {
    flag.StringVar(&kubeconfigPath, "kubeconfigPath", "", "Path to a kubeconfigPath.")

    flag.Parse()

    if kubeconfigPath == "" {
        kubeconfigPath = os.Getenv("KUBECONFIG")
    }

    kubeconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
    if err != nil {
        glog.Fatalf("error building kubeconfigPath: %v", err)
    }

    kubeClient, err := kubernetes.NewForConfig(kubeconfig)
    if err != nil {
        glog.Fatalf("error building kubernetes clientset for kubeClient: %v", err)
    }

    fmt.Printf("got kubeClient: %v", kubeClient)
}
