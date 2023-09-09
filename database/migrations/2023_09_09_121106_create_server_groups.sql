CREATE TABLE IF NOT EXISTS server_groups
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    name       varchar(20) not null,
    settings   json        not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`),
    unique index (name)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;