CREATE TABLE IF NOT EXISTS share_files
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    project_id int          not null,
    name       varchar(20)  not null,
    path       varchar(200) not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`),
    index (project_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;