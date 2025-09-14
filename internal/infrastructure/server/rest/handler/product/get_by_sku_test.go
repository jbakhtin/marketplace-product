package product

import (
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/mock/product"
	"net/http"
	"net/http/httptest"
	"testing"

	mockLogger "github.com/jbakhtin/marketplace-product/internal/infrastructure/logger/mock"
	mockUseCase "github.com/jbakhtin/marketplace-product/internal/infrastructure/mock/product"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockConfig struct{}

// TestSuite для группировки тестов
type ProductHandlerTestSuite struct {
	suite.Suite
	handler     Handler
	mockUseCase *mockUseCase.MockProductService
	mockLogger  *mockLogger.MockLogger
}

func (suite *ProductHandlerTestSuite) SetupTest() {
	suite.mockLogger = new(mockLogger.MockLogger)
	suite.mockUseCase = new(product.MockProductService)
	suite.handler, _ = NewProductHandler(MockConfig{}, suite.mockLogger, suite.mockUseCase)
}

func (suite *ProductHandlerTestSuite) TearDownTest() {
	suite.mockLogger.AssertExpectations(suite.T())
	suite.mockUseCase.AssertExpectations(suite.T())
}

// Тесты для GetProductBySKU
func (suite *ProductHandlerTestSuite) TestGet_Success() {
	// Arrange
	expectedSKU := domain.SKU(123)
	expectedProduct := domain.Product{
		SKU:   expectedSKU,
		Name:  "Test Product",
		Price: 1000,
	}

	suite.mockUseCase.On("GetProductBySKU", mock.Anything, expectedSKU).
		Return(expectedProduct, nil).
		Once()

	// Act
	req := httptest.NewRequest("GET", "/product/get?sku=123", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGet_EmptySKU() {
	// Act - пустой sku параметр
	req := httptest.NewRequest("GET", "/product/get?sku=", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGet_MissingSKU() {
	// Act - отсутствует sku параметр
	req := httptest.NewRequest("GET", "/product/get", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGet_InvalidSKU() {
	// Act - невалидный sku (не число)
	req := httptest.NewRequest("GET", "/product/get?sku=abc", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGet_UseCaseError() {
	// Arrange
	expectedSKU := domain.SKU(999)
	expectedError := errors.New("product not found")

	suite.mockUseCase.On("GetProductBySKU", mock.Anything, expectedSKU).
		Return(domain.Product{}, expectedError).
		Once()

	suite.mockLogger.On("Error", "product not found").Return().Once()

	// Act
	req := httptest.NewRequest("GET", "/product/get?sku=999", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGet_ZeroSKU() {
	// Act
	req := httptest.NewRequest("GET", "/product/get?sku=0", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGet_NegativeSKU() {
	// Act
	req := httptest.NewRequest("GET", "/product/get?sku=-1", nil)
	w := httptest.NewRecorder()
	suite.handler.Get(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

// Запуск test suite
func TestProductHandlerSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}

// Дополнительные unit тесты без suite
func TestGet_EdgeCases(t *testing.T) {
	t.Parallel()
}
