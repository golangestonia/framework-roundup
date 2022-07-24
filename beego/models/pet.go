package models

import (
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

type PetID int64
type CategoryID int64

type PetStatus string

const UnsetID = 0

type Pet struct {
	ID       PetID     `json:"id"`
	Name     string    `json:"name"`
	Category Category  `json:"category"`
	Status   PetStatus `json:"status,omitempty"`
}

type Category struct {
	ID   CategoryID `json:"id"`
	Name string     `json:"name"`
}

func PetIDFromString(v string) (PetID, error) {
	x, err := strconv.Atoi(v)
	if err != nil {
		return UnsetID, fmt.Errorf("failed to parse %q: %w", v, err)
	}
	if x <= 0 {
		return UnsetID, fmt.Errorf("id %v out-of-range", x)
	}
	return PetID(x), nil
}

func init() {
	orm.RegisterModel(new(Pet))
}
