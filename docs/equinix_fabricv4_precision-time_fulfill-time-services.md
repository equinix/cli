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
      --precision-time-service-request-additional-properties string                              precision-time-service-request-additional-properties (JSON)
      --precision-time-service-request-connections string                                        precision-time-service-request-connections (JSON array)
      --precision-time-service-request-ipv4-additional-properties string                         precision-time-service-request-ipv4-additional-properties (JSON)
      --precision-time-service-request-ipv4-default-gateway string                               Gateway Interface IP address
      --precision-time-service-request-ipv4-network-mask string                                  Network Mask
      --precision-time-service-request-ipv4-primary string                                       Primary Timing Server IP Address
      --precision-time-service-request-ipv4-secondary string                                     Secondary Timing Server IP Address
      --precision-time-service-request-name string                                               Precision Time Service name.
      --precision-time-service-request-ntp-advanced-configuration string                         NTP Advanced configuration - MD5 Authentication. (JSON array)
      --precision-time-service-request-order-additional-properties string                        precision-time-service-request-order-additional-properties (JSON)
      --precision-time-service-request-order-customer-reference-number string                    Customer reference number
      --precision-time-service-request-order-order-number string                                 Order Reference Number
      --precision-time-service-request-order-purchase-order-number string                        Purchase order number
      --precision-time-service-request-package-additional-properties string                      precision-time-service-request-package-additional-properties (JSON)
      --precision-time-service-request-package-code string                                       precision-time-service-request-package-code
      --precision-time-service-request-project-additional-properties string                      precision-time-service-request-project-additional-properties (JSON)
      --precision-time-service-request-project-project-id string                                 Subscriber-assigned project ID
      --precision-time-service-request-ptp-advanced-configuration-additional-properties string   precision-time-service-request-ptp-advanced-configuration-additional-properties (JSON)
      --precision-time-service-request-ptp-advanced-configuration-domain int                     The PTP domain value.
      --precision-time-service-request-ptp-advanced-configuration-grant-time int                 Unicast Grant Time in seconds. For Multicast and Hybrid transport modes, grant time defaults to 300 seconds. For Unicast mode, grant time can be between 30 to 7200.
      --precision-time-service-request-ptp-advanced-configuration-log-announce-interval int      precision-time-service-request-ptp-advanced-configuration-log-announce-interval
      --precision-time-service-request-ptp-advanced-configuration-log-delay-req-interval int     precision-time-service-request-ptp-advanced-configuration-log-delay-req-interval
      --precision-time-service-request-ptp-advanced-configuration-log-sync-interval int          precision-time-service-request-ptp-advanced-configuration-log-sync-interval
      --precision-time-service-request-ptp-advanced-configuration-priority1 int                  The priority1 value determines the best primary clock, Lower value indicates higher priority.
      --precision-time-service-request-ptp-advanced-configuration-priority2 int                  The priority2 value differentiates and prioritizes the primary clock to avoid confusion when priority1-value is the same for different primary clocks in a network.
      --precision-time-service-request-ptp-advanced-configuration-time-scale string              precision-time-service-request-ptp-advanced-configuration-time-scale
      --precision-time-service-request-ptp-advanced-configuration-transport-mode string          precision-time-service-request-ptp-advanced-configuration-transport-mode
      --precision-time-service-request-type string                                               precision-time-service-request-type
      --request string                                                                           JSON payload for additional optional fields not exposed as flags
      --service-id string                                                                        Service UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 precision-time](equinix_fabricv4_precision-time.md)	 - Manage precision-time resources

