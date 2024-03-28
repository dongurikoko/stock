# Stock
## 動作確認手順

1. 環境変数の設定
    ```
    MYSQL_USER=root   
    MYSQL_PASSWORD=stock
    MYSQL_HOST=127.0.0.1
    MYSQL_PORT=3301
    MYSQL_DATABASE=stock_api
    ```
2. `docker-compose up`実行
3. `go run cmd/main.go`実行

## 動作確認例
以下、postmanを使用しました  

例１: GET http://localhost:8080/products

応答：
```
{
    "products": [
        {
            "id": 1,
            "product_name": "A.P.C. 'AURELIA' DENIM DRESS",
            "brand_name": "A.P.C.",
            "image_path": [
                "https://stok.store/cdn/shop/files/20211218140947410_E52---apc---COETKF05822IAL_1_M1.jpg"
            ]
        },
        {
            "id": 2,
            "product_name": "STIVALETTO",
            "brand_name": "Alexander McQUEEN",
            "image_path": [
                "https://stok.store/cdn/shop/files/757487WIDU11000_5_P_2023-07-07T07-27-52.533Z.jpg",
                "https://stok.store/cdn/shop/files/757487WIDU11000_2023-07-07T07-27-52.221Z.jpg"
            ]
        }
    ]
}
```
例２：GET http://localhost:8080/products?brandName=A.P.C.

応答
```
{
    "products": [
        {
            "id": 1,
            "product_name": "A.P.C. 'AURELIA' DENIM DRESS",
            "brand_name": "A.P.C.",
            "image_path": [
                "https://stok.store/cdn/shop/files/20211218140947410_E52---apc---COETKF05822IAL_1_M1.jpg"
            ]
        }
    ]
}
```
例３: POST　http://localhost:8080/products  
（問題なく新規追加できます)
<img width="857" alt="スクリーンショット 2024-03-29 5 38 02" src="https://github.com/dongurikoko/stock/assets/108347471/e5b794d7-81d2-4469-84f1-fc913eb0e8c2">
