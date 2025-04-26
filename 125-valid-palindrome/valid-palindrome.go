func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1

	for left < right {
		if isNonAlphaNumeric(s[left]) {
			left += 1
			continue
		}
		if isNonAlphaNumeric(s[right]) {
			right -= 1
			continue
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left += 1
		right -= 1
	}

	return true
}

func isNonAlphaNumeric(character byte) bool {
	return !((character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || (character >= '0' && character <= '9'))
}