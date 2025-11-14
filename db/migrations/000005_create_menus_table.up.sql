create extension if not exists "pgcrypto";

create table
    menus (
        id uuid primary key default gen_random_uuid (),
        menu_id uuid references menus (id) on delete cascade,
        name varchar(100) not null,
        depth int default 0,
        sort_order int default 0,
        created_at timestamp not null default current_timestamp,
        updated_at timestamp default current_timestamp
    );

create index idx_menus_menu_id on menus (menu_id);

create index idx_menus_sort_order on menus (sort_order);

create index idx_menus_depth on menus (depth);