package naming

import (
	"reflect"
	"testing"
)

func TestCanonicalServiceName(t *testing.T) {
	tests := map[string]string{
		// Every service currently onboarded from the SDK.
		"accesstokenv1":   "accesstoken",
		"eiav2":           "eia",
		"fabricv4":        "fabric",
		"lookupv2":        "lookup",
		"orderhistoryv1":  "orderhistory",
		"ordersv2":        "orders",
		"securecabinetv1": "securecabinet",
		"smarthandsv1":    "smarthands",
		"stsv1alpha":      "sts-alpha",

		// Other version shapes the specs use.
		"metalv1":         "metal",
		"examplev10":      "example",
		"examplev2beta":   "example-beta",
		"examplev1alpha2": "example-alpha2",

		// A name whose base ends in "v" must keep it.
		"vrfv2": "vrf",

		// No version suffix to strip.
		"lookup": "lookup",
		"v1":     "v1",
		"":       "",
	}

	for apiName, want := range tests {
		if got := CanonicalServiceName(apiName); got != want {
			t.Errorf("CanonicalServiceName(%q) = %q, want %q", apiName, got, want)
		}
	}
}

func TestServiceAliases(t *testing.T) {
	if got, want := ServiceAliases("fabricv4"), []string{"fabricv4"}; !reflect.DeepEqual(got, want) {
		t.Errorf("ServiceAliases(\"fabricv4\") = %q, want %q", got, want)
	}

	if got := ServiceAliases("lookup"); got != nil {
		t.Errorf("ServiceAliases(\"lookup\") = %q, want nil", got)
	}
}
