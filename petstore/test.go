package petstore

import "context"

func PopulateTestData(ctx context.Context, db *DB) {
	db.Add(ctx, Pet{
		Name:     "Thomas",
		Category: Category{Name: "dog"},
		Status:   "available",
	})
	db.Add(ctx, Pet{
		Name:     "Toby",
		Category: Category{Name: "dog"},
		Status:   "pending",
	})
	db.Add(ctx, Pet{
		Name:     "Floof",
		Category: Category{Name: "cat"},
		Status:   "sold",
	})
}
