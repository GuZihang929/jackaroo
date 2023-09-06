CREATE TABLE `company` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                           `company` varchar(255)  NOT NULL DEFAULT '' COMMENT '公司名称',
                           `number`    bigint   NOT NULL DEFAULT 0,
                           `brief` varchar(255)  NOT NULL DEFAULT '' COMMENT '简介',
                           `picture` varchar(255)  NOT NULL DEFAULT '' COMMENT '图片',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;