package persistence

import (
	"testing"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/persistence/dbmodel"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDBRepaymentOrder 设置 SQLite 测试数据库
func setupDBRepaymentOrder(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&dbmodel.RepaymentOrder{}, &dbmodel.RepaymentOrderItem{})
	assert.NoError(t, err)
	return db
}

// TestRepaymentOrderRepository 还款订单仓储测试
func TestRepaymentOrderRepository(t *testing.T) {
	// 测试用例：保存操作
	t.Run("Save", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			db := setupDBRepaymentOrder(t)
			repo := RepaymentOrderRepository{DB: db}

			order, _ := model.NewRepaymentOrder("1", "card1", 1010.0)
			debitItem, _ := model.NewRepaymentOrderItem("d1", "1", model.ItemTypeDebit, 1010.0, "bank1")
			order.StartPayment(*debitItem)

			err := repo.Save(order)
			assert.NoError(t, err)

			var savedOrder dbmodel.RepaymentOrder
			db.First(&savedOrder, "id = ?", "1")
			assert.Equal(t, "Paying", savedOrder.Status)

			var savedItem dbmodel.RepaymentOrderItem
			db.First(&savedItem, "id = ?", "d1")
			assert.Equal(t, "Debit", savedItem.Type)
		})

		t.Run("Update", func(t *testing.T) {
			db := setupDBRepaymentOrder(t)
			repo := RepaymentOrderRepository{DB: db}

			order, _ := model.NewRepaymentOrder("1", "card1", 1010.0)
			debitItem, _ := model.NewRepaymentOrderItem("d1", "1", model.ItemTypeDebit, 1010.0, "bank1")
			order.StartPayment(*debitItem)
			err := repo.Save(order)
			assert.NoError(t, err)

			order.ConfirmPayment()
			err = repo.Save(order)
			assert.NoError(t, err)

			var updatedOrder dbmodel.RepaymentOrder
			db.First(&updatedOrder, "id = ?", "1")
			assert.Equal(t, "Success", updatedOrder.Status)
		})
	})

	// 测试用例：查询操作
	t.Run("Find", func(t *testing.T) {
		t.Run("FindByIDSuccess", func(t *testing.T) {
			db := setupDBRepaymentOrder(t)
			repo := RepaymentOrderRepository{DB: db}

			order, _ := model.NewRepaymentOrder("1", "card1", 1010.0)
			debitItem, _ := model.NewRepaymentOrderItem("d1", "1", model.ItemTypeDebit, 1010.0, "bank1")
			order.StartPayment(*debitItem)
			db.Create(toDBRepaymentOrder(order))
			db.Create(toDBRepaymentOrderItem(debitItem))

			foundOrder, err := repo.FindByID("1")
			assert.NoError(t, err)
			assert.Equal(t, "card1", foundOrder.CreditCardID)
			assert.Len(t, foundOrder.Items, 1)
			assert.Equal(t, "d1", foundOrder.Items[0].ID)
		})

		t.Run("FindByIDNotFound", func(t *testing.T) {
			db := setupDBRepaymentOrder(t)
			repo := RepaymentOrderRepository{DB: db}

			foundOrder, err := repo.FindByID("nonexistent")
			assert.NoError(t, err)
			assert.Nil(t, foundOrder)
		})

		t.Run("FindByCreditCardIDSuccess", func(t *testing.T) {
			db := setupDBRepaymentOrder(t)
			repo := RepaymentOrderRepository{DB: db}

			order, _ := model.NewRepaymentOrder("1", "card1", 1010.0)
			db.Create(toDBRepaymentOrder(order))

			orders, err := repo.FindByCreditCardID("card1")
			assert.NoError(t, err)
			assert.Len(t, orders, 1)
			assert.Equal(t, "card1", orders[0].CreditCardID)
		})
	})
}
