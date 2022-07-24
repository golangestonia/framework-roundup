package petdb

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/exp/slices"

	"encore.app/pet"
)

var ErrNotFound = errors.New("not found")

type DB struct {
	mu     sync.Mutex
	pets   []pet.Pet
	cats   []pet.Category
	lastID int64
}

func (db *DB) newID() int64 {
	db.lastID++
	return db.lastID
}

func (db *DB) All(ctx context.Context) ([]pet.Pet, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	return slices.Clone(db.pets), nil
}

func (db *DB) Add(ctx context.Context, p pet.Pet) (pet.ID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	p.ID = pet.ID(db.newID())
	db.ensureCategory(&p.Category)
	db.pets = append(db.pets, p)
	return p.ID, nil
}

func (db *DB) ensureCategory(cat *pet.Category) {
	for _, v := range db.cats {
		if v.Name == cat.Name {
			cat.ID = v.ID
			return
		}
	}
	cat.ID = pet.CategoryID(db.newID())
	db.cats = append(db.cats, *cat)
}

func (db *DB) UpdateByID(ctx context.Context, pet pet.Pet) error {
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

func (db *DB) GetByID(ctx context.Context, petID pet.ID) (pet.Pet, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, pet := range db.pets {
		if pet.ID == petID {
			return pet, nil
		}
	}

	return pet.Pet{}, ErrNotFound
}

func (db *DB) GetByStatuses(ctx context.Context, statuses []pet.Status) ([]pet.Pet, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var xs []pet.Pet
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

func (db *DB) DeleteByID(ctx context.Context, petID pet.ID) error {
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
