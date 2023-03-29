BEGIN;

CREATE SEQUENCE publisher_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS main.publishers (
    "id" BIGINT NOT NULL PRIMARY KEY DEFAULT nextval('publisher_id_seq'::regclass),
    "uuid" UUID NOT NULL,
    "name" VARCHAR NOT NULL,
    "kota" VARCHAR NOT NULL
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- INSERT INTO main.publishers (id, name, kota) VALUES(1, 'Gramedia', 'Jakarta');
-- INSERT INTO main.publishers (id, name, kota) VALUES(2, 'Elex Media', 'Jakarta');
-- INSERT INTO main.publishers (id, name, kota) VALUES(3, 'Bentang Pustaka', 'Bandung');
-- INSERT INTO main.publishers (id, name, kota) VALUES(4, 'Pustaka Sinar Harapan', 'Bandung');
-- INSERT INTO main.publishers (id, name, kota) VALUES(5, 'Pustaka Pelajar', 'Jakarta');
-- INSERT INTO main.publishers (id, name, kota) VALUES(6, 'Pustaka Alfabeta', 'Jakarta');
-- INSERT INTO main.publishers (id, name, kota) VALUES(7, 'Pustaka Sinar Harapan', 'Bandung');

ALTER SEQUENCE publisher_id_seq OWNED BY main.publishers.id;

COMMIT;