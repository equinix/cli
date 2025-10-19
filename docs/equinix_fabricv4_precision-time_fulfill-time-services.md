## equinix fabricv4 precision-time fulfill-time-services

Configure Service.

### Synopsis

The API provides capability to Configure/Fulfill the Precision Time Service.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 precision-time fulfill-time-services [flags]
```

### Options

```
  -h, --help                                                                                     help for fulfill-time-services
      --precision-time-service-request-additional-properties string                              precision-time-service-request-additional-properties (required) (JSON)
      --precision-time-service-request-connections string                                        precision-time-service-request-connections (required) (JSON array)
      --precision-time-service-request-ipv4-additional-properties string                         precision-time-service-request-ipv4-additional-properties (required) (JSON)
      --precision-time-service-request-ipv4-default-gateway string                               precision-time-service-request-ipv4-default-gateway
      --precision-time-service-request-ipv4-network-mask string                                  precision-time-service-request-ipv4-network-mask (required)
      --precision-time-service-request-ipv4-primary string                                       precision-time-service-request-ipv4-primary (required)
      --precision-time-service-request-ipv4-secondary string                                     precision-time-service-request-ipv4-secondary (required)
      --precision-time-service-request-name string                                               precision-time-service-request-name (required)
      --precision-time-service-request-ntp-advanced-configuration string                         precision-time-service-request-ntp-advanced-configuration (JSON array)
      --precision-time-service-request-order-additional-properties string                        precision-time-service-request-order-additional-properties (JSON)
      --precision-time-service-request-order-customer-reference-number string                    precision-time-service-request-order-customer-reference-number
      --precision-time-service-request-order-order-number string                                 precision-time-service-request-order-order-number
      --precision-time-service-request-order-purchase-order-number string                        precision-time-service-request-order-purchase-order-number
      --precision-time-service-request-package-additional-properties string                      precision-time-service-request-package-additional-properties (required) (JSON)
      --precision-time-service-request-package-code string                                       precision-time-service-request-package-code (required)
      --precision-time-service-request-project-additional-properties string                      precision-time-service-request-project-additional-properties (JSON)
      --precision-time-service-request-project-project-id string                                 precision-time-service-request-project-project-id
      --precision-time-service-request-ptp-advanced-configuration-additional-properties string   precision-time-service-request-ptp-advanced-configuration-additional-properties (JSON)
      --precision-time-service-request-ptp-advanced-configuration-domain int                     precision-time-service-request-ptp-advanced-configuration-domain
      --precision-time-service-request-ptp-advanced-configuration-grant-time int                 precision-time-service-request-ptp-advanced-configuration-grant-time
      --precision-time-service-request-ptp-advanced-configuration-log-announce-interval int      precision-time-service-request-ptp-advanced-configuration-log-announce-interval
      --precision-time-service-request-ptp-advanced-configuration-log-delay-req-interval int     precision-time-service-request-ptp-advanced-configuration-log-delay-req-interval
      --precision-time-service-request-ptp-advanced-configuration-log-sync-interval int          precision-time-service-request-ptp-advanced-configuration-log-sync-interval
      --precision-time-service-request-ptp-advanced-configuration-priority1 int                  precision-time-service-request-ptp-advanced-configuration-priority1
      --precision-time-service-request-ptp-advanced-configuration-priority2 int                  precision-time-service-request-ptp-advanced-configuration-priority2
      --precision-time-service-request-ptp-advanced-configuration-time-scale string              precision-time-service-request-ptp-advanced-configuration-time-scale
      --precision-time-service-request-ptp-advanced-configuration-transport-mode string          precision-time-service-request-ptp-advanced-configuration-transport-mode
      --precision-time-service-request-type string                                               precision-time-service-request-type (required)
      --request string                                                                           JSON payload for additional optional fields not exposed as flags
      --service-id string                                                                        Service UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 precision-time](equinix_fabricv4_precision-time.md)	 - Manage precision-time resources

