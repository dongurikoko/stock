package handler

import (
	"net/http"
	"stock/domain/model"
	"stock/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	HandleProductCreate(c *gin.Context)
	HandleProductGetByBrand(c *gin.Context)
}

type productHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(pu usecase.ProductUseCase) ProductHandler {
	return &productHandler{
		productUseCase: pu,
	}
}

// 商品の新規作成
func (ph productHandler) HandleProductCreate(c *gin.Context) {
	// フォームデータから取得
	productName := c.PostForm("productName")
	brandName := c.PostForm("brandName")
	imagePath := c.PostFormArray("imagePath")

	if productName == "" || brandName == "" || len(imagePath) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productName, brandName, imagePath are required"})
		return
	}

	if err := ph.productUseCase.Insert(productName, brandName, imagePath); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

// 商品の取得
func (ph productHandler) HandleProductGetByBrand(c *gin.Context) {
	type productResponse struct {
		ID        int      `json:"id"`
		ProductName string `json:"product_name"`
		BrandName string `json:"brand_name"`
		ImagePath []string `json:"image_path"`
	}
	type productsResponse struct {
		Products []*productResponse `json:"products"`
	}

	brandName := c.Query("brandName")

	var products []*model.Product
	var err error
	
	if brandName == "" { // ブランド名が指定されていない場合は全ての商品を取得
		products, err = ph.productUseCase.GetAll()
	}else{	// ブランド名が指定されている場合はそのブランドの商品を取得
		products, err = ph.productUseCase.GetAllByBrand(brandName)
	}
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	var response productsResponse
	for _, product := range products {
		response.Products = append(response.Products, &productResponse{
			ID: product.ID,
			ProductName: product.ProductName,
			BrandName: product.BrandName,
			ImagePath: product.ImagePath,
		})
	}

	c.JSON(http.StatusOK, response)
}
