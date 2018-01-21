

CREATE TABLE IF NOT EXISTS user_dd (
 		uid SERIAL NOT NULL DEFAULT current_timestamp()
 		name STRING NOT NULL DEFAULT 'aaa':::STRING
 		gender SMALLINT NOT NULL
 		password STRING NOT NULL
 		qq SMALLINT NOT NULL
 		account STRING NOT NULL
 		cellphone STRING NOT NULL
 		happy STRING NOT NULL DEFAULT 'aaa':::STRING
 		cash REAL NOT NULL DEFAULT '0':::REAL
 		create_time TIMESTAMP NOT NULL
 		update_time TIMESTAMP NOT NULL
		PRIMARY KEY (uid,qq)
		UNIQUE (uid,name,aaa)
		INDEX user_dd_index_1 (updated_at,status_at)
	 );
	 CREATE TABLE IF NOT EXISTS hs_auth_api (
 		id INT PRIMARY KEY  NOT NULL,
 		api_key varchar NOT NULL,
 		name varchar NOT NULL,
 		type int NOT NULL DEFAULT '0',
 		created_at timestamp NOT NULL DEFAULT current_timestamp(),
 		updated_at int NOT NULL,
 		deleted_at int DEFAULT NULL,
 		status_at int NOT NULL DEFAULT '1'
	 );



	 CREATE TABLE IF NOT EXISTS hs_auth_application (
 		id serial PRIMARY key NOT NULL,
 		secret_key varchar NOT NULL,
 		app_key varchar NOT NULL,
 		name varchar NOT NULL,
 		ip varchar NOT NULL DEFAULT '',
 		type int NOT NULL DEFAULT '0',
 		exp varchar NOT NULL DEFAULT '0',
 		created_at varchar NOT NULL,
 		updated_at varchar NOT NULL,
 		deleted_at varchar DEFAULT NULL,
 		status_at int NOT NULL DEFAULT '1'
	 );



	 CREATE TABLE IF NOT EXISTS hs_auth_permission (
 		id int PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT,
 		app_key varchar NOT NULL,
 		api_key varchar NOT NULL,
 		created_at int NOT NULL;DEFAULT CURRENT_TIMESTAMP,
 		updated_at int NOT NULL;DEFAULT CURRENT_TIMESTAMP,
 		deleted_at int DEFAULT NULL,
 		status_at tinyint NOT NULL;DEFAULT '1'
	 );



	 CREATE TABLE IF NOT EXISTS hs_auth_records (
 		id int PRIMARY key NOT NULL,
 		secret_key varchar NOT NULL,
 		app_key varchar NOT NULL,
 		sign varchar NOT NULL DEFAULT '',
 		token varchar NOT NULL,
 		alg varchar NOT NULL,
 		ip varchar NOT NULL DEFAULT '',
 		exp varchar DEFAULT NULL,
 		iat varchar DEFAULT NULL,
 		type int NOT NULL DEFAULT '0',
 		created_at varchar NOT NULL,
 		updated_at varchar NOT NULL,
 		deleted_at varchar DEFAULT NULL,
 		status_at int NOT NULL DEFAULT '1'
	 );



	 CREATE TABLE IF NOT EXISTS hs_migrations (
 		id int PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT,
 		migration varchar NOT NULL,
 		batch int NOT NULL
	 );
