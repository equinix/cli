## equinix smarthandsv1 smarthands handle-smart-hand-run-jumper-cable-order

This API is used to request a jumper cable to be ran between devices.

### Synopsis

This API is used to request a jumper cable to be ran between devices.

Use --request flag to provide optional JSON payload fields.

```
equinix smarthandsv1 smarthands handle-smart-hand-run-jumper-cable-order [flags]
```

### Options

```
      --authorization string                                authorization field
      --body-additional-properties string                   body-additional-properties (JSON)
      --body-attachments string                             Use this to pass uploaded attachments. Attachments need to be uploaded using the attachments API. Maximum size of an attachment is 2MB with the following formats - bmp, jpg, jpeg, gif, png, tif, tiff, txt, doc, docx, xls, xlsx, ppt, pps, ppsx, pdf, and vsd. (JSON array)
      --body-contacts string                                Use this array to pass ordering contact, notification contacts and technical contact. Only one ordering contact, technical contact is allowed. One or more notification contacts are allowed. Ordering and notification contacts are always registered customers with the customer portal. (JSON array)
      --body-customer-reference-number string               You may use numbers and text in this field to enter reference information for your records. This will also appear in your reports and details. You may use this information to search for this content on the submitted requests page.
      --body-ibx-location-additional-properties string      body-ibx-location-additional-properties (JSON)
      --body-ibx-location-cages string                      body-ibx-location-cages (JSON array)
      --body-ibx-location-ibx string                        body-ibx-location-ibx
      --body-purchase-order-additional-properties string    body-purchase-order-additional-properties (JSON)
      --body-purchase-order-number string                   Purchase Order Number
      --body-purchase-order-purchase-order-type string      body-purchase-order-purchase-order-type
      --body-schedule-additional-properties string          body-schedule-additional-properties (JSON)
      --body-schedule-requested-completion-date string      body-schedule-requested-completion-date (JSON)
      --body-schedule-requested-start-date string           body-schedule-requested-start-date (JSON)
      --body-schedule-schedule-type string                  body-schedule-schedule-type
      --body-service-details-additional-properties string   body-service-details-additional-properties (JSON)
      --body-service-details-cable-id string                Cable ID
      --body-service-details-connector string               body-service-details-connector
      --body-service-details-device-details string          body-service-details-device-details (JSON array)
      --body-service-details-jumper-type string             body-service-details-jumper-type
      --body-service-details-media-type string              body-service-details-media-type
      --body-service-details-provide-tx-rx-light-levels     Provide Tx/Rx Light Levels
      --body-service-details-quantity string                body-service-details-quantity
      --body-service-details-scope-of-work string           Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field.
  -h, --help                                                help for handle-smart-hand-run-jumper-cable-order
      --request string                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix smarthandsv1 smarthands](equinix_smarthandsv1_smarthands.md)	 - Manage smarthands resources

