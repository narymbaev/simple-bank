alter table if exists accounts drop constraint owner_currency_key;

alter table if exists accounts drop constraint accounts_owner_fkey;

drop table if exists users;