create table if not exists users (
    id serial primary key,
    username text unique not null,
    passphrase bytea not null
);
