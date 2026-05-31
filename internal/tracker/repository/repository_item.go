package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/valyala/fasthttp"
	"job4j.ru/go-lang-base/internal/tracker"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

func (r *RepoPg) Create(ctx context.Context, it tracker.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`insert into items(id, name) values($1, $2)`,
		it.ID, it.Name,
	)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}
	return nil
}

func (r *RepoPg) List(ctx context.Context) ([]tracker.Item, error) {
	rows, err := r.pool.Query(ctx, `select id, name from items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) Get(ctx context.Context, id string) (tracker.Item, error) {
	var it tracker.Item
	err := r.pool.QueryRow(
		ctx,
		`select id, name from items where id = $1`,
		id,
	).Scan(&it.ID, &it.Name)

	return it, err
}

func (r *RepoPg) FindByName(ctx context.Context, name string) ([]tracker.Item, error) {
	rows, err := r.pool.Query(ctx, `select id, name from items where name like '%$1%'`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) Delete(ctx *fasthttp.RequestCtx, id string) error {
	deleted, err := r.pool.Exec(ctx, `delete from items where id = $1`, id)
	if err != nil {
		log.Errorw("failed to delete item: %s", id)
		return fmt.Errorf("database error while deleting item %s: %w", id, err)
	}

	if deleted.RowsAffected() == 0 {
		log.Errorw("item with id %s not found", id)
	}

	return nil
}

func (r *RepoPg) Update(ctx *fasthttp.RequestCtx, id string, item tracker.Item) error {
	updated, err := r.pool.Exec(ctx, `update items set name=$1 where id = $2`, item.Name, id)
	if err != nil {
		log.Errorw("failed to update item: %s", id)
		return fmt.Errorf("database error while updating item %s: %w", id, err)
	}

	if updated.RowsAffected() == 0 {
		log.Errorw("item with id %s not found", id)
		return fmt.Errorf("item with id %s not found", id)
	}

	return nil
}
