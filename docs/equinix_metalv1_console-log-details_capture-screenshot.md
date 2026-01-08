## equinix metalv1 console-log-details capture-screenshot

Method for CaptureScreenshot

### Synopsis

Capture a screenshot from the device, if supported, via the BMC.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 console-log-details capture-screenshot [flags]
```

### Options

```
  -h, --help             help for capture-screenshot
      --id string        Device UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 console-log-details](equinix_metalv1_console-log-details.md)	 - Manage console-log-details resources

