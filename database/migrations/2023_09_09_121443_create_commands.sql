CREATE TABLE IF NOT EXISTS commands
(
    `id`            INT UNSIGNED AUTO_INCREMENT,
    project_id      int         not null,
    step            varchar(20) not null,
    sort            int         not null,
    name            varchar(64) not null,
    user            varchar(20) not null,
    script          mediumtext  not null,
    environments    text        not null,
    optional        bool        not null,
    defaultSelected bool        not null,
    created_at      timestamp,
    updated_at      timestamp,
    PRIMARY KEY (`id`),
    index (project_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;