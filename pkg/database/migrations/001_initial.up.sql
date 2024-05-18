create table accounts (
  id integer not null primary key autoincrement,
  name varchar(255) not null,
  iban varchar(34),
  bic varchar(12),
  account_number varchar(10),
  account_type integer not null default 0,
  opening_balance integer not null default 0,
  opening_balance_date datetime not null,
  include_in_net_worth boolean not null,
  created_at datetime not null,
  updated_at datetime not null,
  currency varchar(3) not null default 'EUR',
  notes varchar(1024) not null,
  deleted boolean not null default false
);

create table categories (
  id integer not null primary key autoincrement,
  name varchar(255) not null,
  created_at datime not null,
  updated_at datetime not null,
  color_hex varchar(6) not null
);

create table transactions (
  id integer not null primary key autoincrement,
  from_account_id integer not null,
  to_account_id integer not null,
  total_amount integer not null,
  created_at datetime not null,
  updated_at datetime not null,
  foreign key (from_account_id) references accounts (id),
  foreign key (to_account_id) references accounts (id)
);

create table items (
  id integer not null primary key autoincrement,
  name varchar(255) not null,
  price integer not null,
  transaction_id integer not null,
  category_id integer,
  created_at datetime not null,
  updated_at datetime not null,
  foreign key (transaction_id) references transactions (id),
  foreign key (category_id) references categories (id)
);

create table items_categories (
  id integer not null primary key autoincrement,
  item_id integer not null,
  category_id integer not null,
  foreign key (item_id) references items (id),
  foreign key (category_id) references categories (id)
);

create index index_categories_on_name on categories (name);
