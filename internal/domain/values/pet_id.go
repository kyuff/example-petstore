package values

import (
	"fmt"
)

type PetID int64

func PetIDFromPtr(id *int64) PetID {
	if id == nil {
		return -1
	}

	return PetID(*id)
}

func (id PetID) String() string {
	return fmt.Sprintf("pet-%d", int(id))
}

func (id PetID) Validate() error {
	if id < 0 {
		return fmt.Errorf("pet id cannot be negative: %d", id)
	}

	return nil
}
