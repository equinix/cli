## equinix orderhistoryv1 retrieve-orders p-o-s-t-orders-history

Search Orders History

### Synopsis

Based on filtering criteria, this method returns a list of orders from the last 12 months for IBX locations where the user has permissions.

Use --request flag to provide optional JSON payload fields.

```
equinix orderhistoryv1 retrieve-orders p-o-s-t-orders-history [flags]
```

### Options

```
      --authorization string                        authorization field
      --body-additional-properties string           body-additional-properties (JSON)
      --body-filters-additional-properties string   body-filters-additional-properties (JSON)
      --body-filters-date-range string              body-filters-date-range
      --body-filters-from-date string               Date Format Should be mm/dd/yyyy.<br/> Not applicable when dateRange is provided.
      --body-filters-ibxs string                    ibxs filter (JSON array)
      --body-filters-order-status string            order status filter (JSON array)
      --body-filters-product-types string           Product(order type) filter (JSON array)
      --body-filters-to-date string                 Date Format Should be mm/dd/yyyy.<br/> Not applicable when dateRange is provided
      --body-page-additional-properties string      body-page-additional-properties (JSON)
      --body-page-number int                        Page number indexed from 0.
      --body-page-size int                          Page Size.
      --body-q string                               Query value to be queried against source values(Or operation against all sources).<br/>Supports partial text search
      --body-sorts string                           body-sorts (JSON array)
      --body-source string                          <b>ORDER_NUMBER:</b> Search by order number(1-123456789).<br><b>CUSTOMER_REFERENCE_NUMBER:</b> Search by customer reference number which was entered as part place order.<br><b>TROUBLE_TICKET_NUMBER:</b> Search by trouble ticket numnber(5-123456).<br><b>WORK_ACTIVITY_NUMBER:</b> Search by work order activity number(3-123456). (JSON array)
  -h, --help                                        help for p-o-s-t-orders-history
      --request string                              JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix orderhistoryv1 retrieve-orders](equinix_orderhistoryv1_retrieve-orders.md)	 - Manage retrieve-orders resources

