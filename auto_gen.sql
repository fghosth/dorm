


	 CREATE TABLE `hs_auth_api` (
 		`id` int(10)   unsigned NOT NULL AUTO_INCREMENT,
 		`api_key` varchar(128) NOT NULL,
 		`name` varchar(256) NOT NULL,
 		`type` tinyint(4) NOT NULL DEFAULT '0',
 		`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`deleted_at` timestamp DEFAULT NULL,
 		`status_at` tinyint(4) NOT NULL DEFAULT '1',
		 PRIMARY KEY (`id`)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	


	 CREATE TABLE `hs_auth_application` (
 		`id` int(10)   unsigned NOT NULL AUTO_INCREMENT,
 		`secret_key` varchar(128) NOT NULL,
 		`app_key` varchar(128) NOT NULL,
 		`name` varchar(256) NOT NULL,
 		`ip` varchar(32) NOT NULL DEFAULT '',
 		`type` tinyint(4) NOT NULL DEFAULT '0',
 		`exp` int(11) NOT NULL DEFAULT '0',
 		`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`deleted_at` timestamp DEFAULT NULL,
 		`status_at` tinyint(4) NOT NULL DEFAULT '1',
		 PRIMARY KEY (`id`)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	


	 CREATE TABLE `hs_auth_permission` (
 		`id` int(10)   unsigned NOT NULL AUTO_INCREMENT,
 		`app_key` varchar(128) NOT NULL,
 		`api_key` varchar(256) NOT NULL,
 		`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`deleted_at` timestamp DEFAULT NULL,
 		`status_at` tinyint(4) NOT NULL DEFAULT '1',
		 PRIMARY KEY (`id`)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	

	 CREATE TABLE `hs_auth_application` (
 		`id` int(10)   unsigned NOT NULL AUTO_INCREMENT,
 		`secret_key` varchar(128) NOT NULL,
 		`app_key` varchar(128) NOT NULL,
 		`name` varchar(256) NOT NULL,
 		`ip` varchar(32) NOT NULL DEFAULT '',
 		`type` tinyint(4) NOT NULL DEFAULT '0',
 		`exp` int(11) NOT NULL DEFAULT '0',
 		`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`deleted_at` timestamp DEFAULT NULL,
 		`status_at` tinyint(4) NOT NULL DEFAULT '1',
		 PRIMARY KEY (`id`)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	


	 CREATE TABLE `hs_auth_records` (
 		`id` int(10)   unsigned NOT NULL AUTO_INCREMENT,
 		`secret_key` varchar(128) NOT NULL,
 		`app_key` varchar(128) NOT NULL,
 		`sign` varchar(128) NOT NULL DEFAULT '',
 		`token` varchar(256) NOT NULL,
 		`alg` varchar(64) NOT NULL,
 		`ip` varchar(32) NOT NULL DEFAULT '',
 		`exp` timestamp DEFAULT NULL,
 		`iat` timestamp DEFAULT NULL,
 		`type` tinyint(4) NOT NULL DEFAULT '0',
 		`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		`deleted_at` timestamp DEFAULT NULL,
 		`status_at` tinyint(4) NOT NULL DEFAULT '1',
		 PRIMARY KEY (`id`)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	


	 CREATE TABLE `hs_migrations` (
 		`id` int(10)   unsigned NOT NULL AUTO_INCREMENT,
 		`migration` varchar(255) NOT NULL,
 		`batch` int(11) NOT NULL,
		 PRIMARY KEY (`id`)
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	


	 CREATE TABLE `product_information` (
 		`product_id` bigint NOT NULL,
 		`product_name` varchar NOT NULL,
 		`product_description` varchar NULL,
 		`category_id` varchar NOT NULL,
 		`weight_class` bigint NULL,
 		`warranty_period` bigint NULL,
 		`supplier_id` bigint NULL,
 		`product_status` varchar NULL,
 		`list_price` double NULL,
 		`min_price` double NULL,
 		`catalog_url` varchar NULL,
 		`date_added` varchar 
		 
	 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';
	