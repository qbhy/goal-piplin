CREATE TABLE IF NOT EXISTS `groups`
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    name       varchar(20) not null,
    creator_id int         not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`),
    index (creator_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;