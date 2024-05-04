CREATE TABLE IF NOT EXISTS user_groups
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    group_id   int         not null,
    user_id    int         not null,
    status     varchar(20) not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`),
    index (user_id),
    index (group_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;