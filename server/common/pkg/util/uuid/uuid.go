package uuid

import (
	"strings"

	"github.com/google/uuid"
)

func CreateUuid() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	uu := u.String()
	uu = strings.ReplaceAll(uu, "-", "")
	return uu
}
