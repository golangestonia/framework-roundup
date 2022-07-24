package petdb

import (
	"context"

	"encore.app/pet"
)

var db DB

func All(ctx context.Context) ([]pet.Pet, error) {
	return db.All(ctx)
}

func Add(ctx context.Context, pet pet.Pet) (pet.ID, error) {
	return db.Add(ctx, pet)
}

func UpdateByID(ctx context.Context, pet pet.Pet) error {
	return db.UpdateByID(ctx, pet)
}

func GetByID(ctx context.Context, petID pet.ID) (pet.Pet, error) {
	return db.GetByID(ctx, petID)
}

func GetByStatuses(ctx context.Context, statuses []pet.Status) ([]pet.Pet, error) {
	return db.GetByStatuses(ctx, statuses)
}

func DeleteByID(ctx context.Context, petID pet.ID) error {
	return db.DeleteByID(ctx, petID)
}
