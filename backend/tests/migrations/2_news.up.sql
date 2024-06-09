create table news
(
    id       serial
        primary key,
    title    text      not null,
    text     text      not null,
    datetime timestamp not null
);

alter table news
    owner to hack;

grant delete, insert, select, update on news to hack;

create table categories
(
    id   serial
        primary key,
    name text not null
);

alter table categories
    owner to hack;

grant delete, insert, select, update on categories to hack;

create table "categoriesNews"
(
    id           serial
        primary key,
    "newsID"     integer not null
        references news,
    "categoryID" integer not null
        references categories
);

alter table "categoriesNews"
    owner to hack;

grant delete, insert, select, update on "categoriesNews" to hack;

