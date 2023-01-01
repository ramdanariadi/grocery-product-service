create table category
(
    id         uuid not null
        primary key,
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp,
    category   varchar(255),
    image_url  varchar(255)
);

alter table category
    owner to postgres;

create table products
(
    id             uuid not null
        primary key,
    created_at     timestamp,
    deleted_at     timestamp,
    updated_at     timestamp,
    description    varchar(255),
    image_url      varchar(255),
    is_recommended boolean,
    is_top         boolean,
    name           varchar(255),
    per_unit       integer,
    price          bigint,
    weight         integer,
    category_id    uuid
        constraint fk1cf90etcu98x1e6n9aks3tel3
            references category
);

alter table products
    owner to postgres;

create table cart
(
    id         uuid         not null
        primary key,
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp,
    category   varchar(255),
    image_url  varchar(255),
    name       varchar(255),
    per_unit   integer,
    price      bigint,
    total      integer,
    weight     integer,
    product_id uuid
        constraint fkpu4bcbluhsxagirmbdn7dilm5
            references products,
    user_id    varchar(255) not null
);

alter table cart
    owner to postgres;

create table liked
(
    id         uuid         not null
        primary key,
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp,
    category   varchar(255),
    image_url  varchar(255),
    name       varchar(255),
    per_unit   integer,
    price      bigint,
    weight     integer,
    product_id uuid
        constraint fk4tgscg55hv7our1fju6yu7o2q
            references products,
    user_id    varchar(255) not null
);

alter table liked
    owner to postgres;

create index idx_product_name
    on products (name);

create index idx_product_category
    on products (category_id);

create table transaction
(
    id          uuid not null
        primary key,
    created_at  timestamp,
    deleted_at  timestamp,
    updated_at  timestamp,
    total_price bigint,
    user_id     varchar(255)
);

alter table transaction
    owner to postgres;

create table detail_transaction
(
    id             uuid not null
        primary key,
    created_at     timestamp,
    deleted_at     timestamp,
    updated_at     timestamp,
    image_url      varchar(255),
    name           varchar(255),
    per_unit       integer,
    price          bigint,
    total          integer,
    weight         integer,
    product_id     uuid
        constraint fkf784xfktx9aaei9fyqxfrxdqj
            references products,
    transaction_id uuid
        constraint fkd6mv4qpm8jsp3lviwupuodmm4
            references transaction
);

alter table detail_transaction
    owner to postgres;

create table wishlist
(
    id         uuid         not null
        primary key,
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp,
    category   varchar(255),
    image_url  varchar(255),
    name       varchar(255),
    per_unit   integer,
    price      bigint,
    weight     integer,
    product_id uuid
        constraint fk6p7qhvy1bfkri13u29x6pu8au
            references products,
    user_id    varchar(255) not null
);

alter table wishlist
    owner to postgres;

