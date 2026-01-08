## equinix metalv1 spot-market create-spot-market-request

Create a spot market request

### Synopsis

Creates a new spot market request. Type-specific options (such as operating_system for baremetal devices) should be included in the main data structure alongside hostname and plan. The features attribute allows you to optionally specify what features your server should have. For example, if you require a server with a TPM chip, you may specify `{ "features": { "tpm": "required" } }` (or `{ "features": ["tpm"] }` in shorthand). The request will fail if there are no available servers matching your criteria. Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a preferred value (see the example request below). The request will not fail if we have no servers with that feature in our inventory.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 spot-market create-spot-market-request [flags]
```

### Options

```
  -h, --help                                                                                help for create-spot-market-request
      --id string                                                                           Project UUID (required)
      --request string                                                                      JSON payload for additional optional fields not exposed as flags
      --spot-market-request-create-input-additional-properties string                       spot-market-request-create-input-additional-properties (JSON)
      --spot-market-request-create-input-devices_max int                                    spot-market-request-create-input-devices_max
      --spot-market-request-create-input-devices_min int                                    spot-market-request-create-input-devices_min
      --spot-market-request-create-input-facilities string                                  Deprecated (JSON array)
      --spot-market-request-create-input-instance_parameters-additional-properties string   spot-market-request-create-input-instance_parameters-additional-properties (JSON)
      --spot-market-request-create-input-instance_parameters-always_pxe                     spot-market-request-create-input-instance_parameters-always_pxe
      --spot-market-request-create-input-instance_parameters-billing_cycle string           spot-market-request-create-input-instance_parameters-billing_cycle
      --spot-market-request-create-input-instance_parameters-customdata string              spot-market-request-create-input-instance_parameters-customdata (JSON)
      --spot-market-request-create-input-instance_parameters-description string             spot-market-request-create-input-instance_parameters-description
      --spot-market-request-create-input-instance_parameters-features string                spot-market-request-create-input-instance_parameters-features (JSON array)
      --spot-market-request-create-input-instance_parameters-hostname string                spot-market-request-create-input-instance_parameters-hostname
      --spot-market-request-create-input-instance_parameters-hostnames string               spot-market-request-create-input-instance_parameters-hostnames (JSON array)
      --spot-market-request-create-input-instance_parameters-locked                         Whether the device should be locked, preventing accidental deletion.
      --spot-market-request-create-input-instance_parameters-no_ssh_keys                    spot-market-request-create-input-instance_parameters-no_ssh_keys
      --spot-market-request-create-input-instance_parameters-operating_system string        spot-market-request-create-input-instance_parameters-operating_system
      --spot-market-request-create-input-instance_parameters-plan string                    spot-market-request-create-input-instance_parameters-plan
      --spot-market-request-create-input-instance_parameters-private_ipv4_subnet_size int   spot-market-request-create-input-instance_parameters-private_ipv4_subnet_size
      --spot-market-request-create-input-instance_parameters-project_ssh_keys string        spot-market-request-create-input-instance_parameters-project_ssh_keys (JSON array)
      --spot-market-request-create-input-instance_parameters-public_ipv4_subnet_size int    spot-market-request-create-input-instance_parameters-public_ipv4_subnet_size
      --spot-market-request-create-input-instance_parameters-tags string                    spot-market-request-create-input-instance_parameters-tags (JSON array)
      --spot-market-request-create-input-instance_parameters-termination_time string        spot-market-request-create-input-instance_parameters-termination_time (JSON)
      --spot-market-request-create-input-instance_parameters-user_ssh_keys string           The UUIDs of users whose SSH keys should be included on the provisioned device. (JSON array)
      --spot-market-request-create-input-instance_parameters-userdata string                spot-market-request-create-input-instance_parameters-userdata
      --spot-market-request-create-input-max_bid_price float                                spot-market-request-create-input-max_bid_price
      --spot-market-request-create-input-metro string                                       The metro ID or code the spot market request will be created in.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 spot-market](equinix_metalv1_spot-market.md)	 - Manage spot-market resources

