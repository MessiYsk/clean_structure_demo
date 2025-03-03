package persistence

import (
	"context"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/persistence/dbmodel"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
	"gorm.io/gorm"
)

// RepaymentOrderRepository 还款订单仓储实现
type RepaymentOrderRepository struct {
	db *gorm.DB // 数据库连接
}

// toDBRepaymentOrder 将业务实体转换为数据库主单实体
func toDBRepaymentOrder(order *model.RepaymentOrder) *dbmodel.RepaymentOrder {
	return &dbmodel.RepaymentOrder{
		ID:           order.ID,
		CreditCardID: order.CreditCardID,
		TotalAmount:  order.TotalAmount,
		Status:       string(order.Status),
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

// toDBRepaymentOrderItem 将业务子单转换为数据库子单实体
func toDBRepaymentOrderItem(item *model.RepaymentOrderItem) *dbmodel.RepaymentOrderItem {
	return &dbmodel.RepaymentOrderItem{
		ID:            item.ID,
		OrderID:       item.OrderID,
		Type:          string(item.Type),
		Amount:        item.Amount,
		AccountID:     item.AccountID,
		Status:        string(item.Status),
		TransactionID: item.TransactionID,
		CreatedAt:     item.CreatedAt,
	}
}

// toDomainRepaymentOrder 将数据库主单实体转换为业务实体
func toDomainRepaymentOrder(order *dbmodel.RepaymentOrder, items []model.RepaymentOrderItem) *model.RepaymentOrder {
	return &model.RepaymentOrder{
		ID:           order.ID,
		CreditCardID: order.CreditCardID,
		TotalAmount:  order.TotalAmount,
		Status:       model.OrderStatus(order.Status),
		Items:        items,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

// toDomainRepaymentOrderItem 将数据库子单实体转换为业务子单
func toDomainRepaymentOrderItem(item *dbmodel.RepaymentOrderItem) *model.RepaymentOrderItem {
	return &model.RepaymentOrderItem{
		ID:            item.ID,
		OrderID:       item.OrderID,
		Type:          model.ItemType(item.Type),
		Amount:        item.Amount,
		AccountID:     item.AccountID,
		Status:        model.ItemStatus(item.Status),
		TransactionID: item.TransactionID,
		CreatedAt:     item.CreatedAt,
	}
}

// Save 保存还款订单（主单和子单分开存储）
func (r *RepaymentOrderRepository) Save(order *model.RepaymentOrder) error {
	ctx := context.Background()
	db := r.db
	if tx := ctx.Value("tx"); tx != nil {
		db = tx.(*gorm.DB)
	}

	// 保存主单
	if err := db.Save(toDBRepaymentOrder(order)).Error; err != nil {
		return err
	}

	// 保存子单
	for _, item := range order.Items {
		if err := db.Save(toDBRepaymentOrderItem(&item)).Error; err != nil {
			return err
		}
	}
	return nil
}

// FindByID 按 ID 查询还款订单
func (r *RepaymentOrderRepository) FindByID(id string) (*model.RepaymentOrder, error) {
	var dbOrder dbmodel.RepaymentOrder
	if err := r.db.Where("id = ?", id).First(&dbOrder).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	// 查询子单
	var dbItems []dbmodel.RepaymentOrderItem
	if err := r.db.Where("order_id = ?", id).Find(&dbItems).Error; err != nil {
		return nil, err
	}
	items := make([]model.RepaymentOrderItem, len(dbItems))
	for i, dbItem := range dbItems {
		items[i] = *toDomainRepaymentOrderItem(&dbItem)
	}

	return toDomainRepaymentOrder(&dbOrder, items), nil
}

// FindByCreditCardID 按信用卡 ID 查询订单列表
func (r *RepaymentOrderRepository) FindByCreditCardID(creditCardID string) ([]*model.RepaymentOrder, error) {
	var dbOrders []dbmodel.RepaymentOrder
	if err := r.db.Where("credit_card_id = ?", creditCardID).Find(&dbOrders).Error; err != nil {
		return nil, err
	}

	orders := make([]*model.RepaymentOrder, len(dbOrders))
	for i, dbOrder := range dbOrders {
		var dbItems []dbmodel.RepaymentOrderItem
		if err := r.db.Where("order_id = ?", dbOrder.ID).Find(&dbItems).Error; err != nil {
			return nil, err
		}
		items := make([]model.RepaymentOrderItem, len(dbItems))
		for j, dbItem := range dbItems {
			items[j] = *toDomainRepaymentOrderItem(&dbItem)
		}
		orders[i] = toDomainRepaymentOrder(&dbOrder, items)
	}
	return orders, nil
}
