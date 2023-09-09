CREATE TABLE IF NOT EXISTS deployments
(
    `id`         INT UNSIGNED AUTO_INCREMENT,
    project_id   int          not null,
    version      varchar(64)  not null,
    comment      varchar(200) not null,
    environments text         not null,
    created_at   timestamp,
    updated_at   timestamp,
    PRIMARY KEY (`id`),
    index (project_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;