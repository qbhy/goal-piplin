CREATE TABLE IF NOT EXISTS projects
(
    `id`           INT UNSIGNED AUTO_INCREMENT,
    uuid           varchar(16)  not null,
    name           varchar(20)  not null,
    creator_id     int          not null,
    key_id         int          not null,
    repo_address   varchar(200) not null,
    project_path   varchar(200) not null,
    default_branch varchar(200) not null,
    settings       json         not null,
    created_at     timestamp,
    updated_at     timestamp,
    PRIMARY KEY (`id`),
    unique index (uuid),
    index (creator_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

