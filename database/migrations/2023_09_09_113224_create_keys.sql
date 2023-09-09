CREATE TABLE IF NOT EXISTS `keys`
(
    `id`        INT UNSIGNED AUTO_INCREMENT,
    name        varchar(20) not null,
    public_key  text        not null,
    private_key mediumtext  not null,
    fingerprint varchar(64) not null,
    created_at  timestamp,
    updated_at  timestamp,
    PRIMARY KEY (`id`),
    unique index (fingerprint)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;