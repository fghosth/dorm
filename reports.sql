

CREATE TABLE `reports` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `mp_ID` bigint(20) unsigned NOT NULL COMMENT '物料id或成本id',
                           `category` tinyint(4) NOT NULL COMMENT '1:物料 2:成品',
                           `cid` bigint(20) unsigned NOT NULL COMMENT '成本中心id',
                           `scost` bigint(20) unsigned NOT NULL COMMENT '期初成本',
                           `stime` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                           `etime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           `is_deleted` tinyint(1) NOT NULL DEFAULT 0,
                           `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           `modify_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           `ecost` bigint(20) unsigned NOT NULL COMMENT '期末成本',
                           `s_QTY` double NOT NULL COMMENT '期初数量',
                           `e_QTY` double NOT NULL COMMENT '期末成本',
                           `recipt_QTY` double NOT NULL COMMENT '单据累加总和',
                           `increase_QTY` double NOT NULL COMMENT '增加 数量',
                           `decrease_QTY` double NOT NULL COMMENT '较少数量',
                           `bom` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'bom 物料详情' CHECK (json_valid(`bom`)),
                           `check_increase` double NOT NULL COMMENT '盘点增加数量',
                           `check_decrease` double NOT NULL COMMENT '盘点 减少数量',
                           `scrap_qty` double NOT NULL COMMENT '保费数量',
                           `sold_qty` double NOT NULL COMMENT '销售数量',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

