package database

import (
	"database/sql"
	"testing"

	"github.com/Gustavohsdp/learningGo/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, finalPrice float NOT NULL, PRIMARY KEY (id))")
	suite.db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.db.Close()
}

func TestSuite(testing *testing.T) {
	suite.Run(testing, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestSavingOrder() {
	order, err := entity.NewOrder("123", 10.0, 1.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order

	err = suite.db.QueryRow("select id, price, tax, finalPrice from orders where id = ?",
		order.Id).Scan(&orderResult.Id, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.Id, orderResult.Id)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
