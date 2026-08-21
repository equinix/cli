## equinix metalv1 ip-addresses delete-ip-address

Unassign an ip address

### Synopsis

This call can be used to un-assign an IP assignment or delete an IP reservation. Un-assign an IP address record. Use the assignment UUID you get after attaching the IP. This will remove the relationship between an IP and the device or metal gateway and will make the IP address available to be assigned to another device, once the IP has been un-configured from the network. Delete an IP reservation. Use the reservation UUID you get after adding the IP to the project. This will permanently delete the IP block reservation from the project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ip-addresses delete-ip-address [flags]
```

### Options

```
  -h, --help             help for delete-ip-address
      --id string        IP Address UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ip-addresses](equinix_metalv1_ip-addresses.md)	 - Manage ip-addresses resources

