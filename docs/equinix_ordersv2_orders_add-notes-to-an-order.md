## equinix ordersv2 orders add-notes-to-an-order

Add notes to an order

### Synopsis

This method adds notes to an order by its ID for a user with permission to view this request.

Use --request flag to provide optional JSON payload fields.

```
equinix ordersv2 orders add-notes-to-an-order [flags]
```

### Options

```
  -h, --help                                        help for add-notes-to-an-order
      --note-request-additional-properties string   note-request-additional-properties (JSON)
      --note-request-attachments string             File(s) attached to the Orders. To learn about including attachments in your request, see POST Attachments API. (JSON array)
      --note-request-reference-id string            Unique reference ID associated with notes.
      --note-request-text string                    The text of the note
      --order-id string                             Identifier of the order (required)
      --request string                              JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix ordersv2 orders](equinix_ordersv2_orders.md)	 - Manage orders resources

