package database

import (
	"database/sql"

	"github.com/leandrosouzasnts/go-intensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)", order.Id, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)

	if err != nil {
		return 0, nil
	}

	return total, nil
}
