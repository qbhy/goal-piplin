CREATE TABLE IF NOT EXISTS users
(
    id         INT UNSIGNED AUTO_INCREMENT,
    username   varchar(20) not null,
    nickname   varchar(20) not null,
    avatar     varchar(255),
    role       varchar(20) not null,
    password   varchar(64) not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;