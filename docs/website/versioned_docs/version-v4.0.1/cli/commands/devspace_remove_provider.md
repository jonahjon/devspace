---
title: Command - devspace remove provider
sidebar_label: provider
id: version-v4.0.1-devspace_remove_provider
original_id: devspace_remove_provider
---


Removes a cloud provider from the configuration

## Synopsis


```
devspace remove provider [flags]
```

```
#######################################################
############ devspace remove provider #################
#######################################################
Removes a cloud provider.

Example:
devspace remove provider app.devspace.cloud
#######################################################
```
## Options

```
  -h, --help          help for provider
      --name string   Cloud provider name to use
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

* [devspace remove](../../cli/commands/devspace_remove)	 - Changes devspace configuration
