package casing

import "testing"

func TestCamelToKebab(t *testing.T) {
	cases := map[string]string{
		"BGP":                 "bgp",
		"BGPAction":           "bgp-action",
		"BGPSession":          "bgp-session",
		"ByID":                "by-id",
		"CreateAPIKey":        "create-api-key",
		"CreateIPAssignment":  "create-ip-assignment",
		"CreateSSHKey":        "create-ssh-key",
		"DeviceSSHKeys":       "device-ssh-keys",
		"EIAService":          "eia-service",
		"EvplVCCount":         "evpl-vc-count",
		"FGVCCount":           "fg-vc-count",
		"FindAPIKeys":         "find-api-keys",
		"FindIPAddress":       "find-ip-address",
		"GETOrder":            "get-order",
		"GETRetrieve":         "get-retrieve",
		"GetOAuth":            "get-oauth",
		"RefreshOAuth":        "refresh-oauth",
		"IPAddresses":         "ip-addresses",
		"POSTOrders":          "post-orders",
		"PortVCVlan":          "port-vc-vlan",
		"ProjectAPIKey":       "project-api-key",
		"ProtocolBGPType":     "protocol-bgp-type",
		"RequestASide":        "request-a-side",
		"RequestZSide":        "request-z-side",
		"ConnectionASide":     "connection-a-side",
		"DataZSide":           "data-z-side",
		"SSHKey":              "ssh-key",
		"SSHKeys":             "ssh-keys",
		"UpdateIPAddress":     "update-ip-address",
		"UserAPIKey":          "user-api-key",
		"VLANs":               "vlans",
		"VRFs":                "vrfs",
		"VlanCSPConnection":   "vlan-csp-connection",
		"VrfBGPNeighbors":     "vrf-bgp-neighbors",
		"AccessVCCount":       "access-vc-count",
		"GetConnectionByUuid": "get-connection-by-uuid",
		"ServiceProfilesApi":  "service-profiles-api",
		"ConnectionsApi":      "connections-api",
		"OAuth2TokenApi":      "oauth2-token-api",
		"GetOAuth2Token":      "get-oauth2-token",
		"IPv4Address":         "ipv4-address",
		"IPv6Address":         "ipv6-address",
	}

	for input, want := range cases {
		if got := CamelToKebab(input); got != want {
			t.Errorf("CamelToKebab(%q) = %q, want %q", input, got, want)
		}
	}
}
