create table users
(
    id         serial
        primary key,
    surname    text not null,
    name       text not null,
    patronymic text,
    age        integer,
    photourl   text,
    email      text not null,
    password   text not null,
    push       boolean
);

alter table users
    owner to hack;

grant delete, insert, select, update on users to hack;

create table roles
(
    id   serial
        primary key,
    name text not null
);

alter table roles
    owner to hack;

grant delete, insert, select, update on roles to hack;

create table userroles
(
    id     serial
        primary key,
    userid integer
        constraint userroles_users_id_fk
            references users,
    roleid integer
        constraint userroles_roles_id_fk
            references roles
);

alter table userroles
    owner to hack;

grant delete, insert, select, update on userroles to hack;
