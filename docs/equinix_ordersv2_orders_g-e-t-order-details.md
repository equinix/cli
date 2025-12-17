## equinix ordersv2 orders g-e-t-order-details

Retrieve an order

### Synopsis

This method retrieves an order by its ID for a user with permission to view this request.

Use --request flag to provide optional JSON payload fields.

```
equinix ordersv2 orders g-e-t-order-details [flags]
```

### Options

```
  -h, --help              help for g-e-t-order-details
      --ibxs string       ibxs field (JSON or string)
      --order-id string   Identifier of the Order (required)
      --request string    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix ordersv2 orders](equinix_ordersv2_orders.md)	 - Manage orders resources

