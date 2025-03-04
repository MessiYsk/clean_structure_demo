package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCreditCard 信用卡实体测试
func TestCreditCard(t *testing.T) {
	// 测试用例：构造相关
	t.Run("Construction", func(t *testing.T) {
		t.Run("ValidNewCreditCard", func(t *testing.T) {
			dueDate := time.Now().Add(24 * time.Hour)
			card, err := NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			assert.NoError(t, err)
			assert.Equal(t, "1", card.ID)
			assert.Equal(t, "user1", card.UserID)
			assert.Equal(t, true, card.AutoRepay)
			assert.Equal(t, 0.0, card.Balance)
		})
		t.Run("InvalidUserID", func(t *testing.T) {
			dueDate := time.Now().Add(24 * time.Hour)
			_, err := NewCreditCard("1", "", "1234567890123456", "bank1", 1000.0, dueDate, true)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "用户 ID")
		})
		t.Run("InvalidCreditLimit", func(t *testing.T) {
			dueDate := time.Now().Add(24 * time.Hour)
			_, err := NewCreditCard("1", "user1", "1234567890123456", "bank1", -1000.0, dueDate, true)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "信用额度必须为正数")
		})
	})

	// 测试用例：行为相关
	t.Run("Behavior", func(t *testing.T) {
		t.Run("UpdateBalanceValid", func(t *testing.T) {
			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			err := card.UpdateBalance(500.0)
			assert.NoError(t, err)
			assert.Equal(t, 500.0, card.Balance)
		})
		t.Run("UpdateBalanceExceedLimit", func(t *testing.T) {
			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			err := card.UpdateBalance(-2000.0)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "超过信用额度")
		})
		t.Run("UpdateBalancePositive", func(t *testing.T) {
			dueDate := time.Now().Add(24 * time.Hour)
			card, _ := NewCreditCard("1", "user1", "1234567890123456", "bank1", 1000.0, dueDate, true)
			err := card.UpdateBalance(1500.0)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "信用卡余额不能为正")
		})
	})
}
