## equinix metalv1 organizations create-organization

Create an organization

### Synopsis

Creates an organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations create-organization [flags]
```

### Options

```
      --exclude string                                                    exclude field (JSON or string)
  -h, --help                                                              help for create-organization
      --include string                                                    include field (JSON or string)
      --organization-input-additional-properties string                   organization-input-additional-properties (JSON)
      --organization-input-address-additional-properties string           organization-input-address-additional-properties (JSON)
      --organization-input-address-address string                         organization-input-address-address
      --organization-input-address-address2 string                        organization-input-address-address2
      --organization-input-address-city string                            organization-input-address-city
      --organization-input-address-coordinates string                     organization-input-address-coordinates (JSON)
      --organization-input-address-country string                         organization-input-address-country
      --organization-input-address-state string                           organization-input-address-state
      --organization-input-address-zip_code string                        organization-input-address-zip_code
      --organization-input-billing_address-additional-properties string   organization-input-billing_address-additional-properties (JSON)
      --organization-input-billing_address-address string                 organization-input-billing_address-address
      --organization-input-billing_address-address2 string                organization-input-billing_address-address2
      --organization-input-billing_address-city string                    organization-input-billing_address-city
      --organization-input-billing_address-coordinates string             organization-input-billing_address-coordinates (JSON)
      --organization-input-billing_address-country string                 organization-input-billing_address-country
      --organization-input-billing_address-state string                   organization-input-billing_address-state
      --organization-input-billing_address-zip_code string                organization-input-billing_address-zip_code
      --organization-input-customdata string                              organization-input-customdata (JSON)
      --organization-input-description string                             organization-input-description
      --organization-input-name string                                    organization-input-name
      --organization-input-twitter string                                 organization-input-twitter
      --organization-input-website string                                 organization-input-website
      --request string                                                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

