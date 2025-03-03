package repository

import (
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
)

// ICreditCardRepository 定义信用卡仓储接口
type ICreditCardRepository interface {
	Save(card *model.CreditCard) error                       // 保存信用卡
	FindByID(id string) (*model.CreditCard, error)           // 按 ID 查询信用卡
	FindByUserID(userID string) ([]*model.CreditCard, error) // 按用户 ID 查询信用卡列表
}
