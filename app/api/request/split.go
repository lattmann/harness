package request

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	SplitUUIDParam = "id"
)

func GetUUIDParam(r *http.Request) (uuid.UUID, error) {
	id, err := PathParamOrError(r, SplitUUIDParam)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid uuid for %s: %w", SplitUUIDParam, err)
	}

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid uuid for %s: %w", SplitUUIDParam, err)
	}

	return parsedUUID, nil
}
