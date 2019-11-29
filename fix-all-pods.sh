#!/usr/bin/env bash

kubectl get pod --all-namespaces -o json | jq -r '.items[] | {"namespace": .metadata.namespace, "name": .metadata.name, "ContainersReady": .status.conditions[] | select(.type|contains("ContainersReady")) | .status, "Ready": .status.conditions[] | select(.type|contains("Ready")) | .status } | select(.Ready=="False" and .ContainersReady=="True") | .namespace + " " + .name' | while read namespace name; do
  echo "processing pod: ${namespace}/${name}" >&2
  if [[ -z "$DRYRUN" ]]; then
    if ! ./kube-pod-update-status --namespace "$namespace" "$name"; then
      echo "failed once, trying again" >&2
      if ! ./kube-pod-update-status --namespace "$namespace" "$name"; then
        echo "FAILED FOR POD: ${namespace}/${name}" >&2
      fi
    fi
  fi
done
