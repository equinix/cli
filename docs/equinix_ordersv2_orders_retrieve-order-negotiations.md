## equinix ordersv2 orders retrieve-order-negotiations

Retrieve order negotiations

### Synopsis

This method retrieves order negotiations by its ID for a user with permission to view this request.

Use --request flag to provide optional JSON payload fields.

```
equinix ordersv2 orders retrieve-order-negotiations [flags]
```

### Options

```
  -h, --help              help for retrieve-order-negotiations
      --order-id string   Identifier of the order (required)
      --request string    JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix ordersv2 orders](equinix_ordersv2_orders.md)	 - Manage orders resources

