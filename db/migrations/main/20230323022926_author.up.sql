BEGIN;

CREATE SEQUENCE author_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS main.authors (
    "id" BIGINT NOT NULL PRIMARY KEY DEFAULT nextval('author_id_seq'::regclass),
    "uuid" UUID NOT NULL,
    "name" VARCHAR NOT NULL,
    "gender" VARCHAR NOT NULL
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- INSERT INTO main.authors (id, name, gender) VALUES(1, 'J.K. Rowling', 'F');
-- INSERT INTO main.authors (id, name, gender) VALUES(2, 'J.R.R. Tolkien', 'Male');
-- INSERT INTO main.authors (id, name, gender) VALUES(3, 'Agatha Christie', 'Female');
-- INSERT INTO main.authors (id, name, gender) VALUES(4, 'Stephen King', 'Male');
-- INSERT INTO main.authors (id, name, gender) VALUES(5, 'John Grisham', 'Male');
-- INSERT INTO main.authors (id, name, gender) VALUES(6, 'Danielle Steel', 'Female');
-- INSERT INTO main.authors (id, name, gender) VALUES(7, 'Dan Brown', 'Male');
-- INSERT INTO main.authors (id, name, gender) VALUES(8, 'J.D. Salinger', 'Male');
-- INSERT INTO main.authors (id, name, gender) VALUES(9, 'J.R.R. Martin', 'Male');
-- INSERT INTO main.authors (id, name, gender) VALUES(10, 'George R.R. Martin', 'Male');

ALTER SEQUENCE author_id_seq OWNED BY main.authors.id;

COMMIT;