## equinix fabricv4 cloud-routers create-cloud-router-action

Create Route Table Action

### Synopsis

This API provides capability to refresh route table and bgp session summary information

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers create-cloud-router-action [flags]
```

### Options

```
  -h, --help               help for create-cloud-router-action
      --request string     JSON payload for request body. Available fields: cloud-router-action-request (CloudRouterActionRequest)
      --router-id string   Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

