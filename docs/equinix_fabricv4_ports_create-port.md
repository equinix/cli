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
      --port-request-account-account-name string                        Account name
      --port-request-account-account-number int                         Account number
      --port-request-account-additional-properties string               port-request-account-additional-properties (JSON)
      --port-request-account-global-cust-id string                      Account name
      --port-request-account-global-org-id string                       Global organization identifier
      --port-request-account-global-organization-name string            Global organization name
      --port-request-account-org-id int                                 Customer organization identifier
      --port-request-account-organization-name string                   Customer organization name
      --port-request-account-reseller-account-name string               Reseller account name
      --port-request-account-reseller-account-number int                Reseller account number
      --port-request-account-reseller-org-id int                        Reseller customer organization identifier
      --port-request-account-reseller-ucm-id string                     Reseller account ucmId
      --port-request-account-ucm-id string                              Account ucmId
      --port-request-additional-info string                             Port additional information (JSON array)
      --port-request-additional-properties string                       port-request-additional-properties (JSON)
      --port-request-asn int                                            Port ASN
      --port-request-available-bandwidth int                            Equinix assigned response attribute for Port available bandwidth in Mbps
      --port-request-bandwidth int                                      Equinix assigned response attribute for Port bandwidth in Mbps
                                                                        Deprecated
      --port-request-bmmr-type string                                   port-request-bmmr-type
      --port-request-change-additional-properties string                port-request-change-additional-properties (JSON)
      --port-request-change-created-date-time string                    port-request-change-created-date-time (JSON)
      --port-request-change-data string                                 port-request-change-data (JSON)
      --port-request-change-information string                          Additional information
      --port-request-change-log-additional-properties string            port-request-change-log-additional-properties (JSON)
      --port-request-change-log-created-by string                       Created by User Key
      --port-request-change-log-created-by-email string                 Created by User Email Address
      --port-request-change-log-created-by-full-name string             Created by User Full Name
      --port-request-change-log-created-date-time string                port-request-change-log-created-date-time (JSON)
      --port-request-change-log-deleted-by string                       Deleted by User Key
      --port-request-change-log-deleted-by-email string                 Deleted by User Email Address
      --port-request-change-log-deleted-by-full-name string             Deleted by User Full Name
      --port-request-change-log-deleted-date-time string                port-request-change-log-deleted-date-time (JSON)
      --port-request-change-log-updated-by string                       Updated by User Key
      --port-request-change-log-updated-by-email string                 Updated by User Email Address
      --port-request-change-log-updated-by-full-name string             Updated by User Full Name
      --port-request-change-log-updated-date-time string                port-request-change-log-updated-date-time (JSON)
      --port-request-change-status string                               port-request-change-status
      --port-request-change-type string                                 port-request-change-type
      --port-request-change-updated-date-time string                    port-request-change-updated-date-time (JSON)
      --port-request-change-uuid string                                 Uniquely identifies a change
      --port-request-connections-count int                              Equinix assigned response attribute for Connection count
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
      --port-request-description string                                 Equinix assigned response attribute for Port description
      --port-request-device-additional-properties string                port-request-device-additional-properties (JSON)
      --port-request-device-name string                                 Device name
      --port-request-device-redundancy string                           port-request-device-redundancy (JSON)
      --port-request-device-vc-bandwidth-max int                        Maximum bandwidth allowed for connection.
      --port-request-encapsulation-additional-properties string         port-request-encapsulation-additional-properties (JSON)
      --port-request-encapsulation-tag-protocol-id string               Port encapsulation tag protocol identifier
      --port-request-encapsulation-type string                          port-request-encapsulation-type
      --port-request-end-customer-additional-properties string          port-request-end-customer-additional-properties (JSON)
      --port-request-end-customer-is-disclosed                          Indicate if endCustomer info should be disclosed or not
      --port-request-end-customer-mdm-id string                         port-request-end-customer-mdm-id
      --port-request-end-customer-name string                           port-request-end-customer-name
      --port-request-href string                                        Equinix assigned response attribute for an absolute URL that is the subject of the link's context.
      --port-request-id int                                             Equinix assigned response attribute for Port Id
      --port-request-interface-additional-properties string             port-request-interface-additional-properties (JSON)
      --port-request-interface-type string                              Port interface type
      --port-request-lag-additional-properties string                   port-request-lag-additional-properties (JSON)
      --port-request-lag-enabled                                        If LAG enabled
      --port-request-lag-id string                                      id
      --port-request-lag-member-status string                           member status
      --port-request-lag-name string                                    name
      --port-request-loas string                                        Port Loas (JSON array)
      --port-request-location-additional-properties string              port-request-location-additional-properties (JSON)
      --port-request-location-ibx string                                Deprecated
      --port-request-location-metro-code string                         port-request-location-metro-code
      --port-request-location-metro-href string                         port-request-location-metro-href
      --port-request-location-metro-name string                         port-request-location-metro-name
      --port-request-location-region string                             port-request-location-region
      --port-request-name string                                        Equinix assigned response attribute for Port name
      --port-request-notifications string                               Notification preferences (JSON array)
      --port-request-operation-access-v-c-count int                     Total number of connections.
      --port-request-operation-additional-properties string             port-request-operation-additional-properties (JSON)
      --port-request-operation-connection-count int                     Total number of connections.
      --port-request-operation-evpl-v-c-count int                       Total number of connections.
      --port-request-operation-fg-v-c-count int                         Total number of connections.
      --port-request-operation-op-status-changed-at string              port-request-operation-op-status-changed-at (JSON)
      --port-request-operation-operational-status string                port-request-operation-operational-status
      --port-request-order-additional-properties string                 port-request-order-additional-properties (JSON)
      --port-request-order-customer-reference-id string                 Customer order reference Id
      --port-request-order-order-id string                              Order Identification
      --port-request-order-order-number string                          Order Reference Number
      --port-request-order-purchase-order string                        port-request-order-purchase-order (JSON)
      --port-request-order-signature string                             port-request-order-signature (JSON)
      --port-request-order-uuid string                                  Equinix-assigned order identifier, this is a derived response atrribute
                                                                        Deprecated
      --port-request-package-additional-properties string               port-request-package-additional-properties (JSON)
      --port-request-package-code string                                port-request-package-code
      --port-request-package-type string                                port-request-package-type
      --port-request-physical-port-quantity int                         Number of physical ports
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
      --port-request-service-type string                                Deprecated
      --port-request-settings-additional-properties string              port-request-settings-additional-properties (JSON)
      --port-request-settings-buyout                                    Deprecated
      --port-request-settings-layer3-enabled                            Deprecated
      --port-request-settings-package-type string                       Deprecated
      --port-request-settings-place-vc-order-permission                 Deprecated
      --port-request-settings-shared-port-product string                port-request-settings-shared-port-product
      --port-request-settings-shared-port-type                          port-request-settings-shared-port-type
      --port-request-settings-view-port-permission                      Deprecated
      --port-request-state string                                       port-request-state
      --port-request-tether-ibx string                                  z-side/Equinix ibx
      --port-request-type string                                        port-request-type
      --port-request-used-bandwidth int                                 Equinix assigned response attribute for Port used bandwidth in Mbps
      --port-request-uuid string                                        Equinix assigned response attribute for  port identifier
      --request string                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

