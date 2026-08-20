// Package casing provides identifier case conversion helpers shared by the
// SDK parser and command registration code.
package casing

import "strings"

// irregularCasing maps identifier substrings with non-standard capitalization
// to a form that CamelToKebab can split correctly.
var irregularCasing = map[string]string{
	"OAuth": "Oauth",
}

// acronyms are known multi-letter initialisms that should be kept together as
// a single word instead of being split letter-by-letter, ordered so that
// longer acronyms are matched before any of their shorter prefixes.
var acronyms = []string{
	"HTTPS", "OPTIONS", "DELETE", "PATCH", "HTTP", "UUID", "VLAN", "EVPL",
	"VRF", "BGP", "CSP", "EIA", "API", "SSH", "URL", "POST", "GET", "PUT",
	"VC", "IP", "ID", "FG",
}

func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }
func isLower(r rune) bool { return r >= 'a' && r <= 'z' }

// CamelToKebab converts a CamelCase or camelCase identifier to kebab-case,
// keeping known initialisms (e.g. "BGP", "SSH", "IP") together as a single
// word instead of splitting them letter by letter (e.g. "BGPSession"
// becomes "bgp-session", not "b-g-p-session").
func CamelToKebab(s string) string {
	for from, to := range irregularCasing {
		s = strings.ReplaceAll(s, from, to)
	}

	runes := []rune(s)
	n := len(runes)
	var words []string

	i := 0
	for i < n {
		if !isUpper(runes[i]) {
			j := i
			for j < n && !isUpper(runes[j]) {
				j++
			}
			words = append(words, string(runes[i:j]))
			i = j
			continue
		}

		j := i
		for j < n && isUpper(runes[j]) {
			j++
		}

		if j-i == 1 {
			// A single leading capital starts an ordinary word; absorb any
			// lowercase letters that follow it (e.g. "Get" or "Aside").
			k := j
			for k < n && isLower(runes[k]) {
				k++
			}
			words = append(words, string(runes[i:k]))
			i = k
			continue
		}

		run := string(runes[i:j])
		upperRun := strings.ToUpper(run)
		matchLen := 0
		for _, a := range acronyms {
			if len(a) <= len(upperRun) && strings.HasPrefix(upperRun, a) {
				matchLen = len(a)
				break
			}
		}

		if matchLen > 0 {
			word := run[:matchLen]
			pos := i + matchLen
			if pos == j {
				// The acronym consumed the whole capital run; any lowercase
				// letters immediately after it are a suffix (e.g. the "s"
				// in "VLANs"), not the start of a new word.
				k := pos
				for k < n && isLower(runes[k]) {
					k++
				}
				word += string(runes[pos:k])
				pos = k
			}
			words = append(words, word)
			i = pos
			continue
		}

		// Unrecognized run of capitals: the last capital starts the next
		// word if one follows (e.g. "ASide" -> "A", "Side").
		if j < n && isLower(runes[j]) {
			words = append(words, string(runes[i:j-1]))
			i = j - 1
		} else {
			words = append(words, run)
			i = j
		}
	}

	return strings.ToLower(strings.Join(words, "-"))
}
