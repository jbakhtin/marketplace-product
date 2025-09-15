package product

import (
	"errors"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (suite *ProductHandlerTestSuite) TestGetListSKUs_CheckRequestValidation() {
	for _, testCase := range []struct {
		name               string
		routeParam         string
		expectedStatus     int
		shouldUseCase      bool
		useCaseFirstParam  domain.SKU
		useCaseSecondParam int
		useCaseResponse    []domain.SKU
		useCaseErr         error
		shouldLogger       bool
	}{
		{
			name:               "valid start_after_sku, valid count",
			routeParam:         "start_after_sku=123&count=10",
			expectedStatus:     http.StatusOK,
			shouldUseCase:      true,
			useCaseFirstParam:  domain.SKU(123),
			useCaseSecondParam: 10,
			useCaseResponse:    []domain.SKU{124, 125, 126, 127, 128, 129, 130, 131, 132, 133},
			useCaseErr:         nil,
		},
		{
			name:               "use case error",
			routeParam:         "start_after_sku=123&count=10",
			expectedStatus:     http.StatusInternalServerError,
			shouldUseCase:      true,
			useCaseFirstParam:  domain.SKU(123),
			useCaseSecondParam: 10,
			useCaseResponse:    nil,
			useCaseErr:         errors.New("use case error"),
		},
		{
			name:           "empty start_after_sku, valid count",
			routeParam:     "start_after_sku=&count=10",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "missing start_after_sku, valid count",
			routeParam:     "count=10",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "invalid start_after_sku, valid count",
			routeParam:     "start_after_sku=abc&count=10",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "negative start_after_sku, valid count",
			routeParam:     "start_after_sku=-1&count=10",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "zero start_after_sku, valid count",
			routeParam:     "start_after_sku=0&count=10",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "too large start_after_sku, valid count",
			routeParam:     "start_after_sku=9999999999&count=10",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "valid start_after_sku, empty count",
			routeParam:     "start_after_sku=123&count=",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "valid start_after_sku, missing count",
			routeParam:     "start_after_sku=123",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "valid start_after_sku, invalid count",
			routeParam:     "start_after_sku=123&count=abc",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "valid start_after_sku, negative count",
			routeParam:     "start_after_sku=123&count=-1",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "valid start_after_sku, zero count",
			routeParam:     "start_after_sku=123&count=0",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
		{
			name:           "valid start_after_sku, too large count",
			routeParam:     "start_after_sku=123&count=9999999999",
			expectedStatus: http.StatusBadRequest,
			shouldUseCase:  false,
		},
	} {
		suite.T().Run(testCase.name, func(t *testing.T) {
			if testCase.shouldUseCase {
				suite.mockUseCase.
					On("GetSKUList", mock.Anything, testCase.useCaseFirstParam, testCase.useCaseSecondParam).
					Return(testCase.useCaseResponse, testCase.useCaseErr).
					Once()
			}

			req := httptest.NewRequest("GET", "/product/list?"+testCase.routeParam, nil)
			w := httptest.NewRecorder()
			suite.handler.GetListSKUs(w, req)
			suite.Equal(testCase.expectedStatus, w.Code)
		})
	}
}

// Тесты для GetProductBySKU
func (suite *ProductHandlerTestSuite) TestGetListSKUs_Success() {
	startAfterSKU := domain.SKU(123)
	expectedListSKUs := []domain.SKU{
		domain.SKU(124),
		domain.SKU(125),
		domain.SKU(126),
		domain.SKU(127),
		domain.SKU(128),
		domain.SKU(129),
		domain.SKU(130),
		domain.SKU(130),
		domain.SKU(131),
		domain.SKU(132),
		domain.SKU(133),
	}

	suite.mockUseCase.On("GetSKUList", mock.Anything, startAfterSKU, 10).
		Return(expectedListSKUs, nil).
		Once()

	// Act
	req := httptest.NewRequest("GET", "/product/list?start_after_sku=123&count=10", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetListSKUs_EmptySKU() {
	// Act - пустой sku параметр
	req := httptest.NewRequest("GET", "/product/list?start_after_sku=", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetListSKUs_MissingSKU() {
	// Act - отсутствует sku параметр
	req := httptest.NewRequest("GET", "/product/list", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetListSKUs_InvalidSKU() {
	// Act - невалидный sku (не число)
	req := httptest.NewRequest("GET", "/product/list?start_after_sku=abc", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetListSKUs_UseCaseError() {
	// Arrange
	expectedSKU := domain.SKU(999)
	expectedError := errors.New("product not found")

	suite.mockUseCase.On("GetSKUList", mock.Anything, expectedSKU, 10).
		Return([]domain.SKU{}, expectedError).
		Once()

	// Act
	req := httptest.NewRequest("GET", "/product/list?start_after_sku=999&count=10", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetListSKUs_ZeroStartAfterSKU() {
	// Act
	req := httptest.NewRequest("GET", "/product/list?start_after_sku=0&count=1", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetListSKUs_NegativeSKU() {
	// Act
	req := httptest.NewRequest("GET", "/product/list?start_after_sku=-1", nil)
	w := httptest.NewRecorder()
	suite.handler.GetListSKUs(w, req)

	// Assert
	suite.Equal(http.StatusBadRequest, w.Code)
}
