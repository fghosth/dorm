ALTER TABLE `address` DROP FOREIGN KEY `fk_address_user_1`;

DROP INDEX `dd` ON `user`;
DROP INDEX `city` ON `address`;

DROP TABLE `user`;
DROP TABLE `address`;

CREATE TABLE `user_dd` (
`uid` int(11) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '自增id号',
`name` varchar(64) NOT NULL DEFAULT 'aaa',
`gender` bit(1) NOT NULL,
`password` varchar(255) NOT NULL,
`qq` bit(8) NOT NULL,
`account` char(255) NOT NULL,
`cellphone` varchar(255) NOT NULL,
`happy` enum('aaa','bbb','ccc','ddd') NOT NULL DEFAULT 'aaa',
`cash` double NOT NULL DEFAULT 0,
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`update_time` datetime NOT NULL,
PRIMARY KEY (`uid`,`qq`)
INDEX `dd` (`account` ASC) USING HASH
 UNIQUE KEY `auth_api_api_key_unique` (`uid`),
UNIQUE KEY `rrr` (`name`) USING BTREE,
  KEY `rrrrrrr` (`create_time`),
  KEY `rrrrrrdddr` (`update_time`)
);

CREATE TABLE `address` (
`aid` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`city` varchar(255) NOT NULL,
`privent` varchar(255) NULL,
`address` varchar(255) NOT NULL,
`uid` int(11) UNSIGNED NOT NULL,
PRIMARY KEY (`aid`) ,
INDEX `city` (`city` ASC) USING BTREE
);


ALTER TABLE `address` ADD CONSTRAINT `fk_address_user_1` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`);
