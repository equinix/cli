// Package naming derives user-facing CLI command names from the Equinix API
// service names used by the SDK.
//
// The published API specs carry a version suffix in their names (fabricv4,
// ordersv2, stsv1alpha). That suffix is a trait of the specs rather than
// something a CLI user should have to type, so the command name is the
// canonical service name and the API name is kept as an alias.
package naming

import "regexp"

// apiVersionSuffix matches the trailing version segment of an API service
// name: the "v4" in "fabricv4", or the "v1alpha" in "stsv1alpha".
var apiVersionSuffix = regexp.MustCompile(`^(.+?)v([0-9]+)((?:alpha|beta)[0-9]*)?$`)

// CanonicalServiceName strips the version suffix from an API service name,
// e.g. "fabricv4" -> "fabric" and "ordersv2" -> "orders".
//
// Pre-release APIs are the exception: their stability level is not an
// implementation detail a user should have to look up, so it is kept as a
// trailing word, e.g. "stsv1alpha" -> "sts-alpha". Names without a version
// suffix are returned unchanged.
func CanonicalServiceName(apiName string) string {
	match := apiVersionSuffix.FindStringSubmatch(apiName)
	if match == nil {
		return apiName
	}

	base, stability := match[1], match[3]
	if stability == "" {
		return base
	}

	return base + "-" + stability
}

// ServiceAliases returns the alternate names a service command should answer
// to, so that the versioned API name keeps working after the command was
// renamed to its canonical form. It returns nil when the API name is already
// canonical, since cobra treats an alias equal to the command name as a
// duplicate.
func ServiceAliases(apiName string) []string {
	canonical := CanonicalServiceName(apiName)
	if canonical == apiName {
		return nil
	}

	return []string{apiName}
}
