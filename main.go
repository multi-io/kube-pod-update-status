package main

import (
    "flag"
    v1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/klog"
    "os"
)

func main() {
    var (
        kubeconfigPath string
        namespace string
    )

    flag.StringVar(&kubeconfigPath, "kubeconfigPath", "", "Path to a kubeconfigPath.")
    flag.StringVar(&namespace, "namespace", "", "K8s namespace (default: default)")

    flag.Parse()

    if kubeconfigPath == "" {
        kubeconfigPath = os.Getenv("KUBECONFIG")
    }

    if namespace == "" {
        namespace = "default"
    }

    podName := flag.Arg(0)
    if podName == "" {
        klog.Fatalf("podUpdate name not specified")
    }

    kubeconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
    if err != nil {
        klog.Fatalf("error getting kubeconfig: %v", err)
    }

    kubeClient, err := kubernetes.NewForConfig(kubeconfig)
    if err != nil {
        klog.Fatalf("error building kubernetes client: %v", err)
    }

    pod, err := kubeClient.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
    if err != nil {
        klog.Fatalf("error getting pod %s/%s: %v", namespace, podName, err)
    }

    // copied (incl. dependencies) and adapted from https://github.com/kubernetes/kubernetes/blob/v1.15.5/pkg/controller/util/node/controller_utils.go#L122
    podUpdate := pod.DeepCopy()
    for _, cond := range podUpdate.Status.Conditions {
        if cond.Type == v1.PodReady {
            cond.Status = v1.ConditionTrue
            if !UpdatePodCondition(&podUpdate.Status, &cond) {
                break
            }
            klog.V(2).Infof("Updating ready status of pod %v to true", podUpdate.Name)
            _, err := kubeClient.CoreV1().Pods(podUpdate.Namespace).UpdateStatus(podUpdate)
            if err != nil {
                klog.Fatalf("Failed to update status: %v", err)
            }
            break
        }
    }

    if _, err := kubeClient.CoreV1().Pods(podUpdate.Namespace).UpdateStatus(podUpdate); err != nil {
        klog.Fatalf("Failed to update status: %v", err)
    }

    klog.Infof("Successfully updated pod %s/%s", namespace, podName)
}
