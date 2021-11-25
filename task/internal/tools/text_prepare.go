package tools

func GetPureTitle(text string) string {
	if len(text) == 0 {
		return text
	}
	// TODO `[jira-id] - Title` => Title
	// вернуть просто Title
	return text
}

func GetPureTaskKey(text string) string {
	if len(text) == 0 {
		return text
	}
	// TODO `[jira-id] - Title` => jira-id
	// вернуть просто jira-id или пустую строку
	return text
}
