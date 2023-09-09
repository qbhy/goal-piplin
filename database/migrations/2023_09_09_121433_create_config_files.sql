CREATE TABLE IF NOT EXISTS config_files
(
    `id`         INT UNSIGNED AUTO_INCREMENT,
    project_id   int          not null,
    name         varchar(64)  not null,
    path         varchar(200) not null,
    content      mediumtext   not null,
    environments text         not null,
    created_at   timestamp,
    updated_at   timestamp,
    PRIMARY KEY (`id`),
    index (project_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;