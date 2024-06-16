create table `t_account` (
    `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id` bigint(11) NOT NULL COMMENT '用户ID',
    `username` varchar(128) NOT NULL default '' COMMENT '用户名',
    `password` varchar(255) NOT NULL default '' COMMENT '密码',
    `phone` varchar(32) NOT NULL default '' COMMENT '手机号',
    `email` varchar(64) NOT NULL default '' COMMENT '邮箱',
    `active` tinyint(1) NOT NULL default 1 COMMENT '状态 1:正常 0:禁用',
    `created_at` datetime NOT NULL default CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL default CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime default NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_email` (`email`),
    UNIQUE KEY `idx_phone` (`phone`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `udx_deleted_time_email` (`deleted_at`, `email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;