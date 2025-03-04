package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockRepaymentOrderRepository Mock 还款订单仓储
type mockRepaymentOrderRepository struct {
	mock.Mock
}

func (m *mockRepaymentOrderRepository) Save(order *model.RepaymentOrder) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *mockRepaymentOrderRepository) FindByID(id string) (*model.RepaymentOrder, error) {
	args := m.Called(id)
	return args.Get(0).(*model.RepaymentOrder), args.Error(1)
}

func (m *mockRepaymentOrderRepository) FindByCreditCardID(creditCardID string) ([]*model.RepaymentOrder, error) {
	args := m.Called(creditCardID)
	return args.Get(0).([]*model.RepaymentOrder), args.Error(1)
}

// mockCreditCardRepository Mock 信用卡仓储
type mockCreditCardRepository struct {
	mock.Mock
}

// Mock PaymentService 和 PayoutService
type mockPaymentService struct {
	mock.Mock
}

func (m *mockPaymentService) CreatePayment(orderID string, amount float64, bankAccountID string) (string, error) {
	return "mock_cashier_url", nil
}

type mockPayoutService struct {
	mock.Mock
}

func (m *mockPayoutService) CreatePayout(orderID string, creditAmount, feeAmount float64, creditAccountID, feeAccountID string) error {
	return nil
}

func (m *mockCreditCardRepository) Save(card *model.CreditCard) error {
	args := m.Called(card)
	return args.Error(0)
}

func (m *mockCreditCardRepository) FindByID(id string) (*model.CreditCard, error) {
	args := m.Called(id)
	return args.Get(0).(*model.CreditCard), args.Error(1)
}

func (m *mockCreditCardRepository) FindByUserID(userID string) ([]*model.CreditCard, error) {
	args := m.Called(userID)
	return args.Get(0).([]*model.CreditCard), args.Error(1)
}

// mockTransactionManager Mock 事务管理器
type mockTransactionManager struct {
	mock.Mock
}

func (m *mockTransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	args := m.Called(ctx, fn)
	// 模拟事务执行
	if fn != nil {
		return fn(ctx)
	}
	return args.Error(0)
}

// setupRepaymentUseCase 设置测试用的 RepaymentUseCase
func setupRepaymentUseCase() (*RepaymentUseCase, *mockCreditCardRepository, *mockRepaymentOrderRepository, *mockTransactionManager, *mockPaymentService, *mockPayoutService) {
	cardRepo := &mockCreditCardRepository{}
	orderRepo := &mockRepaymentOrderRepository{}
	txManager := &mockTransactionManager{}
	paymentSvc := &mockPaymentService{}
	payoutSvc := &mockPayoutService{}

	uc := &RepaymentUseCase{
		CreditCardRepository:     cardRepo,
		RepaymentOrderRepository: orderRepo,
		TransactionManager:       txManager,
		PaymentService:           paymentSvc,
		PayoutService:            payoutSvc,
	}
	return uc, cardRepo, orderRepo, txManager, paymentSvc, payoutSvc
}

// TestRepaymentUseCase 还款用例测试
func TestRepaymentUseCase(t *testing.T) {
	// 测试用例：手动还款
	t.Run("ManualRepay", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			uc, cardRepo, orderRepo, txManager, paymentSvc, payoutSvc := setupRepaymentUseCase()
			card := &model.CreditCard{ID: "card1", BankAccountID: "bank1"}
			cardRepo.On("FindByID", "card1").Return(card, nil)
			orderRepo.On("Save", mock.Anything).Return(nil).Twice()
			paymentSvc.On("CreatePayment", mock.Anything, 1010.0, "bank1").Return("mock_url", nil)
			txManager.On("WithTransaction", mock.Anything, mock.Anything).Return(nil)

			resp, err := uc.ManualRepay(context.Background(), "card1", 1000.0, 10.0)
			assert.NoError(t, err)
			assert.Equal(t, "mock_url", resp.CashierURL)
			cardRepo.AssertExpectations(t)
			orderRepo.AssertExpectations(t)
			txManager.AssertExpectations(t)
			paymentSvc.AssertExpectations(t)
			payoutSvc.AssertNotCalled(t, "CreatePayout") // 未调用出款服务
		})

		t.Run("CardNotFound", func(t *testing.T) {
			uc, cardRepo, orderRepo, txManager, _, _ := setupRepaymentUseCase()
			cardRepo.On("FindByID", "card1").Return((*model.CreditCard)(nil), errors.New("card not found"))
			txManager.On("WithTransaction", mock.Anything, mock.Anything).Return(errors.New("card not found"))

			_, err := uc.ManualRepay(context.Background(), "card1", 1000.0, 10.0)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "card not found")
			cardRepo.AssertExpectations(t)
			orderRepo.AssertNotCalled(t, "Save")
		})

		t.Run("PaymentServiceError", func(t *testing.T) {
			uc, cardRepo, orderRepo, txManager, paymentSvc, _ := setupRepaymentUseCase()
			card := &model.CreditCard{ID: "card1", BankAccountID: "bank1"}
			cardRepo.On("FindByID", "card1").Return(card, nil)
			orderRepo.On("Save", mock.Anything).Return(nil).Once()
			paymentSvc.On("CreatePayment", mock.Anything, 1010.0, "bank1").Return("", errors.New("payment failed"))
			txManager.On("WithTransaction", mock.Anything, mock.Anything).Return(errors.New("payment failed"))

			_, err := uc.ManualRepay(context.Background(), "card1", 1000.0, 10.0)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "payment failed")
			cardRepo.AssertExpectations(t)
			orderRepo.AssertExpectations(t)
			txManager.AssertExpectations(t)
			paymentSvc.AssertExpectations(t)
		})
	})
}
