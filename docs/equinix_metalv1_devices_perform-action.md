## equinix metalv1 devices perform-action

Perform an action

### Synopsis

Performs an action for the given device.  Possible actions include: power_on, power_off, reboot, reinstall, and rescue (reboot the device into rescue OS.)

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices perform-action [flags]
```

### Options

```
      --device-action-input-additional-properties string   device-action-input-additional-properties (JSON)
      --device-action-input-deprovision_fast reinstall     When type is reinstall, enabling fast deprovisioning will bypass full disk wiping.
      --device-action-input-force_delete                   May be required to perform actions under certain conditions
      --device-action-input-ipxe_script_url reinstall      When type is reinstall, use this `ipxe_script_url` (`operating_system` must be `custom_ipxe`, defaults to the current `ipxe_script_url`)
      --device-action-input-operating_system reinstall     When type is reinstall, use this `operating_system` (defaults to the current `operating system`)
      --device-action-input-preserve_data reinstall        When type is reinstall, preserve the existing data on all disks except the operating-system disk.
      --device-action-input-type string                    device-action-input-type
  -h, --help                                               help for perform-action
      --id string                                          Device UUID (required)
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

