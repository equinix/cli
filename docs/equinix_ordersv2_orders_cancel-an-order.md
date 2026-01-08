## equinix ordersv2 orders cancel-an-order

Cancel an order

### Synopsis

This method cancels an order by its ID for a user with permission to view this request.

Use --request flag to provide optional JSON payload fields.

```
equinix ordersv2 orders cancel-an-order [flags]
```

### Options

```
      --cancel-request-additional-properties string   cancel-request-additional-properties (JSON)
      --cancel-request-attachments string             File(s) attached (JSON array)
      --cancel-request-line-ids lineId                Refers to the lineId of product/service. (JSON array)
      --cancel-request-reason string                  Reason for cancellation
  -h, --help                                          help for cancel-an-order
      --order-id string                               Identifier of the order (required)
      --request string                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix ordersv2 orders](equinix_ordersv2_orders.md)	 - Manage orders resources

