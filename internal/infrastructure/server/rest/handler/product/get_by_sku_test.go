package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jbakhtin/marketplace-product/internal/infrastructure/mock/product"

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

func (suite *ProductHandlerTestSuite) TestGet_CheckRequestValidation() {
	for _, testCase := range []struct {
		name              string
		routeParam        string
		expectedStatus    int
		shouldUseCase     bool
		useCaseFirstParam domain.SKU
		useCaseResponse   domain.Product
		useCaseErr        error
	}{
		{
			name:              "success",
			routeParam:        "sku=123",
			expectedStatus:    http.StatusOK,
			shouldUseCase:     true,
			useCaseFirstParam: domain.SKU(123),
			useCaseResponse:   domain.Product{SKU: 123, Name: "Test Product", Price: 1000},
			useCaseErr:        nil,
		},
		{
			name:           "empty sku",
			routeParam:     "sku=",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "missing sku",
			routeParam:     "",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "missing sku",
			routeParam:     "",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "invalid sku",
			routeParam:     "sku=abc",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "negative sku",
			routeParam:     "sku=-1",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:              "use case error",
			routeParam:        "sku=10",
			expectedStatus:    http.StatusInternalServerError,
			shouldUseCase:     true,
			useCaseFirstParam: domain.SKU(10),
			useCaseResponse:   domain.Product{},
			useCaseErr:        errors.New("use case error"),
		},
	} {
		suite.T().Run(testCase.name, func(t *testing.T) {
			if testCase.shouldUseCase {
				suite.mockUseCase.
					On("GetProductBySKU", mock.Anything, testCase.useCaseFirstParam).
					Return(testCase.useCaseResponse, testCase.useCaseErr).
					Once()
			}

			if testCase.useCaseErr != nil {
				suite.mockLogger.
					On("Error", testCase.useCaseErr.Error()).
					Return().
					Once()
			}

			req := httptest.NewRequest("GET", "/product/get?"+testCase.routeParam, nil)
			w := httptest.NewRecorder()
			suite.handler.Get(w, req)
			suite.Equal(testCase.expectedStatus, w.Code)
		})
	}
}

// Запуск test suite
func TestProductHandlerSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}

// Дополнительные unit тесты без suite
func TestGet_EdgeCases(t *testing.T) {
	t.Parallel()
}
