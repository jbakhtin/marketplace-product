package product

import (
	"net/http"
	"testing"

	mockLogger "github.com/jbakhtin/marketplace-product/internal/infrastructure/logger/mock"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/mock/product"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockConfig struct{}

type ProductHandlerTestSuite struct {
	suite.Suite
	handler     *Handler
	mockUseCase *product.MockProductService
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

func (suite *ProductHandlerTestSuite) TestGetProductBySKU_CheckRequestValidation() {
	for _, testCase := range []struct {
		name              string
		path              string
		expectedStatus    int
		shouldUseCase     bool
		useCaseFirstParam domain.SKU
		useCaseResponse   domain.Product
		useCaseErr        error
		shouldLogger      bool
	}{
		{
			name:              "success",
			path:              "/products/get?sku=123",
			expectedStatus:    http.StatusOK,
			shouldUseCase:     true,
			useCaseFirstParam: domain.SKU(123),
			useCaseResponse:   domain.Product{SKU: 123, Name: "Test Product", Price: 1000},
		},
		{
			name:           "empty sku",
			path:           "/products/get?sku=",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing sku",
			path:           "/products/get",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid sku",
			path:           "/products/get?sku=abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "negative sku",
			path:           "/products/get?sku=-1",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "zero sku",
			path:           "/products/get?sku=0",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "too large sku",
			path:           "/products/get?sku=9999999999",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:              "use case internal error",
			path:              "/products/get?sku=10",
			expectedStatus:    http.StatusInternalServerError,
			shouldUseCase:     true,
			useCaseFirstParam: domain.SKU(10),
			useCaseErr:        errors.New("use case internal error"),
			shouldLogger:      true,
		},
		{
			name:              "product not found",
			path:              "/products/get?sku=10",
			expectedStatus:    http.StatusNotFound,
			shouldUseCase:     true,
			useCaseFirstParam: domain.SKU(10),
			useCaseErr:        domain.NotFound,
		},
	} {
		suite.T().Run(testCase.name, func(t *testing.T) {
			if testCase.shouldUseCase {
				suite.mockUseCase.
					On("GetProductBySKU", mock.Anything, testCase.useCaseFirstParam).
					Return(testCase.useCaseResponse, testCase.useCaseErr).
					Once()
			}

			if testCase.shouldLogger {
				suite.mockLogger.
					On("Error", testCase.useCaseErr.Error()).
					Return().
					Once()
			}

			rec := suite.serveRequest(http.MethodGet, testCase.path)
			suite.Equal(testCase.expectedStatus, rec.Code)
		})
	}
}

func TestProductHandlerSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}

func TestGetProductBySKU_EdgeCases(t *testing.T) {
	t.Parallel()
}
