## equinix fabricv4 deployments create-topology-deployment

Create a new topology deployment

### Synopsis

The deployment API is used to creates new deployment topologies. It allows users to define the properties of the deployment, including Fabric and CSP providers.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 deployments create-topology-deployment [flags]
```

### Options

```
      --deployment-account-account-name string               Account name
      --deployment-account-account-number int                Account number
      --deployment-account-additional-properties string      deployment-account-additional-properties (JSON)
      --deployment-account-global-cust-id string             Account name
      --deployment-account-global-org-id string              Global organization identifier
      --deployment-account-global-organization-name string   Global organization name
      --deployment-account-org-id int                        Customer organization identifier
      --deployment-account-organization-name string          Customer organization name
      --deployment-account-reseller-account-name string      Reseller account name
      --deployment-account-reseller-account-number int       Reseller account number
      --deployment-account-reseller-org-id int               Reseller customer organization identifier
      --deployment-account-reseller-ucm-id string            Reseller account ucmId
      --deployment-account-ucm-id string                     Account ucmId
      --deployment-additional-properties string              deployment-additional-properties (JSON)
      --deployment-description string                        deployment-description
      --deployment-name string                               deployment-name
      --deployment-notifications string                      Preferences for notifications on deployement status changes (JSON array)
      --deployment-order-additional-properties string        deployment-order-additional-properties (JSON)
      --deployment-order-billing-tier string                 Billing tier for connection bandwidth
      --deployment-order-customer-reference-number string    Customer reference number
      --deployment-order-order-id string                     Order Identification
      --deployment-order-order-number string                 Order Reference Number
      --deployment-order-purchase-order-number string        Purchase order number
      --deployment-order-term-length int                     Term length in months, valid values are 1, 12, 24, 36 where 1 is the default value (for on-demand case).
      --deployment-project-additional-properties string      deployment-project-additional-properties (JSON)
      --deployment-project-project-id string                 Subscriber-assigned project ID
      --deployment-providers string                          deployment-providers (JSON array)
      --deployment-topology-additional-properties string     deployment-topology-additional-properties (JSON)
      --deployment-topology-uuid string                      deployment-topology-uuid
      --deployment-type string                               deployment-type
  -h, --help                                                 help for create-topology-deployment
      --request string                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 deployments](equinix_fabricv4_deployments.md)	 - Manage deployments resources

