package use_case

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mockRepo "github.com/jbakhtin/marketplace-product/internal/infrastructure/storage/mock"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

// MockLogger для тестов
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(msg string, fields ...any) {
	m.Called(msg, fields)
}

func (m *MockLogger) Info(msg string, fields ...any) {
	m.Called(msg, fields)
}

func (m *MockLogger) Warn(msg string, fields ...any) {
	m.Called(msg, fields)
}

func (m *MockLogger) Error(msg string, fields ...any) {
	m.Called(msg, fields)
}

func (m *MockLogger) Fatal(msg string, fields ...any) {
	m.Called(msg, fields)
}

// TestSuite для группировки тестов
type ProductUseCaseTestSuite struct {
	suite.Suite
	useCase    *ProductUseCase
	mockRepo   *mockRepo.ProductRepository
	mockLogger *MockLogger
}

func (suite *ProductUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(mockRepo.ProductRepository)
	suite.mockLogger = new(MockLogger)
	suite.useCase = &ProductUseCase{
		logger:            suite.mockLogger,
		productRepository: suite.mockRepo,
	}
}

func (suite *ProductUseCaseTestSuite) TearDownTest() {
	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockLogger.AssertExpectations(suite.T())
}

// Тесты для GetProductBySKU
func (suite *ProductUseCaseTestSuite) TestGetProductBySKU_Success() {
	// Arrange
	expectedSKU := domain.SKU(123)
	expectedProduct := domain.Product{
		SKU:   expectedSKU,
		Name:  "Test Product",
		Price: 1000,
	}

	suite.mockRepo.On("GetProductBySKU", mock.Anything, expectedSKU).
		Return(expectedProduct, nil).
		Once()

	// Act
	result, err := suite.useCase.GetProductBySKU(context.Background(), expectedSKU)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedProduct, result)
}

func (suite *ProductUseCaseTestSuite) TestGetProductBySKU_RepositoryError() {
	// Arrange
	expectedSKU := domain.SKU(999)
	expectedError := errors.New("product not found")

	suite.mockRepo.On("GetProductBySKU", mock.Anything, expectedSKU).
		Return(domain.Product{}, expectedError).
		Once()

	// Act
	result, err := suite.useCase.GetProductBySKU(context.Background(), expectedSKU)

	// Assert
	suite.Error(err)
	suite.Equal(expectedError, err)
	suite.Equal(domain.Product{}, result)
}

func (suite *ProductUseCaseTestSuite) TestGetProductBySKU_EmptyProduct() {
	// Arrange
	expectedSKU := domain.SKU(456)
	emptyProduct := domain.Product{}

	suite.mockRepo.On("GetProductBySKU", mock.Anything, expectedSKU).
		Return(emptyProduct, nil).
		Once()

	// Act
	result, err := suite.useCase.GetProductBySKU(context.Background(), expectedSKU)

	// Assert
	suite.NoError(err)
	suite.Equal(emptyProduct, result)
}

// Тесты для GetSKUList
func (suite *ProductUseCaseTestSuite) TestGetSKUList_Success() {
	// Arrange
	startSKU := domain.SKU(100)
	count := 5
	expectedSKUs := []domain.SKU{101, 102, 103, 104, 105}

	suite.mockRepo.On("GetSKUList", mock.Anything, startSKU, count).
		Return(expectedSKUs, nil).
		Once()

	// Act
	result, err := suite.useCase.GetSKUList(context.Background(), startSKU, count)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedSKUs, result)
}

func (suite *ProductUseCaseTestSuite) TestGetSKUList_RepositoryError() {
	// Arrange
	startSKU := domain.SKU(200)
	count := 10
	expectedError := errors.New("database connection failed")

	suite.mockRepo.On("GetSKUList", mock.Anything, startSKU, count).
		Return(nil, expectedError).
		Once()

	// Act
	result, err := suite.useCase.GetSKUList(context.Background(), startSKU, count)

	// Assert
	suite.Error(err)
	suite.Equal(expectedError, err)
	suite.Nil(result)
}

func (suite *ProductUseCaseTestSuite) TestGetSKUList_EmptyResult() {
	// Arrange
	startSKU := domain.SKU(300)
	count := 3
	emptySKUs := []domain.SKU{}

	suite.mockRepo.On("GetSKUList", mock.Anything, startSKU, count).
		Return(emptySKUs, nil).
		Once()

	// Act
	result, err := suite.useCase.GetSKUList(context.Background(), startSKU, count)

	// Assert
	suite.NoError(err)
	suite.Equal(emptySKUs, result)
	suite.Len(result, 0)
}

func (suite *ProductUseCaseTestSuite) TestGetSKUList_ZeroCount() {
	// Arrange
	startSKU := domain.SKU(400)
	count := 0
	expectedSKUs := []domain.SKU{}

	suite.mockRepo.On("GetSKUList", mock.Anything, startSKU, count).
		Return(expectedSKUs, nil).
		Once()

	// Act
	result, err := suite.useCase.GetSKUList(context.Background(), startSKU, count)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedSKUs, result)
}

// Тесты для NewProductUseCase
func (suite *ProductUseCaseTestSuite) TestNewProductUseCase_Success() {
	// Act
	useCase, err := NewProductUseCase(suite.mockLogger, suite.mockRepo)

	// Assert
	suite.NoError(err)
	suite.NotNil(useCase)
	suite.Equal(suite.mockLogger, useCase.logger)
	suite.Equal(suite.mockRepo, useCase.productRepository)
}

// Запуск test suite
func TestProductUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ProductUseCaseTestSuite))
}

// Дополнительные unit тесты без suite
func TestProductUseCase_GetProductBySKU_EdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		sku            domain.SKU
		repoReturn     domain.Product
		repoError      error
		expectedResult domain.Product
		expectedError  bool
	}{
		{
			name:           "zero SKU",
			sku:            0,
			repoReturn:     domain.Product{SKU: 0, Name: "Zero Product", Price: 0},
			repoError:      nil,
			expectedResult: domain.Product{SKU: 0, Name: "Zero Product", Price: 0},
			expectedError:  false,
		},
		{
			name:           "negative SKU",
			sku:            -1,
			repoReturn:     domain.Product{SKU: -1, Name: "Negative Product", Price: 100},
			repoError:      nil,
			expectedResult: domain.Product{SKU: -1, Name: "Negative Product", Price: 100},
			expectedError:  false,
		},
		{
			name:           "large SKU",
			sku:            999999,
			repoReturn:     domain.Product{SKU: 999999, Name: "Large Product", Price: 5000},
			repoError:      nil,
			expectedResult: domain.Product{SKU: 999999, Name: "Large Product", Price: 5000},
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(mockRepo.ProductRepository)
			mockLogger := new(MockLogger)

			mockRepo.On("GetProductBySKU", mock.Anything, tt.sku).
				Return(tt.repoReturn, tt.repoError).
				Once()

			useCase := &ProductUseCase{
				logger:            mockLogger,
				productRepository: mockRepo,
			}

			// Act
			result, err := useCase.GetProductBySKU(context.Background(), tt.sku)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestProductUseCase_GetSKUList_EdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		startSKU       domain.SKU
		count          int
		repoReturn     []domain.SKU
		repoError      error
		expectedResult []domain.SKU
		expectedError  bool
	}{
		{
			name:           "negative count",
			startSKU:       100,
			count:          -5,
			repoReturn:     []domain.SKU{},
			repoError:      nil,
			expectedResult: []domain.SKU{},
			expectedError:  false,
		},
		{
			name:           "large count",
			startSKU:       200,
			count:          10000,
			repoReturn:     []domain.SKU{201, 202, 203},
			repoError:      nil,
			expectedResult: []domain.SKU{201, 202, 203},
			expectedError:  false,
		},
		{
			name:           "single SKU",
			startSKU:       500,
			count:          1,
			repoReturn:     []domain.SKU{501},
			repoError:      nil,
			expectedResult: []domain.SKU{501},
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(mockRepo.ProductRepository)
			mockLogger := new(MockLogger)

			mockRepo.On("GetSKUList", mock.Anything, tt.startSKU, tt.count).
				Return(tt.repoReturn, tt.repoError).
				Once()

			useCase := &ProductUseCase{
				logger:            mockLogger,
				productRepository: mockRepo,
			}

			// Act
			result, err := useCase.GetSKUList(context.Background(), tt.startSKU, tt.count)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
