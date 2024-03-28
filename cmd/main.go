package main

import (
	"log"
	"stock/config"
	"stock/handler"
	"stock/infra/persistence"
	"stock/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// 依存関係を注入
	db, _ := config.GetConn()
	productPersistence := persistence.NewProductPersistence(db)
	productUseCase := usecase.NewProductUseCase(productPersistence)
	productHandler := handler.NewProductHandler(productUseCase)

	engine := gin.Default()

	// 商品関連のエンドポイント
	engine.POST("/products", productHandler.HandleProductCreate)
	engine.GET("/products", productHandler.HandleProductGetByBrand)
	
	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := engine.Run(":8080")
	if err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
