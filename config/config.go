package config

import "strings"

// Join combines the tokens into an env path.
func Join(tokens ...string) string {
	all := []string{}
	for _, tok := range tokens {
		all = append(all, strings.Split(strings.ToLower(tok), "_")...)
	}
	return strings.Join(all, "_")
}
