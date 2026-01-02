package phonenumber

func IsValid(phone string) bool {
	// TODO - technical debt - using regex to validate
	if len(phone) != 11 {
		return false
	}
	if phone[0:2] != "09" {
		return false
	}

	return true
}
