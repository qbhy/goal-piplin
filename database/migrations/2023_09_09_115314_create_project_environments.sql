CREATE TABLE IF NOT EXISTS project_environments
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    project_id int         not null,
    name       varchar(40) not null,
    settings   json        not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`),
    index (project_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;