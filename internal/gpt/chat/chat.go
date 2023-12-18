package chat

import (
	"strings"

	"github.com/google/uuid"
)

const TempChatIdPrefix = "temp-"

func GenerateTempChatId() string {
	u := uuid.New()
	return TempChatIdPrefix + strings.ReplaceAll(u.String(), "-", "")
}
