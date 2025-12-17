## equinix metalv1 devices update-device

Update the device

### Synopsis

Updates the device.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices update-device [flags]
```

### Options

```
      --device-update-input-additional-properties string   device-update-input-additional-properties (JSON)
      --device-update-input-always_pxe                     device-update-input-always_pxe
      --device-update-input-billing_cycle string           device-update-input-billing_cycle
      --device-update-input-customdata string              device-update-input-customdata (JSON)
      --device-update-input-description string             device-update-input-description
      --device-update-input-firmware_set_id string         device-update-input-firmware_set_id
      --device-update-input-hostname string                device-update-input-hostname
      --device-update-input-ipxe_script_url string         device-update-input-ipxe_script_url
      --device-update-input-locked                         Whether the device should be locked, preventing accidental deletion.
      --device-update-input-network_frozen                 If true, this instance can not be converted to a different network type.
      --device-update-input-spot_instance                  Can be set to false to convert a spot-market instance to on-demand.
      --device-update-input-tags string                    device-update-input-tags (JSON array)
      --device-update-input-userdata string                device-update-input-userdata
      --exclude string                                     exclude field (JSON or string)
  -h, --help                                               help for update-device
      --id string                                          Device UUID (required)
      --include string                                     include field (JSON or string)
      --request string                                     JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

