func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    chars := strings.Split(s, "")

	for i := 0; i < len(indices); i++ {
		index := indices[i]
		source := sources[i]
		target := targets[i]

		if len(source) <= len(s[index:]) && s[index:index+len(source)] == source {
			chars[index] = target

			for j := index + 1; j < index + len(source); j++ {
				chars[j] = ""
			}
		}
	}

	return strings.Join(chars, "")
}