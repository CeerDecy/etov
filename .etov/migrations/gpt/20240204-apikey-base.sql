CREATE TABLE APIKEY
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `uid`        INTEGER      NOT NULL COMMENT '用户ID',
    `key_name`   VARCHAR(255) NOT NULL COMMENT 'KEY名称',
    `apikey`     VARCHAR(255) NOT NULL COMMENT 'key值',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`),
    INDEX `index_uid` (`uid`)
);

INSERT INTO  APIKEY (uid, key_name, apikey,created_at,updated_at)VALUES (1,'ChatGPT3.5','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',now(),now());
INSERT INTO  APIKEY (uid, key_name, apikey,created_at,updated_at)VALUES (1,'ChatGPT4.0','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',now(),now());
