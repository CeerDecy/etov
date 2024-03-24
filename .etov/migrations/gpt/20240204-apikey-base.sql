CREATE TABLE APIKEY
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `uid`        INTEGER      NOT NULL COMMENT '用户ID',
    `key_name`   VARCHAR(255) NOT NULL COMMENT 'KEY名称',
    `apikey`     VARCHAR(255) NOT NULL COMMENT 'key值',
    `model_tag`  VARCHAR(255) NOT NULL COMMENT '模型',
    `host`       VARCHAR(255) NOT NULL COMMENT 'api地址',
    `is_deleted` INT      DEFAULT 0 COMMENT '是否被删除',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`),
    INDEX `index_uid` (`uid`)
);

INSERT INTO APIKEY (uid, key_name, apikey, model_tag, host, is_deleted, created_at, updated_at)
VALUES (1, 'ChatGPT3.5', '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', 'gpt-3.5-turbo',
        'https://api.openai.com/v1', 0, now(), now());
INSERT INTO APIKEY (uid, key_name, apikey, model_tag, host, is_deleted, created_at, updated_at)
VALUES (1, 'ChatGPT4.0', 'sk-Lskydw569JiYB3XW9bF9Fd413cB940D688C28c28Cc363e17', 'gpt-4-32k',
        'https://kapkey.chatgptapi.org.cn/v1', 0, now(), now());
