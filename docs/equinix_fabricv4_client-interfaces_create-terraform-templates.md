## equinix fabricv4 client-interfaces create-terraform-templates

Generate Terraform Deployment Templates

### Synopsis

The Client Interfaces API is used to generate Terraform Templates based on Deployment details.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 client-interfaces create-terraform-templates [flags]
```

### Options

```
      --client-interfaces-additional-properties string   client-interfaces-additional-properties (required) (JSON)
      --client-interfaces-description string             client-interfaces-description (required)
      --client-interfaces-name string                    client-interfaces-name (required)
      --client-interfaces-type string                    client-interfaces-type (required)
      --deployment-id string                             Deployment UUID (required)
  -h, --help                                             help for create-terraform-templates
      --request string                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 client-interfaces](equinix_fabricv4_client-interfaces.md)	 - Manage client-interfaces resources

