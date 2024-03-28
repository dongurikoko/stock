CREATE SCHEMA IF NOT EXISTS `stock_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `stock_api` ;

SET CHARSET utf8mb4;
-- -----------------------------------------------------
-- Table `stock_api`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `stock_api`.`products` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'プロダクトID',
  `product_name` VARCHAR(255) NOT NULL COMMENT 'プロダクト名',
  `brand` VARCHAR(255) NOT NULL COMMENT 'ブランド名',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'プロダクト情報';

-- -----------------------------------------------------
-- Table `stock_api`.`products_images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `stock_api`.`products_images` (
  `image_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '画像ID',
  `product_id` INT UNSIGNED NOT NULL COMMENT 'プロダクトID',
  `image_path` VARCHAR(255) NOT NULL COMMENT '画像パス',
  Foreign Key (`product_id`) REFERENCES `products`(`id`),
  PRIMARY KEY (`image_id`))
ENGINE = InnoDB
COMMENT = '画像情報';
