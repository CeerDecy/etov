CREATE TABLE USERS
(
    `id`         INTEGER PRIMARY KEY AUTO_INCREMENT,
    `nickname`   varchar(255) DEFAULT '' COMMENT '昵称',
    `email`      varchar(255) DEFAULT '' COMMENT '邮箱',
    `phone`      char(11)     DEFAULT '' COMMENT '手机号',
    `password`   char(255)     DEFAULT '' COMMENT '密码',
    `avatar`     char(255)     DEFAULT '' COMMENT '头像地址',
    `api_key`    varchar(255) DEFAULT '' COMMENT 'OpenAI的key',
    `salt`       char(255)      DEFAULT '' COMMENT '盐值',
    `validate`   char(1)      DEFAULT 'N' COMMENT '是否通过学生验证',
    `created_at` datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime     DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`)
);