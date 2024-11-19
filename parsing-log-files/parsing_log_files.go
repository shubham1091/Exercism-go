package parsinglogfiles

import (
	"regexp"
	"strings"
)

// 1. Identify garbled log lines
func IsValidLine(text string) bool {
	validPrefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

// 2. Split the log line
func SplitLogLine(text string) []string {
	separatorPattern := `<[~*=\-]*>`
	re := regexp.MustCompile(separatorPattern)
	return re.Split(text, -1)
}

// 3. Count the number of lines containing password in quoted text
func CountQuotedPasswords(lines []string) int {
	passwordPattern := `"[^"]*password[^"]*"`
	re := regexp.MustCompile(passwordPattern)
	count := 0
	for _, line := range lines {
		if re.MatchString(strings.ToLower(line)) {
			count++
		}
	}
	return count
}

// 4. Remove artifacts from log
func RemoveEndOfLineText(text string) string {
	endOfLinePattern := `end-of-line\d+`
	re := regexp.MustCompile(endOfLinePattern)
	return re.ReplaceAllString(text, "")
}

// 5. Tag lines with user names
func TagWithUserName(lines []string) []string {
	userPattern := `User\s+(\S+)`
	re := regexp.MustCompile(userPattern)
	var result []string
	for _, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			userName := match[1]
			result = append(result, "[USR] "+userName+" "+line)
		} else {
			result = append(result, line)
		}
	}
	return result
}
