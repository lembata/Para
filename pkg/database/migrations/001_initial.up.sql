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
  notes varchar(1024) not null
);

create table categories (
  id integer not null primary key autoincrement,
  name varchar(255) not null,
  created_at datetime not null,
  updated_at datetime not null,
  color_hex varchar(6) not null
);

create table items (
  id integer not null primary key autoincrement,
  name varchar(255) not null,
  price decimal(10,2) not null,
  created_at datetime not null,
  updated_at datetime not null,
  category_id integer,
  foreign key (category_id) references categories (id)
);


create table transactions (
  id integer not null primary key autoincrement,
  account_id integer not null,
  quantity integer not null,
  created_at datetime not null,
  updated_at datetime not null,
  total decimal(10, 2) not null,
  foreign key (account_id) references accounts (id)
);


create table transaction_items (
  id integer not null primary key autoincrement,
  transaction_id integer not null,
  item_id integer not null,
  created_at datetime not null,
  updated_at datetime not null,
  foreign key (transaction_id) references transactions (id),
  foreign key (item_id) references items (id)
);

create index index_categories_on_name on categories (name);
