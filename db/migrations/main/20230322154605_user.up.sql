BEGIN;

CREATE TABLE IF NOT EXISTS main.users
(
    id                    BIGINT      NOT NULL,
    uuid                  UUID         NOT NULL,
    name                  VARCHAR(128) NOT NULL,
    email                 VARCHAR(128) NOT NULL,
    password              VARCHAR(255) NOT NULL,
    created_by            VARCHAR(128) NOT NULL,
    updated_by            VARCHAR(128) NOT NULL,
    deleted_by            VARCHAR(128),
    created_at            TIMESTAMPTZ  NOT NULL,
    updated_at            TIMESTAMPTZ  NOT NULL,
    deleted_at            TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;