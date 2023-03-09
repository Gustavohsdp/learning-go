package database

import (
	"database/sql"

	"github.com/Gustavohsdp/learningGo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (repository *OrderRepository) Save(order *entity.Order) error {
	_, err := repository.Db.Exec("insert into orders (id, price, tax, finalPrice) values(?, ?, ?, ?)",
		order.Id, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) GetTotal() (int, error) {
	var total int
	err := repository.Db.QueryRow("select count(*) from orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
