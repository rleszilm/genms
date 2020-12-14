package generator

func strRef(s string) *string {
	return &s
}

func cloneValues(m map[string]interface{}) map[string]interface{} {
	n := map[string]interface{}{}

	for k, v := range m {
		n[k] = v
	}
	return n
}
