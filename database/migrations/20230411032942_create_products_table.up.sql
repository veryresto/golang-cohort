CREATE TABLE products (
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_category_id` INT NULL,
    `title` VARCHAR ( 255 ) NOT NULL,
    `image` VARCHAR ( 255 )  NULL,
    `video` VARCHAR ( 255 )  NULL,
    `description` VARCHAR ( 255 )  NULL,
    `is_highlighted` boolean DEFAULT 0 NOT NULL,
    `price` INT NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY ( `id` ),
    INDEX idx_products_created_by ( `created_by` ) ,
    INDEX idx_products_updated_by ( `updated_by` ) ,
    INDEX idx_products_product_category_id ( `product_category_id` ) ,
    CONSTRAINT FK_product_product_category_id FOREIGN KEY (`product_category_id`) REFERENCES product_category(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_product_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_product_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`)  ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;