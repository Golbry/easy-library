CREATE TABLE `book_history`
(
    `id`                   bigint unsigned    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `book_id`              varchar(127)       NOT NULL DEFAULT '' COMMENT '书籍ID',
    `book_name`            varchar(256)       NOT NULL DEFAULT '' COMMENT '书籍名称',
    `status`               tinyint            NOT NULL DEFAULT 0 COMMENT '当前状态，0-有效，1-删除',
    `borrower`             varchar(256)       NOT NULL DEFAULT '' COMMENT '借阅人',
    `borrow_num`           int unsigned       NOT NULL DEFAULT 0 COMMENT '借阅数量',
    `ctime`                timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime`                timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    FOREIGN KEY `for_idx_book_id`(`book_id`) REFERENCES `book_global_info` (`book_id`),
    KEY `idx_book_name` (`book_name`),
    KEY `idx_borrower`(`borrower`),
    KEY `idx_ctime` (`ctime`),
    KEY `idx_utime` (`utime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='书籍借阅历史表';