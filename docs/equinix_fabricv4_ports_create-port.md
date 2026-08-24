## equinix fabricv4 ports create-port

Create Port

### Synopsis

Creates Equinix Fabric™ Port.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports create-port [flags]
```

### Options

```
      --dry-run                                                         dry-run field
  -h, --help                                                            help for create-port
      --port-request-account-account-number int                         Account number
      --port-request-account-additional-properties string               port-request-account-additional-properties (JSON)
      --port-request-additional-properties string                       port-request-additional-properties (JSON)
      --port-request-bandwidth int                                      Equinix assigned response attribute for Port bandwidth in Mbps
                                                                        Deprecated
      --port-request-bmmr-type string                                   port-request-bmmr-type
      --port-request-connectivity-source-type string                    port-request-connectivity-source-type
      --port-request-demarcation-point-additional-properties string     port-request-demarcation-point-additional-properties (JSON)
      --port-request-demarcation-point-cabinet-unique-space-id string   Port cabinet unique space id
      --port-request-demarcation-point-cage-unique-space-id string      Port cage unique space id
      --port-request-demarcation-point-connector-type string            Port connector type
      --port-request-demarcation-point-ibx string                       A-side/Equinix ibx
      --port-request-demarcation-point-patch-panel string               Port patch panel
      --port-request-demarcation-point-patch-panel-name string          Port patch panel
                                                                        Deprecated
      --port-request-demarcation-point-patch-panel-port-a string        Port patch panel port A
      --port-request-demarcation-point-patch-panel-port-b string        Port patch panel port B
      --port-request-encapsulation-additional-properties string         port-request-encapsulation-additional-properties (JSON)
      --port-request-encapsulation-tag-protocol-id string               Port encapsulation tag protocol identifier
      --port-request-encapsulation-type string                          port-request-encapsulation-type
      --port-request-lag-enabled                                        Indicates whether Link Aggregation Group (LAG) is enabled on this port
      --port-request-loas string                                        Port Loas (JSON array)
      --port-request-location-additional-properties string              port-request-location-additional-properties (JSON)
      --port-request-location-metro-code string                         port-request-location-metro-code
      --port-request-notifications string                               Notification preferences (JSON array)
      --port-request-order-additional-properties string                 port-request-order-additional-properties (JSON)
      --port-request-order-customer-reference-id string                 Customer order reference Id
      --port-request-order-purchase-order string                        port-request-order-purchase-order (JSON)
      --port-request-order-signature string                             port-request-order-signature (JSON)
      --port-request-package-additional-properties string               port-request-package-additional-properties (JSON)
      --port-request-package-code string                                port-request-package-code
      --port-request-physical-ports string                              Physical ports that implement this port (JSON array)
      --port-request-physical-ports-count int                           port-request-physical-ports-count
      --port-request-physical-ports-speed int                           Physical Ports Speed in Mbps
      --port-request-physical-ports-type string                         port-request-physical-ports-type
      --port-request-project-additional-properties string               port-request-project-additional-properties (JSON)
      --port-request-project-project-id string                          Subscriber-assigned project ID
      --port-request-redundancy-additional-properties string            port-request-redundancy-additional-properties (JSON)
      --port-request-redundancy-enabled                                 Access point redundancy
      --port-request-redundancy-group string                            Port UUID of respective primary port
                                                                        Deprecated
      --port-request-redundancy-priority string                         port-request-redundancy-priority
      --port-request-service-code string                                port-request-service-code
      --port-request-service-type string                                Deprecated
      --port-request-settings-additional-properties string              port-request-settings-additional-properties (JSON)
      --port-request-settings-buyout                                    Deprecated
      --port-request-settings-layer3-enabled                            Deprecated
      --port-request-settings-package-type string                       Deprecated
      --port-request-settings-place-vc-order-permission                 Deprecated
      --port-request-settings-view-port-permission                      Deprecated
      --port-request-tether-ibx string                                  z-side/Equinix ibx
      --port-request-type string                                        port-request-type
      --request string                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

