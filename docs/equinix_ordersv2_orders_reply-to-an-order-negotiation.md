## equinix ordersv2 orders reply-to-an-order-negotiation

Reply to an order negotiation

### Synopsis

This method approves or rejects an order negotiation for a user with permission to view this request.

Use --request flag to provide optional JSON payload fields.

```
equinix ordersv2 orders reply-to-an-order-negotiation [flags]
```

### Options

```
  -h, --help                                                help for reply-to-an-order-negotiation
      --negotiations-request-action string                  negotiations-request-action
      --negotiations-request-additional-properties string   negotiations-request-additional-properties (JSON)
      --negotiations-request-reason string                  Reason for cancelling the negotiation
      --negotiations-request-reference-id string            Unique identifier to reference specific activity or order line id.
      --order-id string                                     Identifier of the order (required)
      --request string                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix ordersv2 orders](equinix_ordersv2_orders.md)	 - Manage orders resources

