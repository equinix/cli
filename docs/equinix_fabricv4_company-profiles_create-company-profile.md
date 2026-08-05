## equinix fabricv4 company-profiles create-company-profile

Execute create-company-profile operation

### Synopsis

Execute the create-company-profile operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 company-profiles create-company-profile [flags]
```

### Options

```
      --company-profile-request-additional-properties string   company-profile-request-additional-properties (JSON)
      --company-profile-request-contact-url string             company-profile-request-contact-url
      --company-profile-request-description string             company-profile-request-description
      --company-profile-request-name string                    company-profile-request-name
      --company-profile-request-notifications string           company-profile-request-notifications (JSON array)
      --company-profile-request-summary string                 company-profile-request-summary
      --company-profile-request-type string                    company-profile-request-type
      --company-profile-request-web-url string                 company-profile-request-web-url
  -h, --help                                                   help for create-company-profile
      --request string                                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 company-profiles](equinix_fabricv4_company-profiles.md)	 - Manage company-profiles resources

