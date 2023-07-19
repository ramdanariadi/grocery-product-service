create table shops
(
    id         varchar(36)  not null
        primary key,
    name       varchar(100) not null,
    address    varchar(300),
    user_id    varchar(36)
        constraint fk_shops_user
            references users,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table shops
    owner to postgres;

create index idx_shops_deleted_at
    on shops (deleted_at);

create index idx_shops_user_id
    on shops (user_id);

