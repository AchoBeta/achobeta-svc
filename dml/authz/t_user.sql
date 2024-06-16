create table `t_user`(
    `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `nickname` varchar(32) NOT NULL default '' COMMENT '昵称',
    `gender` tinyint(1) NOT NULL default 0 COMMENT '性别 0:未知 1:男 2:女',
    `avatar` varchar(255) NOT NULL default '' COMMENT '头像 link',
    `introduction` text COMMENT '个人简介',
    `user_version` int(11) NOT NULL default 0 COMMENT '用户版本',
    `created_at` datetime NOT NULL default CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL default CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime default NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;