-- +goose Up
CREATE TABLE IF NOT EXISTS `go_crm_user`
(
    `usr_id`               INT AUTO_INCREMENT COMMENT 'Account ID',

    `usr_email`            VARCHAR(255) NOT NULL COMMENT 'Email',
    `usr_phone`            VARCHAR(20)  NOT NULL COMMENT 'Phone',
    `usr_username`         VARCHAR(100) NOT NULL COMMENT 'Username',

    `usr_password`         VARCHAR(255) NOT NULL COMMENT 'Hashed Password',

    `usr_created_at`       BIGINT COMMENT 'Created Time',
    `usr_updated_at`       BIGINT COMMENT 'Updated Time',

    `usr_create_ip_at`     VARCHAR(45) COMMENT 'Created IP',
    `usr_last_login_at`    BIGINT COMMENT 'Last Login Time',
    `usr_last_login_ip_at` VARCHAR(45) COMMENT 'Last Login IP',

    `usr_login_times`      INT     DEFAULT 0 COMMENT 'Login Times',

    `usr_status`           TINYINT DEFAULT 1 COMMENT '1:active,0:inactive',

    PRIMARY KEY (`usr_id`),

    UNIQUE KEY `uk_usr_email` (`usr_email`),
    UNIQUE KEY `uk_usr_phone` (`usr_phone`),
    UNIQUE KEY `uk_usr_username` (`usr_username`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS `go_crm_user`;
