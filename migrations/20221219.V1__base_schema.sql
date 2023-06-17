create table public.users
(
    id         text not null
        primary key,
    username   text,
    password   text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table public.users
    owner to postgres;

create index idx_users_deleted_at
    on public.users (deleted_at);
create table public.categories
(
    id         text not null
        primary key,
    category   text,
    image_url  text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table public.categories
    owner to postgres;

create index idx_categories_deleted_at
    on public.categories (deleted_at);

create table public.products
(
    id             text not null
        primary key,
    price          bigint,
    weight         bigint,
    category_id    text
        constraint fk_products_category
            references public.categories,
    per_unit       bigint,
    description    text,
    image_url      text,
    name           text,
    is_recommended boolean,
    is_top         boolean,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone
);

alter table public.products
    owner to postgres;

create index idx_products_deleted_at
    on public.products (deleted_at);

create table public.wishlists
(
    id         text not null
        primary key,
    name       text,
    price      bigint,
    weight     bigint,
    category   text,
    per_unit   bigint,
    image_url  text,
    product_id text,
    user_id    text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table public.wishlists
    owner to postgres;

create index idx_wishlists_deleted_at
    on public.wishlists (deleted_at);

create table public.carts
(
    id         text not null
        primary key,
    name       text,
    price      bigint,
    weight     bigint,
    category   text,
    per_unit   bigint,
    total      bigint,
    image_url  text,
    product_id text,
    user_id    text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table public.carts
    owner to postgres;

create index idx_carts_deleted_at
    on public.carts (deleted_at);

create table public.transactions
(
    id          text not null
        primary key,
    user_id     text,
    total_price bigint,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
);

alter table public.transactions
    owner to postgres;

create index idx_transactions_deleted_at
    on public.transactions (deleted_at);

create table public.transaction_details
(
    id             text not null
        primary key,
    product_id     text
        constraint fk_transaction_details_product
            references public.products,
    transaction_id text
        constraint fk_transactions_transaction_details
            references public.transactions,
    price          bigint,
    weight         bigint,
    category       text,
    per_unit       bigint,
    description    text,
    image_url      text,
    name           text,
    category_id    text,
    total          bigint,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone
);

alter table public.transaction_details
    owner to postgres;

create index idx_transaction_details_deleted_at
    on public.transaction_details (deleted_at);

create table public.roles
(
    id   uuid not null
        primary key,
    name varchar(255)
);

alter table public.roles
    owner to postgres;

create table public.users_roles
(
    user_id  varchar(255) not null
        constraint fk2o0jvgh89lemvvo17cbqvdxaa
            references public.users,
    roles_id uuid         not null
        constraint fka62j07k5mhgifpp955h37ponj
            references public.roles
);

alter table public.users_roles
    owner to postgres;

