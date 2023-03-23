BEGIN;

CREATE TABLE IF NOT EXISTS main.publishers (
    "id" BIGINT NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "kota" VARCHAR NOT NULL,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

COMMIT;