package pets

import (
	"context"
	"errors"

	"encore.app/pet"
	"encore.app/petdb"
	"encore.dev/beta/errs"
)

type Pets struct {
	List []pet.Pet
}

type ListParams struct {
	Status []string
}

//encore:api public method=GET path=/pets
func List(ctx context.Context, opts *ListParams) (Pets, error) {
	var filter []pet.Status
	if opts != nil {
		for _, v := range opts.Status {
			filter = append(filter, pet.Status(v))
		}
	}

	var pets []pet.Pet
	var err error
	if len(filter) == 0 {
		pets, err = petdb.All(ctx)
	} else {
		pets, err = petdb.GetByStatuses(ctx, filter)
	}

	if err != nil {
		return Pets{}, errs.WrapCode(err, errs.Internal, "database error")
	}

	return Pets{List: pets}, nil
}

//encore:api public method=GET path=/pet/:id
func Get(ctx context.Context, id int) (*pet.Pet, error) {
	pet, err := petdb.GetByID(ctx, pet.ID(id))
	if err != nil {
		if errors.Is(err, petdb.ErrNotFound) {
			return nil, errs.WrapCode(err, errs.NotFound, "pet not found")
		}
		return nil, errs.WrapCode(err, errs.Internal, "database error")
	}

	return &pet, nil
}

type CreatedPet struct {
	ID pet.ID
}

//encore:api public method=PUT path=/pet
func Create(ctx context.Context, p *pet.Pet) (*CreatedPet, error) {
	if p == nil {
		return nil, &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "pet not specified",
		}
	}

	id, err := petdb.Add(ctx, *p)
	if err != nil {
		return nil, errs.WrapCode(err, errs.Internal, "database error")
	}

	return &CreatedPet{ID: id}, nil
}

//encore:api public method=POST path=/pet/:id
func Update(ctx context.Context, id int, p *pet.Pet) error {
	if p == nil {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "pet not specified",
		}
	}
	p.ID = pet.ID(id)

	err := petdb.UpdateByID(ctx, *p)
	if err != nil {
		if errors.Is(err, petdb.ErrNotFound) {
			return errs.WrapCode(err, errs.NotFound, "pet not found")
		}
		return errs.WrapCode(err, errs.Internal, "database error")
	}

	return nil
}

//encore:api public method=DELETE path=/pet/:id
func Delete(ctx context.Context, id int) error {
	err := petdb.DeleteByID(ctx, pet.ID(id))
	if err != nil {
		if errors.Is(err, petdb.ErrNotFound) {
			return errs.WrapCode(err, errs.NotFound, "pet not found")
		}
		return errs.WrapCode(err, errs.Internal, "database error")
	}

	return nil
}
