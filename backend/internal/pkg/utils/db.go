package utils

import "strings"

var (
	likeQueryEscapeTargetList = []string{"\\", "%", "_"}
	likeQueryEscapeList       = []string{"\\\\", "\\%", "\\_"}
)

func EscapeLikeQuery(query string) string {
	result := strings.TrimSpace(query)

	for i := range likeQueryEscapeTargetList {
		result = strings.ReplaceAll(result, likeQueryEscapeTargetList[i], likeQueryEscapeList[i])
	}

	return result
}
