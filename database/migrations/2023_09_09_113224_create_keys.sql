CREATE TABLE IF NOT EXISTS `keys`
(
    `id`        INT UNSIGNED AUTO_INCREMENT,
    name        varchar(20) not null,
    public_key  text        not null,
    private_key mediumtext  not null,
    created_at  timestamp,
    updated_at  timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;