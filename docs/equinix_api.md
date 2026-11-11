## equinix api

Make a raw API request to the given Equinix API path

### Synopsis

Make a raw authenticated request to the Equinix API using your configured credentials.

```
equinix api [url-path] [flags]
```

### Options

```
  -d, --data string           Data to send with POST/PUT requests
      --debug[=true|false]    Enable debug logging for HTTP requests and responses
  -f, --format string         Format to use for output (json or yaml) (default "json")
  -h, --help[=true|false]     help for api
  -X, --method string         HTTP method to use (default "GET")
      --portal[=true|false]   Use Equinix Portal API (cookie auth)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
```

### SEE ALSO

* [equinix](equinix.md)	 - Equinix CLI

