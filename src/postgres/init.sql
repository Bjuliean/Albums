CREATE TABLE albums
(
    id SERIAL PRIMARY KEY,
    title Varchar(255) NOT NULL,
    artist Varchar(255) NOT NULL,
    price Decimal
);

INSERT INTO albums(id, title, artist, price) VALUES (0, 'abobatitle', 'abobaauthor', 255.55);
INSERT INTO albums(id, title, artist, price) VALUES (1, 'Blue Train', 'John Coltrane', 56.99);
INSERT INTO albums(id, title, artist, price) VALUES (2, 'Jeru', 'Gerry Mulligan', 17.99);
INSERT INTO albums(id, title, artist, price) VALUES (3, 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);