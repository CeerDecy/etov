CREATE TABLE TOOL
(
    `id`          INTEGER PRIMARY KEY AUTO_INCREMENT,
    `name`        VARCHAR(255) NOT NULL COMMENT '工具名称',
    `logo`        VARCHAR(255) NOT NULL COMMENT '工具logo',
    `url`         VARCHAR(255) NOT NULL COMMENT '工具链接',
    `description` TEXT         NOT NULL COMMENT '工具描述',
    `is_public`   CHAR(1)      NOT NULL DEFAULT 'N' COMMENT '是否公开',
    `disabled`    CHAR(1)      NOT NULL DEFAULT 'N' COMMENT '是否不可用',
    `author_id`   INTEGER      NOT NULL COMMENT '作者ID',
    `created_at`  datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime              DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`)
);

INSERT INTO TOOL (name, logo, url, description, is_public, author_id, created_at, updated_at, deleted_at)
VALUES ('ChatGPT3.5',
        '/api/static/chat3.5.png',
        '/tools/chat',
        'ChatGPT3.5是一种基于深度学习的聊天机器人。 它使用的是GPT-3（Generative Pre-trained Transformer 3）模型的改进版。 GPT-3是OpenAI研究团队开发的一个自然语言生成模型，可以生成高质量的自然语言文本。',
        'Y',
        1,
        now(),
        now(),
        null);

INSERT INTO TOOL (name, logo, url, description, is_public,disabled, author_id, created_at, updated_at, deleted_at)
VALUES ('ChatGPT4.0',
        '/api/static/chat4.0.png',
        '/tools/chat',
        'ChatGPT 4.0是一种自然语言处理模型，旨在让机器理解人类语言，它是由OpenAI开发的，并且在多个自然语言处理任务中的表现都比之前的版本更好，具有更强的适应性和通用性。',
        'Y',
        'Y',
        1,
        now(),
        now(),
        null);

INSERT INTO TOOL (name, logo, url, description, is_public,disabled, author_id, created_at, updated_at, deleted_at)
VALUES ('AI论文降重',
        '/api/static/browser01.png',
        '/tools/reduce-duplication',
        'AI论文降重（AI Paper Paraphrasing）是一种利用人工智能技术来对学术论文进行重写和改写的方法。它旨在帮助研究人员、学生和作者避免抄袭，并提供更多的原创性。
通过使用自然语言处理（NLP）和机器学习技术，AI论文降重系统可以分析输入的论文文本，并生成与原始文本意思相同但表达方式不同的新文本。这种重写过程通常包括更改句子结构、替换同义词、重新排列段落等操作，以确保新文本的独特性和可读性',
        'Y',
        'N',
        1,
        now(),
        now(),
        null);

INSERT INTO TOOL (name, logo, url, description, is_public,disabled, author_id, created_at, updated_at, deleted_at)
VALUES ('AI翻译',
        '/api/static/translate.png',
        '/tools/chat',
        'AI翻译（AI Translation）是一种利用人工智能技术来实现自动翻译的方法。它使用机器学习和自然语言处理技术，将一种语言的文本或口语转换为另一种语言，以实现跨语言交流和理解。
AI翻译系统通过训练大量的语言数据和模型，可以自动识别和理解不同语言之间的语义和语法规则。它可以处理多种语言对，例如英语到中文、中文到法语等，并提供快速、准确的翻译结果。',
        'Y',
        'Y',
        1,
        now(),
        now(),
        null);

INSERT INTO TOOL (name, logo, url, description, is_public,disabled, author_id, created_at, updated_at, deleted_at)
VALUES ('段落总结',
        '/api/static/book.png',
        '/tools/chat',
        'AI段落总结是一种利用人工智能技术来自动提取和生成段落摘要的方法。它旨在帮助人们快速理解和获取文本段落的关键信息，节省阅读时间和提高工作效率。
通过使用自然语言处理（NLP）和机器学习技术，AI段落总结系统可以分析输入的段落文本，并从中提取出最重要的句子或短语，形成一个简洁而准确的摘要。这种摘要通常包括段落的主题、关键观点和重要细节，以便读者能够快速了解段落的内容。',
        'Y',
        'Y',
        1,
        now(),
        now(),
        null);

INSERT INTO TOOL (name, logo, url, description, is_public,disabled, author_id, created_at, updated_at, deleted_at)
VALUES ('AI写作',
        '/api/static/certificate.png',
        '/tools/chat',
        'AI写作是一种利用人工智能技术来生成和创作文本的方法。它使用自然语言处理（NLP）和机器学习技术，使计算机能够模拟人类的写作能力，并生成高质量、流畅的文本内容。
AI写作系统可以通过学习大量的文本数据和模型训练，理解语言的语法、语义和上下文，并生成与人类写作相似的文章、故事、新闻、评论等。这些系统可以根据给定的主题、风格和要求，自动生成具有逻辑性、连贯性和创造性的文本。',
        'Y',
        'Y',
        1,
        now(),
        now(),
        null);