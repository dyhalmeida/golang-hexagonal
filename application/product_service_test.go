package application_test

import (
	"github/com/dyhalmeida/golang-hexagonal/application"
	mock_application "github/com/dyhalmeida/golang-hexagonal/application/mocks"

	// "github.com/go-delve/delve/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
	
}

func TestProductService_Create(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Any product", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	product.EXPECT().Enable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}