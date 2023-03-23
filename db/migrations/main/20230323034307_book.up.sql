BEGIN;

CREATE TABLE IF NOT EXISTS main.books (
    "id" BIGINT NOT NULL PRIMARY KEY,
    "isbn" VARCHAR NOT NULL,
    "title" VARCHAR NOT NULL,
    "author_id" BIGINT NOT NULL
        CONSTRAINT books_atuhor_id_foreign REFERENCES main.authors(id) ON UPDATE CASCADE ON DELETE CASCADE,
    "publisher_id" BIGINT NOT NULL
        CONSTRAINT books_publisher_id_foreign REFERENCES main.publishers(id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

COMMIT;