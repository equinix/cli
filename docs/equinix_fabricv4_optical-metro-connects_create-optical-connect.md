## equinix fabricv4 optical-metro-connects create-optical-connect

Create Optical Metro Connect Service

### Synopsis

Create a single Optical Metro Connect circuit.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 optical-metro-connects create-optical-connect [flags]
```

### Options

```
  -h, --help                                                                   help for create-optical-connect
      --optical-connect-post-request-a-side-additional-properties string       optical-connect-post-request-a-side-additional-properties (JSON)
      --optical-connect-post-request-a-side-connector-type string              optical-connect-post-request-a-side-connector-type
      --optical-connect-post-request-a-side-patch-panel-id string              Unique identifier of the patch panel.
      --optical-connect-post-request-a-side-patch-panel-port-a string          Specify the desired port number for Port A. <br> When ports are not provided, next available ports will be used.
      --optical-connect-post-request-a-side-patch-panel-port-b string          Specify the desired port number for Port B. <br> When ports are not provided, next available ports will be used. <br> Required for Connector type FC, SC and ST only.
      --optical-connect-post-request-account-account-name string               Account name
      --optical-connect-post-request-account-account-number int                Account number
      --optical-connect-post-request-account-additional-properties string      optical-connect-post-request-account-additional-properties (JSON)
      --optical-connect-post-request-account-global-cust-id string             Account name
      --optical-connect-post-request-account-global-org-id string              Global organization identifier
      --optical-connect-post-request-account-global-organization-name string   Global organization name
      --optical-connect-post-request-account-org-id int                        Customer organization identifier
      --optical-connect-post-request-account-organization-name string          Customer organization name
      --optical-connect-post-request-account-reseller-account-name string      Reseller account name
      --optical-connect-post-request-account-reseller-account-number int       Reseller account number
      --optical-connect-post-request-account-reseller-org-id int               Reseller customer organization identifier
      --optical-connect-post-request-account-reseller-ucm-id string            Reseller account ucmId
      --optical-connect-post-request-account-ucm-id string                     Account ucmId
      --optical-connect-post-request-additional-properties string              optical-connect-post-request-additional-properties (JSON)
      --optical-connect-post-request-bandwidth int                             Connection bandwidth Mbps. <br> Available bandwidths depend on the IBX pair. <br> 1000 - 1 Gbps. <br> 10000 - 10 Gbps. <br> 100000 - 100 Gbps. <br>
      --optical-connect-post-request-bmmr-type string                          optical-connect-post-request-bmmr-type
      --optical-connect-post-request-connection-destination-type string        optical-connect-post-request-connection-destination-type
      --optical-connect-post-request-notifications string                      Contacts to notify about connection configuration and status changes (JSON array)
      --optical-connect-post-request-order-additional-properties string        optical-connect-post-request-order-additional-properties (JSON)
      --optical-connect-post-request-order-customer-reference-id string        Your own reference for this connection.
      --optical-connect-post-request-order-order-number string                 Equinix order reference number.
      --optical-connect-post-request-order-purchase-order-number string        Purchase order number reference.
      --optical-connect-post-request-path-type string                          optical-connect-post-request-path-type
      --optical-connect-post-request-redundancy-additional-properties string   optical-connect-post-request-redundancy-additional-properties (JSON)
      --optical-connect-post-request-redundancy-group string                   Redundancy group identifier
      --optical-connect-post-request-redundancy-priority string                optical-connect-post-request-redundancy-priority
      --optical-connect-post-request-type string                               optical-connect-post-request-type
      --optical-connect-post-request-z-side-additional-properties string       optical-connect-post-request-z-side-additional-properties (JSON)
      --optical-connect-post-request-z-side-connector-type string              optical-connect-post-request-z-side-connector-type
      --optical-connect-post-request-z-side-loa string                         optical-connect-post-request-z-side-loa (JSON)
      --optical-connect-post-request-z-side-location string                    optical-connect-post-request-z-side-location (JSON)
      --optical-connect-post-request-z-side-patch-panel-id string              Unique identifier of the patch panel.
      --optical-connect-post-request-z-side-patch-panel-port-a string          Specify the desired port number for Port A. <br> When ports are not provided, next available ports will be used.
      --optical-connect-post-request-z-side-patch-panel-port-b string          Specify the desired port number for Port B. <br> When ports are not provided, next available ports will be used. <br> Required for Connector type FC, SC and ST only.
      --request string                                                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 optical-metro-connects](equinix_fabricv4_optical-metro-connects.md)	 - Manage optical-metro-connects resources

