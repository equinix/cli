package cmd

import (
	"testing"
)

// TestNoConflictingServiceNames fails the build when two service commands
// claim the same name. Canonical service names drop the API version suffix,
// so two versions of the same API (e.g. fabricv4 and a future fabricv5) would
// otherwise silently shadow one another.
func TestNoConflictingServiceNames(t *testing.T) {
	claimedBy := map[string]string{}

	for _, cmd := range rootCmd.Commands() {
		for _, name := range append([]string{cmd.Name()}, cmd.Aliases...) {
			if owner, taken := claimedBy[name]; taken {
				t.Errorf("command name %q is claimed by both %q and %q", name, owner, cmd.Name())
				continue
			}
			claimedBy[name] = cmd.Name()
		}
	}
}

// TestServiceCommandsAreCanonical checks that each onboarded service is
// reachable under its canonical name and still under its versioned API name.
func TestServiceCommandsAreCanonical(t *testing.T) {
	for apiName, canonical := range map[string]string{
		"accesstokenv1":   "accesstoken",
		"eiav2":           "eia",
		"fabricv4":        "fabric",
		"lookupv2":        "lookup",
		"orderhistoryv1":  "orderhistory",
		"ordersv2":        "orders",
		"securecabinetv1": "securecabinet",
		"smarthandsv1":    "smarthands",
		"stsv1alpha":      "sts-alpha",
	} {
		for _, name := range []string{canonical, apiName} {
			cmd, _, err := rootCmd.Find([]string{name})
			if err != nil {
				t.Errorf("looking up %q: %v", name, err)
				continue
			}
			if cmd.Name() != canonical {
				t.Errorf("%q resolved to command %q, want %q", name, cmd.Name(), canonical)
			}
		}
	}
}
