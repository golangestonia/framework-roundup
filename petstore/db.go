package petstore

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/exp/slices"
)

var ErrNotFound = errors.New("not found")

type DB struct {
	mu     sync.Mutex
	pets   []Pet
	cats   []Category
	lastID int64
}

func (db *DB) newID() int64 {
	db.lastID++
	return db.lastID
}

func (db *DB) All(ctx context.Context) ([]Pet, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	return slices.Clone(db.pets), nil
}

func (db *DB) Add(ctx context.Context, pet Pet) (PetID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	pet.ID = PetID(db.newID())
	db.ensureCategory(&pet.Category)
	db.pets = append(db.pets, pet)
	return pet.ID, nil
}

func (db *DB) ensureCategory(cat *Category) {
	for _, v := range db.cats {
		if v.Name == cat.Name {
			cat.ID = v.ID
			return
		}
	}
	cat.ID = CategoryID(db.newID())
	db.cats = append(db.cats, *cat)
}

func (db *DB) UpdateByID(ctx context.Context, pet Pet) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i := range db.pets {
		if db.pets[i].ID == pet.ID {
			db.pets[i] = pet
			db.ensureCategory(&pet.Category)
		}
	}

	return ErrNotFound
}

func (db *DB) GetByID(ctx context.Context, petID PetID) (Pet, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, pet := range db.pets {
		if pet.ID == petID {
			return pet, nil
		}
	}

	return Pet{}, ErrNotFound
}

func (db *DB) GetByStatuses(ctx context.Context, statuses []PetStatus) ([]Pet, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var xs []Pet
	for _, pet := range db.pets {
		for _, status := range statuses {
			if pet.Status == status {
				xs = append(xs, pet)
				break
			}
		}
	}

	return xs, nil
}

func (db *DB) DeleteByID(ctx context.Context, petID PetID) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	next := db.pets[:0]
	removed := false
	for _, pet := range db.pets {
		if pet.ID == petID {
			removed = true
			continue
		}
		next = append(next, pet)
	}
	db.pets = next
	if !removed {
		return ErrNotFound
	}

	return nil
}
