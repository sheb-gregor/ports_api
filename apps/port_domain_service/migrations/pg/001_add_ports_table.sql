-- +migrate Up
CREATE TABLE IF NOT EXISTS ports
(
    unlocode   VARCHAR(5) UNIQUE PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    city       VARCHAR(255) NOT NULL,
    country    VARCHAR(255) NOT NULL,
    timezone   VARCHAR(255) NOT NULL,
    code       VARCHAR(255) NOT NULL,
    extra      JSONB  DEFAULT '{}',
    created_at BIGINT DEFAULT date_part('epoch', timezone('utc', now())),
    updated_at BIGINT DEFAULT date_part('epoch', timezone('utc', now()))
);


-- +migrate Down
DROP TABLE IF EXISTS ports;
