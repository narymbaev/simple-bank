create table accounts (
    id bigserial primary key,
    owner varchar not null,
    balance bigint not null,
    currency varchar not null,
    created_at timestamptz not null default now()
);

create table entries (
    id bigserial primary key,
    account_id bigint not null references accounts,
    amount bigint not null,
    created_at timestamptz not null default now()
);

create table transfers
(
    id              bigserial primary key,
    from_account_id bigint      not null references accounts,
    to_account_id   bigint      not null references accounts,
    amount          bigint      not null,
    created_at      timestamptz not null default now()
);

comment on column entries.amount is 'can be negative or positive';
comment on column transfers.amount is 'must be positive';

