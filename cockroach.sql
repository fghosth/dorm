CREATE TABLE IF NOT EXISTS product_information (
    product_id           INT PRIMARY KEY NOT NULL,
    product_name         STRING(50) UNIQUE NOT NULL,
    product_description  STRING(2000),
    category_id          STRING(1) NOT NULL CHECK (category_id IN ('A','B','C')) ,
    weight_class         INT DEFAULT 3 ,
    warranty_period      INT CONSTRAINT valid_warranty CHECK (warranty_period BETWEEN 0 AND 24),
    supplier_id          INT,
    product_status       STRING(20),
    list_price           DECIMAL(8,2),
    min_price            DECIMAL(8,2),
    catalog_url          STRING(50) UNIQUE,
    date_added           DATE DEFAULT CURRENT_DATE(),
    CONSTRAINT price_check CHECK (list_price >= min_price),
    INDEX date_added_idx (date_added),
    INDEX supp_id_prod_status_idx (supplier_id, product_status)
);
