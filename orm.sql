/*
 Navicat Premium Data Transfer

 Source Server         : mysql_localhost
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost:3306
 Source Schema         : orm

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 18/12/2017 11:08:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userid` bigint(20) NOT NULL,
  `name` varchar(16) NOT NULL,
  `cash` float NOT NULL,
  `gender` bit(1) NOT NULL,
  `card` enum('123123','123123432','ewrwer','234wer') NOT NULL,
  `create_time` date NOT NULL,
  `payment` double DEFAULT NULL,
  `address` json DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'derek', 3424, b'1', 'ewrwer', '2017-12-18', 324234, '{\"city\": \"shanghai\"}');
INSERT INTO `user` VALUES (2, 'pucca', 34324, b'1', 'ewrwer', '2017-12-18', 3244, '{\"city\": \"beijing\"}');
INSERT INTO `user` VALUES (3, 'frank', 343.33, b'1', 'ewrwer', '2017-12-18', 32.44, '{\"city\": \"zhejiang\"}');
INSERT INTO `user` VALUES (4, 'slinda', 34333.3, b'1', 'ewrwer', '2017-12-18', 3222.44, '{\"city\": \"guangzhou\"}');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
