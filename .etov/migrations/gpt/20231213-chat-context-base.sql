CREATE TABLE CHAT_CONTEXT
(
    `id`         INTEGER PRIMARY KEY AUTO_INCREMENT,
    `chat_id`    INTEGER NOT NULL COMMENT '会话ID',
    `content`    TEXT COMMENT '正文',
    `index`      INTEGER COMMENT '会话索引',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`),
    INDEX `index_chat_id` (`chat_id`)
)