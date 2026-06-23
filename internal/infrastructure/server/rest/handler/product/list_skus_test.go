package product

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/stretchr/testify/mock"
)

func (suite *ProductHandlerTestSuite) TestListProductSKUs_CheckRequestValidation() {
	for _, testCase := range []struct {
		name               string
		path               string
		expectedStatus     int
		shouldUseCase      bool
		useCaseFirstParam  domain.SKU
		useCaseSecondParam int
		useCaseResponse    []domain.SKU
		useCaseErr         error
	}{
		{
			name:               "valid start_after_sku, valid count",
			path:               "/products/list?start_after_sku=123&count=10",
			expectedStatus:     http.StatusOK,
			shouldUseCase:      true,
			useCaseFirstParam:  domain.SKU(123),
			useCaseSecondParam: 10,
			useCaseResponse:    []domain.SKU{124, 125, 126, 127, 128, 129, 130, 131, 132, 133},
		},
		{
			name:               "use case error",
			path:               "/products/list?start_after_sku=123&count=10",
			expectedStatus:     http.StatusInternalServerError,
			shouldUseCase:      true,
			useCaseFirstParam:  domain.SKU(123),
			useCaseSecondParam: 10,
			useCaseErr:         errors.New("use case error"),
		},
		{
			name:           "empty start_after_sku, valid count",
			path:           "/products/list?start_after_sku=&count=10",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing start_after_sku, valid count",
			path:           "/products/list?count=10",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid start_after_sku, valid count",
			path:           "/products/list?start_after_sku=abc&count=10",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "negative start_after_sku, valid count",
			path:           "/products/list?start_after_sku=-1&count=10",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "zero start_after_sku, valid count",
			path:           "/products/list?start_after_sku=0&count=10",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "too large start_after_sku, valid count",
			path:           "/products/list?start_after_sku=9999999999&count=10",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "valid start_after_sku, empty count",
			path:           "/products/list?start_after_sku=123&count=",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "valid start_after_sku, missing count",
			path:           "/products/list?start_after_sku=123",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "valid start_after_sku, invalid count",
			path:           "/products/list?start_after_sku=123&count=abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "valid start_after_sku, negative count",
			path:           "/products/list?start_after_sku=123&count=-1",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "valid start_after_sku, zero count",
			path:           "/products/list?start_after_sku=123&count=0",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "valid start_after_sku, too large count",
			path:           "/products/list?start_after_sku=123&count=9999999999",
			expectedStatus: http.StatusBadRequest,
		},
	} {
		suite.T().Run(testCase.name, func(t *testing.T) {
			if testCase.shouldUseCase {
				suite.mockUseCase.
					On("GetSKUList", mock.Anything, testCase.useCaseFirstParam, testCase.useCaseSecondParam).
					Return(testCase.useCaseResponse, testCase.useCaseErr).
					Once()
			}

			rec := suite.serveRequest(http.MethodGet, testCase.path)
			suite.Equal(testCase.expectedStatus, rec.Code)
		})
	}
}

func (suite *ProductHandlerTestSuite) TestListProductSKUs_Success() {
	startAfterSKU := domain.SKU(123)
	expectedListSKUs := []domain.SKU{
		domain.SKU(124),
		domain.SKU(125),
		domain.SKU(126),
	}

	suite.mockUseCase.On("GetSKUList", mock.Anything, startAfterSKU, 10).
		Return(expectedListSKUs, nil).
		Once()

	rec := suite.serveRequest(http.MethodGet, "/products/list?start_after_sku=123&count=10")
	suite.Equal(http.StatusOK, rec.Code)
}
