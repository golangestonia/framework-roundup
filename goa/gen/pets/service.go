// Code generated by goa v3.7.13, DO NOT EDIT.
//
// pets service
//
// Command:
// $ goa gen github.com/golangestonia/framework-roundup/goa/design

package pets

import (
	"context"

	petsviews "github.com/golangestonia/framework-roundup/goa/gen/pets/views"
)

// The pet management service
type Service interface {
	// AllPets implements AllPets.
	AllPets(context.Context) (res PetCollection, err error)
	// PetByID implements PetByID.
	PetByID(context.Context, *PetByIDPayload) (res *Pet, err error)
	// CreatePet implements CreatePet.
	CreatePet(context.Context, *Pet) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "pets"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"AllPets", "PetByID", "CreatePet"}

// Pet is the result type of the pets service PetByID method.
type Pet struct {
	ID       *int64
	Name     *string
	Category *struct {
		ID   *int64
		Name *string
	}
	Status *string
}

// PetByIDPayload is the payload type of the pets service PetByID method.
type PetByIDPayload struct {
	ID int
}

// PetCollection is the result type of the pets service AllPets method.
type PetCollection []*Pet

// NewPetCollection initializes result type PetCollection from viewed result
// type PetCollection.
func NewPetCollection(vres petsviews.PetCollection) PetCollection {
	return newPetCollection(vres.Projected)
}

// NewViewedPetCollection initializes viewed result type PetCollection from
// result type PetCollection using the given view.
func NewViewedPetCollection(res PetCollection, view string) petsviews.PetCollection {
	p := newPetCollectionView(res)
	return petsviews.PetCollection{Projected: p, View: "default"}
}

// NewPet initializes result type Pet from viewed result type Pet.
func NewPet(vres *petsviews.Pet) *Pet {
	return newPet(vres.Projected)
}

// NewViewedPet initializes viewed result type Pet from result type Pet using
// the given view.
func NewViewedPet(res *Pet, view string) *petsviews.Pet {
	p := newPetView(res)
	return &petsviews.Pet{Projected: p, View: "default"}
}

// newPetCollection converts projected type PetCollection to service type
// PetCollection.
func newPetCollection(vres petsviews.PetCollectionView) PetCollection {
	res := make(PetCollection, len(vres))
	for i, n := range vres {
		res[i] = newPet(n)
	}
	return res
}

// newPetCollectionView projects result type PetCollection to projected type
// PetCollectionView using the "default" view.
func newPetCollectionView(res PetCollection) petsviews.PetCollectionView {
	vres := make(petsviews.PetCollectionView, len(res))
	for i, n := range res {
		vres[i] = newPetView(n)
	}
	return vres
}

// newPet converts projected type Pet to service type Pet.
func newPet(vres *petsviews.PetView) *Pet {
	res := &Pet{
		ID:     vres.ID,
		Name:   vres.Name,
		Status: vres.Status,
	}
	if vres.Category != nil {
		res.Category = &struct {
			ID   *int64
			Name *string
		}{
			ID:   vres.Category.ID,
			Name: vres.Category.Name,
		}
	}
	return res
}

// newPetView projects result type Pet to projected type PetView using the
// "default" view.
func newPetView(res *Pet) *petsviews.PetView {
	vres := &petsviews.PetView{
		ID:     res.ID,
		Name:   res.Name,
		Status: res.Status,
	}
	if res.Category != nil {
		vres.Category = &struct {
			ID   *int64
			Name *string
		}{
			ID:   res.Category.ID,
			Name: res.Category.Name,
		}
	}
	return vres
}