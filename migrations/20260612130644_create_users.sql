-- +goose Up

create table users
(
    id            uuid primary key,
    email         text        not null unique,
    password_hash text        not null,
    is_active     boolean     not null default true,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now()
);

create index idx_users_email on users (email);

-- +goose Down


drop table users;
