package main

import "strings"

func numUniqueEmails(emails []string) int {
	seen := make(map[string]bool)
	for _, email := range emails{
		i := strings.Index(email, "@")
		local := email[:i]
		rest := email[i:]

		if strings.Contains(local, "+") {
			local = email[:strings.Index(email,"+")]
		}

		local = strings.Replace(local, ".","",-1)

		seen[local+rest] = true
	}
	return len(seen)
}