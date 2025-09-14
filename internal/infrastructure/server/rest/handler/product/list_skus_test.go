package product

import (
	"errors"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
)

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
