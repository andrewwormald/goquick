create table users (
    id bigint not null,
    name varchar(255)not null,
);

create table transactions (
    id bigint auto increment primary,
    type int,
);
