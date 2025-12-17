## equinix orderhistoryv1 retrieve-orders g-e-t-retrieve-orders-locations

Retrieve order permissible IBX locations

### Synopsis

The API returns locations (IBXs and Cages) where the user's organization has a presence. The user may be able to view orders related to these locations, depending on their permissions.

Use --request flag to provide optional JSON payload fields.

```
equinix orderhistoryv1 retrieve-orders g-e-t-retrieve-orders-locations [flags]
```

### Options

```
      --authorization string   authorization field
  -h, --help                   help for g-e-t-retrieve-orders-locations
      --request string         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix orderhistoryv1 retrieve-orders](equinix_orderhistoryv1_retrieve-orders.md)	 - Manage retrieve-orders resources

