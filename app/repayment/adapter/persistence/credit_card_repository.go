package persistence

import (
	"context"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/persistence/dbmodel"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
	"gorm.io/gorm"
)

// CreditCardRepository 信用卡仓储实现
type CreditCardRepository struct {
	DB *gorm.DB // 数据库连接
}

// toDBCreditCard 将业务实体转换为数据库实体
func toDBCreditCard(card *model.CreditCard) *dbmodel.CreditCard {
	return &dbmodel.CreditCard{
		ID:            card.ID,
		UserID:        card.UserID,
		CardNumber:    card.CardNumber,
		BankAccountID: card.BankAccountID,
		AutoRepay:     card.AutoRepay,
		CreditLimit:   card.CreditLimit,
		Balance:       card.Balance,
		DueDate:       card.DueDate,
		CreatedAt:     card.CreatedAt,
	}
}

// toDomainCreditCard 将数据库实体转换为业务实体
func toDomainCreditCard(card *dbmodel.CreditCard) *model.CreditCard {
	return &model.CreditCard{
		ID:            card.ID,
		UserID:        card.UserID,
		CardNumber:    card.CardNumber,
		BankAccountID: card.BankAccountID,
		AutoRepay:     card.AutoRepay,
		CreditLimit:   card.CreditLimit,
		Balance:       card.Balance,
		DueDate:       card.DueDate,
		CreatedAt:     card.CreatedAt,
	}
}

// Save 保存信用卡信息
func (r *CreditCardRepository) Save(card *model.CreditCard) error {
	ctx := context.Background()
	DB := r.DB
	if tx := ctx.Value("tx"); tx != nil {
		DB = tx.(*gorm.DB)
	}
	return DB.Save(toDBCreditCard(card)).Error
}

// FindByID 按 ID 查询信用卡
func (r *CreditCardRepository) FindByID(id string) (*model.CreditCard, error) {
	var dbCard dbmodel.CreditCard
	if err := r.DB.Where("id = ?", id).First(&dbCard).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toDomainCreditCard(&dbCard), nil
}

// FindByUserID 按用户 ID 查询信用卡列表
func (r *CreditCardRepository) FindByUserID(userID string) ([]*model.CreditCard, error) {
	var dbCards []*dbmodel.CreditCard
	if err := r.DB.Where("user_id = ?", userID).Find(&dbCards).Error; err != nil {
		return nil, err
	}
	cards := make([]*model.CreditCard, len(dbCards))
	for i, dbCard := range dbCards {
		cards[i] = toDomainCreditCard(dbCard)
	}
	return cards, nil
}
