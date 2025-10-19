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
      --deployment-account-account-name string               deployment-account-account-name
      --deployment-account-account-number int                deployment-account-account-number
      --deployment-account-additional-properties string      deployment-account-additional-properties (required) (JSON)
      --deployment-account-global-cust-id string             deployment-account-global-cust-id
      --deployment-account-global-org-id string              deployment-account-global-org-id
      --deployment-account-global-organization-name string   deployment-account-global-organization-name
      --deployment-account-org-id int                        deployment-account-org-id
      --deployment-account-organization-name string          deployment-account-organization-name
      --deployment-account-reseller-account-name string      deployment-account-reseller-account-name
      --deployment-account-reseller-account-number int       deployment-account-reseller-account-number
      --deployment-account-reseller-org-id int               deployment-account-reseller-org-id
      --deployment-account-reseller-ucm-id string            deployment-account-reseller-ucm-id
      --deployment-account-ucm-id string                     deployment-account-ucm-id
      --deployment-additional-properties string              deployment-additional-properties (required) (JSON)
      --deployment-description string                        deployment-description
      --deployment-name string                               deployment-name (required)
      --deployment-notifications string                      deployment-notifications (JSON array)
      --deployment-order-additional-properties string        deployment-order-additional-properties (JSON)
      --deployment-order-billing-tier string                 deployment-order-billing-tier
      --deployment-order-customer-reference-number string    deployment-order-customer-reference-number
      --deployment-order-order-id string                     deployment-order-order-id
      --deployment-order-order-number string                 deployment-order-order-number
      --deployment-order-purchase-order-number string        deployment-order-purchase-order-number
      --deployment-order-term-length int                     deployment-order-term-length
      --deployment-project-additional-properties string      deployment-project-additional-properties (required) (JSON)
      --deployment-project-project-id string                 deployment-project-project-id (required)
      --deployment-providers string                          deployment-providers (required) (JSON array)
      --deployment-topology-additional-properties string     deployment-topology-additional-properties (JSON)
      --deployment-topology-uuid string                      deployment-topology-uuid
      --deployment-type string                               deployment-type (required)
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

