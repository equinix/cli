## equinix fabricv4 networks create-network

Create Network

### Synopsis

This API provides capability to create user's Fabric Network

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 networks create-network [flags]
```

### Options

```
      --dry-run                                                      dry-run field (required)
  -h, --help                                                         help for create-network
      --network-post-request-additional-properties string            network-post-request-additional-properties (required) (JSON)
      --network-post-request-location-additional-properties string   network-post-request-location-additional-properties (JSON)
      --network-post-request-location-ibx string                     network-post-request-location-ibx
      --network-post-request-location-metro-code string              network-post-request-location-metro-code
      --network-post-request-location-metro-href string              network-post-request-location-metro-href
      --network-post-request-location-metro-name string              network-post-request-location-metro-name
      --network-post-request-location-region string                  network-post-request-location-region
      --network-post-request-name string                             network-post-request-name (required)
      --network-post-request-notifications string                    network-post-request-notifications (required) (JSON array)
      --network-post-request-project-additional-properties string    network-post-request-project-additional-properties (JSON)
      --network-post-request-project-project-id string               network-post-request-project-project-id
      --network-post-request-scope string                            network-post-request-scope (required)
      --network-post-request-type string                             network-post-request-type (required)
      --request string                                               JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 networks](equinix_fabricv4_networks.md)	 - Manage networks resources

