
CREATE TABLE product_information (
	product_id INT NOT NULL,
	product_name STRING(50) NOT NULL,
	product_description STRING(2000) NULL,
	category_id STRING(1) NOT NULL,
	weight_class INT NULL,
	warranty_period INT NULL,
	supplier_id INT NULL,
	product_status STRING(20) NULL,
	list_price DECIMAL(8,2) NULL,
	min_price DECIMAL(8,2) NULL,
	catalog_url STRING(50) NULL,
	date_added DATE NULL DEFAULT '2017-12-20':::DATE,
	CONSTRAINT "primary" PRIMARY KEY (product_id ASC),
	UNIQUE INDEX product_information_product_name_key (product_name ASC),
	UNIQUE INDEX product_information_catalog_url_key (catalog_url ASC),
	INDEX date_added_idx (date_added ASC),
	INDEX supp_id_prod_status_idx (supplier_id ASC, product_status ASC),
	FAMILY "primary" (product_id, product_name, product_description, category_id, weight_class, warranty_period, supplier_id, product_status, list_price, min_price, catalog_url, date_added),
	CONSTRAINT price_check CHECK (list_price >= min_price),
	CONSTRAINT check_category_id CHECK (category_id IN ('A':::STRING, 'B':::STRING, 'C':::STRING)),
	CONSTRAINT valid_warranty CHECK (warranty_period BETWEEN 0 AND 24)
);
