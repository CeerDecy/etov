CREATE TABLE CHATS
(
    `id`         INTEGER PRIMARY KEY AUTO_INCREMENT,
    `uid`        INTEGER not null,
    `title`      varchar(255) DEFAULT '' COMMENT '会话标题',
    `created_at` datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime     DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`),
    INDEX `index_uid` (`uid`)
);