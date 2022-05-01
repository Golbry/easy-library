CREATE TABLE `user_info`
(
    `id`                   bigint unsigned    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `user_name`            varchar(256)       NOT NULL DEFAULT '' COMMENT '小程序appid',
    `password`             varchar(256)       NOT NULL DEFAULT '' COMMENT '用户密码',
    `status`               tinyint            NOT NULL DEFAULT 0 COMMENT '当前状态,0-有效，1-已删除',
    `level`                tinyint            NOT NULL DEFAULT 0 COMMENT '用户等级，0-普通用户，1-管理员，2-超级管理员',
    `ctime`                timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime`                timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_idx_user_name` (`user_name`),
    KEY `idx_ctime` (`ctime`),
    KEY `idx_utime` (`utime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户信息表';