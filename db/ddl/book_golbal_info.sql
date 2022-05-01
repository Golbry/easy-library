CREATE TABLE `book_global_info`
(
    `id`                   bigint unsigned    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `book_id`              varchar(127)       NOT NULL DEFAULT '' COMMENT '书籍ID',
    `book_name`            varchar(256)       NOT NULL DEFAULT '' COMMENT '书籍名称',
    `author`               varchar(256)       NOT NULL DEFAULT '' COMMENT '作者',
    `press`                varchar(256)       NOT NULL DEFAULT '' COMMENT '出版社',
    `status`               tinyint            NOT NULL DEFAULT 0 COMMENT '当前状态,0-有效，1-已删除',
    `book_sum`             int unsigned       NOT NULL DEFAULT 0 COMMENT '书籍总数',
    `book_current_num`     int unsigned       NOT NULL DEFAULT 0 COMMENT '当前剩余数量',
    `ctime`                timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime`                timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_book_id`(`book_id`),
    KEY `idx_book_name` (`book_name`),
    KEY `idx_author`(`author`),
    KEY `idx_ctime` (`ctime`),
    KEY `idx_utime` (`utime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='书籍信息全局表';