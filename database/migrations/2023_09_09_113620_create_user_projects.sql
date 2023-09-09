CREATE TABLE IF NOT EXISTS user_projects
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    project_id int         not null,
    user_id    int         not null,
    status     varchar(20) not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`),
    index (user_id),
    index (project_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;