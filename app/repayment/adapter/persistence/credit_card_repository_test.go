package persistence

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/persistence/dbmodel"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
)

// setupDB 设置 SQLite 测试数据库
func setupDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&dbmodel.CreditCard{})
	assert.NoError(t, err)
	return db
}

// TestCreditCardRepository 信用卡仓储测试
func TestCreditCardRepository(t *testing.T) {
	// 测试用例：保存操作
	t.Run("Save", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			db := setupDB(t)
			repo := CreditCardRepository{DB: db}

			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := model.NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			err := repo.Save(card)
			assert.NoError(t, err)

			var savedCard dbmodel.CreditCard
			db.First(&savedCard, "id = ?", "1")
			assert.Equal(t, "user1", savedCard.UserID)
		})
		t.Run("Update", func(t *testing.T) {
			db := setupDB(t)
			repo := CreditCardRepository{DB: db}

			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := model.NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			err := repo.Save(card)
			assert.NoError(t, err)

			card.Balance = 500.0
			err = repo.Save(card)
			assert.NoError(t, err)

			var updatedCard dbmodel.CreditCard
			db.First(&updatedCard, "id = ?", "1")
			assert.Equal(t, 500.0, updatedCard.Balance)
		})
	})
	// 测试用例：查询操作
	t.Run("Find", func(t *testing.T) {
		t.Run("FindByIDSuccess", func(t *testing.T) {
			db := setupDB(t)
			repo := CreditCardRepository{DB: db}

			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := model.NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			db.Create(toDBCreditCard(card))

			foundCard, err := repo.FindByID("1")
			assert.NoError(t, err)
			assert.Equal(t, "user1", foundCard.UserID)
		})
		t.Run("FindByIDNotFound", func(t *testing.T) {
			db := setupDB(t)
			repo := CreditCardRepository{DB: db}

			foundCard, err := repo.FindByID("nonexistent")
			assert.NoError(t, err)
			assert.Nil(t, foundCard)
		})
		t.Run("FindByUserIDSuccess", func(t *testing.T) {
			db := setupDB(t)
			repo := CreditCardRepository{DB: db}

			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := model.NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			db.Create(toDBCreditCard(card))

			cards, err := repo.FindByUserID("user1")
			assert.NoError(t, err)
			assert.Len(t, cards, 1)
			assert.Equal(t, "user1", cards[0].UserID)
		})
	})
}
