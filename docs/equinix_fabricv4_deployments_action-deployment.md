## equinix fabricv4 deployments action-deployment

Deploy, Dry Run or Destroy Deployment

### Synopsis

The deployment action is used to deploy, dry run or destroy a deployment. The request body must contain the type of action to be performed and the connection details. The response will contain the status of the action.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 deployments action-deployment [flags]
```

### Options

```
  -h, --help             help for action-deployment
      --param-1 string   param-1 (required)
      --request string   Raw JSON payload for optional request fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
```

### SEE ALSO

* [equinix fabricv4 deployments](equinix_fabricv4_deployments.md)	 - Manage deployments resources

