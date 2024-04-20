create table migrations (
  id integer not null primary key autoincrement,
  migration varchar(255) not null,
  applied_at datetime not null,
  PRIMARY KEY (id)
);

create table accounts (
  id integer not null primary key autoincrement,
  name varchar(255) not null,
  created_at datetime not null,
  updated_at datetime not null,
  currency varchar(3) not null default EUR,
  initial_balance decimal(10, 2) not null
);

drop index index_categories_on_name;
drop table transaction_items;
drop table transactions;
drop table items;
drop table categories;
drop table accounts;
