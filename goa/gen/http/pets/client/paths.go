// Code generated by goa v3.7.13, DO NOT EDIT.
//
// HTTP request path constructors for the pets service.
//
// Command:
// $ goa gen github.com/golangestonia/framework-roundup/goa/design

package client

import (
	"fmt"
)

// AllPetsPetsPath returns the URL path to the pets service AllPets HTTP endpoint.
func AllPetsPetsPath() string {
	return "/pets"
}

// PetByIDPetsPath returns the URL path to the pets service PetByID HTTP endpoint.
func PetByIDPetsPath(id int) string {
	return fmt.Sprintf("/pet/%v", id)
}

// CreatePetPetsPath returns the URL path to the pets service CreatePet HTTP endpoint.
func CreatePetPetsPath() string {
	return "/pet"
}
