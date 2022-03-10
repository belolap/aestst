-- Реализация для PostgreSQL.
BEGIN;

-- Автор (важно хранить ФИО).
CREATE TABLE "authors" (
    "id" SERIAL PRIMARY KEY,
    -- Можно не разбивать, сделать просто name. Сделал разбивку,
    -- чтобы облегчить в дальнейшем задачи сортировки по фамилии.
    "first_name" VARCHAR(128) NOT NULL,
    "middle_name" VARCHAR(128),
    "last_name" VARCHAR(128) NOT NULL
);

-- Книга (важно хранить название).
CREATE TABLE "books" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR NOT NULL
);

-- Соответствие авторов книгам (много ко многим).
CREATE TABLE "book_authors" (
    "id" SERIAL PRIMARY KEY,
    "book_id" INTEGER NOT NULL REFERENCES "books" ("id"),
    "author_id" INTEGER NOT NULL REFERENCES "authors" ("id")
);

-- Издание (важно хранить ISBN).
CREATE TABLE "publications" (
    "id" SERIAL PRIMARY KEY,
    "book_id" INTEGER NOT NULL REFERENCES "books" ("id"),
    -- ISBN имеет формат максимум в 13 символов, сделал больше, чтобы не париться, если формат поменяется потом.
    "isbn" VARCHAR(32) NOT NULL
);

-- Гонорар (дату, сумму, кому выплачен).
CREATE TABLE "fees" (
    "id" SERIAL PRIMARY KEY,
    "date" DATE NOT NULL,
    "publication_id" INTEGER NOT NULL REFERENCES "publications" ("id"),
    "author_id" INTEGER NOT NULL REFERENCES "authors" ("id"),
    -- Допустим, что деньги храним как целые числа.
    "fee" INTEGER NOT NULL
);

END;
