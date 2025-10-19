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
      --dry-run                                                         dry-run field (required)
  -h, --help                                                            help for create-port
      --port-request-account-account-name string                        port-request-account-account-name
      --port-request-account-account-number int                         port-request-account-account-number
      --port-request-account-additional-properties string               port-request-account-additional-properties (required) (JSON)
      --port-request-account-global-cust-id string                      port-request-account-global-cust-id
      --port-request-account-global-org-id string                       port-request-account-global-org-id
      --port-request-account-global-organization-name string            port-request-account-global-organization-name
      --port-request-account-org-id int                                 port-request-account-org-id
      --port-request-account-organization-name string                   port-request-account-organization-name
      --port-request-account-reseller-account-name string               port-request-account-reseller-account-name
      --port-request-account-reseller-account-number int                port-request-account-reseller-account-number
      --port-request-account-reseller-org-id int                        port-request-account-reseller-org-id
      --port-request-account-reseller-ucm-id string                     port-request-account-reseller-ucm-id
      --port-request-account-ucm-id string                              port-request-account-ucm-id
      --port-request-additional-info string                             port-request-additional-info (JSON array)
      --port-request-additional-properties string                       port-request-additional-properties (required) (JSON)
      --port-request-asn int                                            port-request-asn
      --port-request-available-bandwidth int                            port-request-available-bandwidth
      --port-request-bandwidth int                                      port-request-bandwidth
      --port-request-bmmr-type string                                   port-request-bmmr-type
      --port-request-change-additional-properties string                port-request-change-additional-properties (JSON)
      --port-request-change-created-date-time string                    port-request-change-created-date-time (JSON)
      --port-request-change-data string                                 port-request-change-data (JSON)
      --port-request-change-information string                          port-request-change-information
      --port-request-change-log-additional-properties string            port-request-change-log-additional-properties (JSON)
      --port-request-change-log-created-by string                       port-request-change-log-created-by
      --port-request-change-log-created-by-email string                 port-request-change-log-created-by-email
      --port-request-change-log-created-by-full-name string             port-request-change-log-created-by-full-name
      --port-request-change-log-created-date-time string                port-request-change-log-created-date-time (JSON)
      --port-request-change-log-deleted-by string                       port-request-change-log-deleted-by
      --port-request-change-log-deleted-by-email string                 port-request-change-log-deleted-by-email
      --port-request-change-log-deleted-by-full-name string             port-request-change-log-deleted-by-full-name
      --port-request-change-log-deleted-date-time string                port-request-change-log-deleted-date-time (JSON)
      --port-request-change-log-updated-by string                       port-request-change-log-updated-by
      --port-request-change-log-updated-by-email string                 port-request-change-log-updated-by-email
      --port-request-change-log-updated-by-full-name string             port-request-change-log-updated-by-full-name
      --port-request-change-log-updated-date-time string                port-request-change-log-updated-date-time (JSON)
      --port-request-change-status string                               port-request-change-status
      --port-request-change-type string                                 port-request-change-type
      --port-request-change-updated-date-time string                    port-request-change-updated-date-time (JSON)
      --port-request-change-uuid string                                 port-request-change-uuid
      --port-request-connections-count int                              port-request-connections-count
      --port-request-connectivity-source-type string                    port-request-connectivity-source-type (required)
      --port-request-demarcation-point-additional-properties string     port-request-demarcation-point-additional-properties (JSON)
      --port-request-demarcation-point-cabinet-unique-space-id string   port-request-demarcation-point-cabinet-unique-space-id
      --port-request-demarcation-point-cage-unique-space-id string      port-request-demarcation-point-cage-unique-space-id
      --port-request-demarcation-point-connector-type string            port-request-demarcation-point-connector-type
      --port-request-demarcation-point-ibx string                       port-request-demarcation-point-ibx
      --port-request-demarcation-point-patch-panel string               port-request-demarcation-point-patch-panel
      --port-request-demarcation-point-patch-panel-name string          port-request-demarcation-point-patch-panel-name
      --port-request-demarcation-point-patch-panel-port-a string        port-request-demarcation-point-patch-panel-port-a
      --port-request-demarcation-point-patch-panel-port-b string        port-request-demarcation-point-patch-panel-port-b
      --port-request-description string                                 port-request-description
      --port-request-device-additional-properties string                port-request-device-additional-properties (JSON)
      --port-request-device-name string                                 port-request-device-name
      --port-request-device-redundancy string                           port-request-device-redundancy (JSON)
      --port-request-device-vc-bandwidth-max int                        port-request-device-vc-bandwidth-max
      --port-request-encapsulation-additional-properties string         port-request-encapsulation-additional-properties (required) (JSON)
      --port-request-encapsulation-tag-protocol-id string               port-request-encapsulation-tag-protocol-id
      --port-request-encapsulation-type string                          port-request-encapsulation-type
      --port-request-end-customer-additional-properties string          port-request-end-customer-additional-properties (JSON)
      --port-request-end-customer-is-disclosed                          port-request-end-customer-is-disclosed
      --port-request-end-customer-mdm-id string                         port-request-end-customer-mdm-id
      --port-request-end-customer-name string                           port-request-end-customer-name
      --port-request-href string                                        port-request-href
      --port-request-id int                                             port-request-id
      --port-request-interface-additional-properties string             port-request-interface-additional-properties (JSON)
      --port-request-interface-type string                              port-request-interface-type
      --port-request-lag-additional-properties string                   port-request-lag-additional-properties (JSON)
      --port-request-lag-enabled                                        port-request-lag-enabled
      --port-request-lag-id string                                      port-request-lag-id
      --port-request-lag-member-status string                           port-request-lag-member-status
      --port-request-lag-name string                                    port-request-lag-name
      --port-request-loas string                                        port-request-loas (JSON array)
      --port-request-location-additional-properties string              port-request-location-additional-properties (required) (JSON)
      --port-request-location-ibx string                                port-request-location-ibx
      --port-request-location-metro-code string                         port-request-location-metro-code
      --port-request-location-metro-href string                         port-request-location-metro-href
      --port-request-location-metro-name string                         port-request-location-metro-name
      --port-request-location-region string                             port-request-location-region
      --port-request-name string                                        port-request-name
      --port-request-notifications string                               port-request-notifications (JSON array)
      --port-request-operation-access-v-c-count int                     port-request-operation-access-v-c-count
      --port-request-operation-additional-properties string             port-request-operation-additional-properties (JSON)
      --port-request-operation-connection-count int                     port-request-operation-connection-count
      --port-request-operation-evpl-v-c-count int                       port-request-operation-evpl-v-c-count
      --port-request-operation-fg-v-c-count int                         port-request-operation-fg-v-c-count
      --port-request-operation-op-status-changed-at string              port-request-operation-op-status-changed-at (JSON)
      --port-request-operation-operational-status string                port-request-operation-operational-status
      --port-request-order-additional-properties string                 port-request-order-additional-properties (JSON)
      --port-request-order-customer-reference-id string                 port-request-order-customer-reference-id
      --port-request-order-order-id string                              port-request-order-order-id
      --port-request-order-order-number string                          port-request-order-order-number
      --port-request-order-purchase-order string                        port-request-order-purchase-order (JSON)
      --port-request-order-signature string                             port-request-order-signature (JSON)
      --port-request-order-uuid string                                  port-request-order-uuid
      --port-request-package-additional-properties string               port-request-package-additional-properties (JSON)
      --port-request-package-code string                                port-request-package-code
      --port-request-package-type string                                port-request-package-type
      --port-request-physical-port-quantity int                         port-request-physical-port-quantity
      --port-request-physical-ports string                              port-request-physical-ports (JSON array)
      --port-request-physical-ports-count int                           port-request-physical-ports-count
      --port-request-physical-ports-speed int                           port-request-physical-ports-speed (required)
      --port-request-physical-ports-type string                         port-request-physical-ports-type (required)
      --port-request-project-additional-properties string               port-request-project-additional-properties (JSON)
      --port-request-project-project-id string                          port-request-project-project-id
      --port-request-redundancy-additional-properties string            port-request-redundancy-additional-properties (JSON)
      --port-request-redundancy-enabled                                 port-request-redundancy-enabled
      --port-request-redundancy-group string                            port-request-redundancy-group
      --port-request-redundancy-priority string                         port-request-redundancy-priority
      --port-request-service-type string                                port-request-service-type
      --port-request-settings-additional-properties string              port-request-settings-additional-properties (required) (JSON)
      --port-request-settings-buyout                                    port-request-settings-buyout
      --port-request-settings-layer3-enabled                            port-request-settings-layer3-enabled
      --port-request-settings-package-type string                       port-request-settings-package-type
      --port-request-settings-place-vc-order-permission                 port-request-settings-place-vc-order-permission
      --port-request-settings-shared-port-product string                port-request-settings-shared-port-product
      --port-request-settings-shared-port-type                          port-request-settings-shared-port-type
      --port-request-settings-view-port-permission                      port-request-settings-view-port-permission
      --port-request-state string                                       port-request-state
      --port-request-tether-ibx string                                  port-request-tether-ibx
      --port-request-type string                                        port-request-type (required)
      --port-request-used-bandwidth int                                 port-request-used-bandwidth
      --port-request-uuid string                                        port-request-uuid
      --request string                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

