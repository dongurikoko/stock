package persistence

import (
	"database/sql"
	"fmt"
	"stock/domain/model"
	"stock/domain/repository"
)

type productPersistence struct {
	Conn *sql.DB
}

func NewProductPersistence(conn *sql.DB) repository.ProductRepository {
	return &productPersistence{
		Conn: conn,
	}
}

func (pp productPersistence) Insert(productName string, brandName string, imagePath []string) error {
	// productsテーブルへの挿入とproducts_imagesテーブルへの挿入をトランザクションで行う
	tx, err := pp.Conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction in Insert: %w", err)
	}
	defer tx.Rollback()

	// productsテーブルへの挿入
	result, err := tx.Exec("INSERT INTO products (product_name, brand) VALUES (?, ?)", productName, brandName)
	if err != nil {
		return fmt.Errorf("failed to insert product in Insert: %w", err)
	}

	// 挿入したproductsテーブルのIDを取得
	productID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id in Insert: %w", err)
	}

	// products_imagesテーブルへの挿入
	for _, path := range imagePath {
		if _, err := tx.Exec("INSERT INTO products_images (product_id, image_path) VALUES (?, ?)", productID, path); err != nil {
			return fmt.Errorf("failed to insert product image in Insert: %w", err)
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction in Insert: %w", err)
	}

	return nil

}

// 作成したProductの一覧を取得する
func (pp productPersistence) GetAll() ([]*model.Product, error) {
	// productsテーブルとproducts_imagesテーブルを結合して取得
	rows, err := pp.Conn.Query("SELECT p.id, p.product_name, p.brand, pi.image_path FROM products p LEFT JOIN products_images pi ON p.id = pi.product_id")
	if err != nil {
		return nil, fmt.Errorf("failed to select product in GetAll: %w", err)
	}
	defer rows.Close()

	return convertToProduct(rows)
}

// ブランド名を指定してProductの一覧を取得する
func (pp productPersistence) GetAllByBrand(brandName string) ([]*model.Product, error) {
	// productsテーブルとproducts_imagesテーブルを結合して取得
	rows, err := pp.Conn.Query("SELECT p.id, p.product_name, p.brand, pi.image_path FROM products p LEFT JOIN products_images pi ON p.id = pi.product_id WHERE p.brand = ?", brandName)
	if err != nil {
		return nil, fmt.Errorf("failed to select product in GetAllByBrand: %w", err)
	}
	defer rows.Close()

	return convertToProduct(rows)
}

// rows型データをProduct型に変換する
func convertToProduct(rows *sql.Rows) ([]*model.Product, error) {
	productsMap := make(map[int]*model.Product)
	for rows.Next() {
		var id int
		var productName, brand, imagePath string
		if err := rows.Scan(&id, &productName, &brand, &imagePath); err != nil {
			return nil, fmt.Errorf("failed to scan product in convertToProduct: %w", err)
		}

		// 既に同じIDの商品がマップにあるか確認
		if _, exists := productsMap[id]; !exists {
			productsMap[id] = &model.Product{
				ID:          id,
				ProductName: productName,
				BrandName:       brand,
				ImagePath:   []string{}, // 新しい商品なので空のスライスで初期化
			}
		}

		// imagePathが空ではない場合のみ追加
		if imagePath != "" {
			productsMap[id].ImagePath = append(productsMap[id].ImagePath, imagePath)
		}
	}

	// mapからスライスに変換して結果を返す
	var products []*model.Product
	for _, product := range productsMap {
		products = append(products, product)
	}

	return products, nil
}
