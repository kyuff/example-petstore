package values

import "github.com/kyuff/validate"

type PetStatus string

const (
	PetStatusAVAILABLE PetStatus = "AVAILABLE"
	PetStatusPENDING   PetStatus = "PENDING"
	PetStatusSOLD      PetStatus = "SOLD"
)

var AllPetStatus = []PetStatus{
	PetStatusAVAILABLE,
	PetStatusPENDING,
	PetStatusSOLD,
}

func (v PetStatus) String() string {
	return string(v)
}

func (v PetStatus) Validate() error {
	return validate.SliceContainsf(AllPetStatus, v, "PetStatus not one of %v", AllPetStatus)
}
