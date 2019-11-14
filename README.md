Program that forcefully sets the ready status of a K8s pod to true.

Use at your own risk.

Written in Go. I might just have been to stupid to do the same thing using kubectl --raw on the status subresource.

This is meant to fix pods broken by https://github.com/kubernetes/kubernetes/issues/80968. The fix-all-pods.sh script
identifies all broken pods in a cluster and runs the program on them all. The identification may not work 100%. E.g.
it doesn't take node readiness into account. So again, use with care.

Usage:

```
$ make # build it
$ ./fix-all-pods.sh
processing pod: kube-system/canal-t6z96
F1114 23:49:28.535716   16040 main.go:70] Failed to update status: Operation cannot be fulfilled on pods "canal-t6z96": the object has been modified; please apply your changes to the latest version and try again
failed once, trying again
I1114 23:49:28.618383   16041 main.go:73] Successfully updated pod kube-system/canal-t6z96
processing pod: kube-system/coredns-6c6659d8f7-lr7tl
F1114 23:49:28.753948   16042 main.go:70] Failed to update status: Operation cannot be fulfilled on pods "coredns-6c6659d8f7-lr7tl": the object has been modified; please apply your changes to the latest version and try again
failed once, trying again
I1114 23:49:28.826809   16043 main.go:73] Successfully updated pod kube-system/coredns-6c6659d8f7-lr7tl
processing pod: kube-system/kube-proxy-2rm4v
...
$ ./fix-all-pods.sh  # run it again to check that it does nothing now
$ 
```
