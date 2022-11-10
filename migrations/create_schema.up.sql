create table clients (
    id bigserial primary key,
    name varchar(255) not null
);

create table balances (
    id bigserial primary key,
    amount bigint not null,
    updated_at timestamp not null default now(),
    client_id bigserial,
    foreign key (client_id) references clients (id)
);

create table transactions (
    id bigserial primary key,
    sender_id bigserial,
    receiver_id bigserial,
    foreign key (sender_id) references clients (id),
    foreign key (receiver_id) references clients (id),
    amount bigint not null,
    status varchar(30) not null
)
