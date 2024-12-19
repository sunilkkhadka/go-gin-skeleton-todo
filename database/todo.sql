CREATE TABLE IF NOT EXISTS `todos`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `title`       VARCHAR(100)                NOT NULL,
    `description` TEXT                         NULL,
    `status`      ENUM('pending', 'completed', 'in-progress') NOT NULL DEFAULT 'pending',
    `created_at`  DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  DATETIME                    NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
