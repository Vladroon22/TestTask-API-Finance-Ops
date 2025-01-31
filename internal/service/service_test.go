package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Vladroon22/TestTask-Bank-Operation/internal/entity"

	"github.com/Vladroon22/TestTask-Bank-Operation/internal/mocks"
)

func TestIncreaseUserBalance(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockServicer(ctrl)
	service := NewService(mockRepo)

	ctx := context.Background()
	userID := 1
	amount := 100.0

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().IncreaseUserBalance(ctx, userID, amount).Return(nil)

		err := service.IncreaseUserBalance(ctx, userID, amount)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")
		mockRepo.EXPECT().IncreaseUserBalance(ctx, userID, amount).Return(expectedErr)

		err := service.IncreaseUserBalance(ctx, userID, amount)
		assert.Equal(t, expectedErr, err)
	})
}

func TestTransferMoney(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockServicer(ctrl)
	service := NewService(mockRepo)

	ctx := context.Background()
	userFrom := "user1"
	userTo := "user2"
	fromUserID := 1
	toUserID := 2
	amount := 50.0

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().TransferMoney(ctx, userFrom, userTo, fromUserID, toUserID, amount).Return(nil)

		err := service.TransferMoney(ctx, userFrom, userTo, fromUserID, toUserID, amount)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")
		mockRepo.EXPECT().TransferMoney(ctx, userFrom, userTo, fromUserID, toUserID, amount).Return(expectedErr)

		err := service.TransferMoney(ctx, userFrom, userTo, fromUserID, toUserID, amount)
		assert.Equal(t, expectedErr, err)
	})
}

func TestGetLastTxs(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockServicer(ctrl)
	service := NewService(mockRepo)

	ctx := context.Background()
	expectedTxs := []entity.Tx{
		{Sender_name: "user1", Receiver_name: "user2", Amount: 100.0},
		{Sender_name: "user1", Receiver_name: "user2", Amount: 250.0},
		{Sender_name: "user1", Receiver_name: "user2", Amount: 300.0},
		{Sender_name: "user1", Receiver_name: "user2", Amount: 550.0},
		{Sender_name: "user1", Receiver_name: "user2", Amount: 240.0},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetLastTxs(ctx, 1).Return(expectedTxs, nil)

		txs, err := service.GetLastTxs(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, expectedTxs, txs)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")
		mockRepo.EXPECT().GetLastTxs(ctx, -1).Return(nil, expectedErr)

		txs, err := service.GetLastTxs(ctx, -1)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, txs)
	})
}
