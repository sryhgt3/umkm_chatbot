package repository

import (
	"context"
	"umkm-chatbot/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreRepository interface {
	CreateStore(ctx context.Context, store *model.Store) error
	GetStoreByID(ctx context.Context, id int) (*model.Store, error)
	GetAllStores(ctx context.Context) ([]*model.Store, error)
}

type storeRepository struct {
	db *pgxpool.Pool
}

func NewStoreRepository(db *pgxpool.Pool) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) CreateStore(ctx context.Context, store *model.Store) error {
	query := `
		INSERT INTO stores (name, created_at, updated_at)
		VALUES ($1, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query, store.Name).Scan(&store.ID, &store.CreatedAt, &store.UpdatedAt)
	return err
}

func (r *storeRepository) GetStoreByID(ctx context.Context, id int) (*model.Store, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM stores
		WHERE id = $1
	`
	store := &model.Store{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&store.ID,
		&store.Name,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return store, nil
}

func (r *storeRepository) GetAllStores(ctx context.Context) ([]*model.Store, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM stores
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []*model.Store
	for rows.Next() {
		store := &model.Store{}
		err := rows.Scan(
			&store.ID,
			&store.Name,
			&store.CreatedAt,
			&store.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}
