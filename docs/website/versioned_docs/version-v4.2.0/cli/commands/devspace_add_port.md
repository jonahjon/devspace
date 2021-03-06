---
title: Command - devspace add port
sidebar_label: port
id: version-v4.2.0-devspace_add_port
original_id: devspace_add_port
---


Add a new port forward configuration

## Synopsis


```
devspace add port [flags]
```

```
#######################################################
################ devspace add port ####################
#######################################################
Add a new port mapping to this project's devspace.yaml

Format is port(:remotePort) comma separated, e.g.
devspace add port 8080:80,3000
#######################################################
```
## Options

```
  -h, --help                    help for port
      --label-selector string   Comma separated key=value label-selector list (e.g. release=test)
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

* [devspace add](../../cli/commands/devspace_add)	 - Convenience command: adds something to devspace.yaml
