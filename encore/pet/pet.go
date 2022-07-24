package pet

import (
	"fmt"
	"strconv"
)

type ID int64
type CategoryID int64

type Status string

const UnsetID = 0

type Pets []Pet

type Pet struct {
	ID       ID       `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Status   Status   `json:"status,omitempty"`
}

type Category struct {
	ID   CategoryID `json:"id"`
	Name string     `json:"name"`
}

func IDFromString(v string) (ID, error) {
	x, err := strconv.Atoi(v)
	if err != nil {
		return UnsetID, fmt.Errorf("failed to parse %q: %w", v, err)
	}
	if x <= 0 {
		return UnsetID, fmt.Errorf("id %v out-of-range", x)
	}
	return ID(x), nil
}
