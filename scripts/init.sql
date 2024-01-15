create table if not exists suppliers(
    name varchar(255) not null primary key,
    address varchar(255),
    delivery_time interval not null
);

create table if not exists cake_decorations(
    article integer not null primary key,
    name varchar(255) not null,
    measurement_unit varchar(255) not null,
    count integer not null,
    main_supplier varchar(255),
        foreign key (main_supplier) references suppliers(name),
    image varchar(255),
    type varchar(255) not null,
    price_per_unit varchar not null,
    weight_per_unit integer not null
);

create table if not exists ingredients(
    article integer not null primary key,
    name varchar(255) not null,
    measurement_unit varchar(255) not null,
    count integer not null,
    main_supplier varchar(255),
        foreign key (main_supplier) references suppliers(name),
    image varchar(255),
    type varchar(255) not null,
    price_per_unit varchar(255) not null,
    state_standard varchar(255),
    pre_packing varchar(255),
    characteristic text
);

create table if not exists products(
    name varchar(255) not null primary key,
    size varchar(255) not null
);

create table if not exists equipment_types(
    equipment_type varchar(255) not null primary key
);

create table if not exists equipment(
    labeling varchar(255) not null primary key,
    equipment_type varchar(255) not null,
        foreign key (equipment_type) references equipment_types(equipment_type),
    characteristic text
);

create table if not exists users(
    login varchar(255) not null unique,
    password varchar(255) not null,
    primary key (login, password),
    role varchar(255) not null,
    full_name varchar(255),
    photo varchar(255)
);

create table if not exists operation_specifications(
    product varchar(255) not null,
        foreign key (product) references products(name),
    operation varchar(255) not null,
    serial_number integer not null,
    primary key (product, operation, serial_number),
    equipment_type varchar(255),
        foreign key (equipment_type) references equipment_types(equipment_type),
    operation_time interval not null
);

create table if not exists cake_decoration_specifications(
    product varchar(255) not null,
        foreign key (product) references products(name),
    cake_decoration integer not null,
        foreign key (cake_decoration) references cake_decorations(article),
    primary key (product, cake_decoration),
    count integer not null
);

create table if not exists ingredient_specifications(
    product varchar(255) not null,
        foreign key (product) references products(name),
    ingredient integer not null,
        foreign key (ingredient) references ingredients(article),
    primary key (product, ingredient),
    count integer not null
);

create table if not exists semi_finished_products(
    product varchar(255) not null,
        foreign key (product) references products(name),
    semi_finished_product varchar(255) not null,
        foreign key (semi_finished_product) references products(name),
    primary key (product, semi_finished_product),
    count integer not null
);

create table if not exists orders(
    order_number integer not null,
    date date not null,
    primary key (order_number, date),
    order_name varchar(255) not null,
    product varchar(255) not null,
        foreign key (product) references products(name),
    customer varchar(255) not null,
        foreign key (customer) references users(login),
    manager varchar(255),
        foreign key (manager) references users(login),
    price varchar(255) not null,
    planned_date date not null,
    examples varchar(255)
);

create table if not exists instrument (
    name varchar(255) not null primary key,
    description text,
    equipment_type varchar(255),
        foreign key (equipment_type) references equipment_types(equipment_type),
    wear_rate integer not null,
    supplier varchar(255),
        foreign key (supplier) references suppliers(name),
    date_of_purchase timestamp not null,
    count integer not null
);

