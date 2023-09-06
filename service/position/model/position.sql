CREATE TABLE `position` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                            `cid` bigint  NOT NULL DEFAULT 0,
                            `company` varchar(255)  NOT NULL DEFAULT '' ,
                            `title` varchar(255)  NOT NULL DEFAULT '' ,
                            `job_category` varchar(255)  NOT NULL DEFAULT '0' ,
                            `job_type_name` varchar(255)  NOT NULL DEFAULT '0' ,
                            `job_detail` longtext ,
                            `job_location` varchar(255)  NOT NULL DEFAULT '0',
                            `push_time`   varchar(255)  NOT NULL DEFAULT '0',
                            `fetch_time`  varchar(255)  NOT NULL DEFAULT '0',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;