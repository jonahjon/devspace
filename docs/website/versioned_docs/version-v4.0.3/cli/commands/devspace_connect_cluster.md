---
title: Command - devspace connect cluster
sidebar_label: cluster
id: version-v4.0.3-devspace_connect_cluster
original_id: devspace_connect_cluster
---


Connects an existing cluster to DevSpace Cloud

## Synopsis


```
devspace connect cluster [flags]
```

```
#######################################################
############ devspace connect cluster #################
#######################################################
Connects an existing cluster to DevSpace Cloud.

Examples:
devspace connect cluster 
#######################################################
```
## Options

```
      --admission-controller   Deploy the admission controller (default true)
      --cert-manager           Deploy a cert manager (default true)
      --context string         The kube context to use
      --domain string          The domain to use
  -h, --help                   help for cluster
      --ingress-controller     Deploy an ingress controller (default true)
      --key string             The encryption key to use
      --name string            The cluster name to create
      --open-ui                Opens the UI and displays the cluster overview
      --provider string        The cloud provider to use
      --use-domain             Use an automatic domain for the cluster
      --use-hostnetwork        Use the host network for the ingress controller instead of a loadbalancer
```

### Options inherited from parent commands

```
      --debug                 Prints the stack trace if an error occurs
      --kube-context string   The kubernetes context to use
  -n, --namespace string      The kubernetes namespace to use
      --no-warn               If true does not show any warning when deploying into a different namespace or kube-context than before
  -p, --profile string        The devspace profile to use (if there is any)
      --silent                Run in silent mode and prevents any devspace log output except panics & fatals
  -s, --switch-context        Switches and uses the last kube context and namespace that was used to deploy the DevSpace project
      --var strings           Variables to override during execution (e.g. --var=MYVAR=MYVALUE)
```

## See Also

* [devspace connect](../../cli/commands/devspace_connect)	 - Connect an external cluster to devspace cloud
