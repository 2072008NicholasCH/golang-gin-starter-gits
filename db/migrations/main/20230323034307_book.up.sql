BEGIN;

CREATE SEQUENCE book_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS main.books (
    "id" BIGINT NOT NULL PRIMARY KEY DEFAULT nextval('book_id_seq'::regclass),
    "uuid" UUID NOT NULL,
    "isbn" VARCHAR NOT NULL,
    "title" VARCHAR NOT NULL,
    "author_id" BIGINT NOT NULL
        CONSTRAINT books_atuhor_id_foreign REFERENCES main.authors(id) ON UPDATE CASCADE ON DELETE CASCADE,
    "publisher_id" BIGINT NOT NULL
        CONSTRAINT books_publisher_id_foreign REFERENCES main.publishers(id) ON UPDATE CASCADE ON DELETE CASCADE
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(1, '9786020302717', 'Harry Potter and the Philosopher''s Stone', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(2, '9786020302724', 'Harry Potter and the Chamber of Secrets', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(3, '9786020302731', 'Harry Potter and the Prisoner of Azkaban', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(4, '9786020302748', 'Harry Potter and the Goblet of Fire', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(5, '9786020302755', 'Harry Potter and the Order of the Phoenix', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(6, '9786020302762', 'Harry Potter and the Half-Blood Prince', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(7, '9786020302779', 'Harry Potter and the Deathly Hallows', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(8, '9786020302786', 'Harry Potter and the Cursed Child', 1, 1);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(9, '8580001058174', 'The Hobbit', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(10, '9786020302793', 'The Lord of the Rings', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(11, '9786020302809', 'The Fellowship of the Ring', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(12, '9786020302816', 'The Two Towers', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(13, '9786020302823', 'The Return of the King', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(14, '9786020302830', 'The Silmarillion', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(15, '9786020302847', 'The Children of Hurin', 2, 2);
-- INSERT INTO main.books (id, isbn, title, author_id, publisher_id) VALUES(16, '1234567890123', 'And Then There Were None', 3, 3);

ALTER SEQUENCE book_id_seq OWNED BY main.books.id;

COMMIT;