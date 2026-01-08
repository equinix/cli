## equinix securecabinetv1 orders create-order

Order a new Secure Cabinet deployment

### Synopsis

Order a new Secure Cabinet deployment and recommended additional products at IBX data centers worldwide.

Use --request flag to provide optional JSON payload fields.

```
equinix securecabinetv1 orders create-order [flags]
```

### Options

```
  -h, --help                                                                  help for create-order
      --order-create-request-account-number string                            Identifier of a billing account, with permission to order colocation products.
      --order-create-request-additional-properties string                     order-create-request-additional-properties (JSON)
      --order-create-request-contract-term string                             order-create-request-contract-term
      --order-create-request-customer-reference string                        Supplementary identifier for your discretionary use. For example, your internal identifier.
      --order-create-request-end-customer-name string                         End customer name. Applicable and required if a reseller billing account is used.
      --order-create-request-ibx-code string                                  IBX data center identifier.
      --order-create-request-order-item-additional-properties string          order-create-request-order-item-additional-properties (JSON)
      --order-create-request-order-item-cabinet-dimensions string             order-create-request-order-item-cabinet-dimensions (JSON)
      --order-create-request-order-item-draw-capacity float                   Maximum, combined power draw of all cabinets in your order, measured in kVA. Applicable values, in 0.5 increments, depend on the IBX data center.
      --order-create-request-order-item-fabric-port                           Indicates if a single, primary Fabric port should be included and delivered to one of the ordered cabinets.
      --order-create-request-order-item-number-of-cabinets int                The number of ordered cabinets.
      --order-create-request-order-item-pdus                                  Indicates if an Equinix-recommended set of two PDUs, for single-phase circuit, per cabinet should be installed.
      --order-create-request-purchase-order-number string                     Purchase order number. Optional, unless the billing account requires it.
      --order-create-request-technical-contact-additional-properties string   order-create-request-technical-contact-additional-properties (JSON)
      --order-create-request-technical-contact-email string                   Email address.
      --order-create-request-technical-contact-first-name string              First name.
      --order-create-request-technical-contact-last-name string               Last name.
      --order-create-request-technical-contact-phone string                   order-create-request-technical-contact-phone (JSON)
      --request string                                                        JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix securecabinetv1 orders](equinix_securecabinetv1_orders.md)	 - Manage orders resources

