package platform

import "strings"

// findKeyWhichHasLongestMatchValue returns key in map which has longest matched value with target.
// If no value was matched with target, this returns defaultKey.
func findKeyWhichHasLongestMatchValue[E ~string](m map[E][]string, target string) E {
	var (
		matchKey   E
		matchValue string
	)
	for key, values := range m {
		for _, value := range values {
			if strings.Contains(target, value) && len(matchValue) < len(value) {
				matchKey = key
				matchValue = value
			}
		}
	}
	return matchKey
}
