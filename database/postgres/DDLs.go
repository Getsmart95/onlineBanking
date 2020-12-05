package postgres

const clientsTable = `create table if not exists clients(
	id integer primary key autoincrement,
	name text not null,
	surname text not null,
	login not null unique,
	password text not null check ( length( password ) >= 8 ),
	age integer not null,
	gender text not null,
	phone integer not null,
	status boolean not null,
	verified_at datetime default CURENT_TIMESTAMP
)`

const clientsAccountsTable = `create table if not exists accounts(
	id integer primary key autoincrement,
	client_id integer references clients,
	account_number integer not null,
	balance integer not null check ( balance >= 0 ),
	status boolean not null,
	card_number integer not null unique check ( length( card_number ) = 16 ),
	//cvv integer not null check ( length ( cvv ) = 3 ),
	//pin_code integer not null check ( length ( pin_code ) = 4 ),
	limit_transfer integer not null default 3000,
	limit_payment integer not null default 4000,
	created_at datetime default CURENT_TIMESTAMP,
	until_at datetime default CURRENT_TIMESTAMP + 3
)`

const ATMsTable = `create table if not exists atms(
	id integer primary key autoincrement,
	address_id references atmaddresses,
	status boolean not null,
	created_at datetime CURRENT_TIMESTAMP
)`

const servicesTable = `create table if not exists services(
	id integer primary key autoincrement,
	name text not null
)`

const historiesTable = `create table if nont exists histories(
	id integer primary key autoincrement,
	sender_id references clients,
	recipient_id references clients,
	money integer not null,
	message text not null,
	service_id references services,
	created_at datetime default CURRENT_TIMESTAMP
)`

const atmAddresses  =  `create table if not exists addresses(
	id integer primary key autoincrement,
	country text not null,
	city text not null,
	street text not null,
	home integer not null,
	apartment integer not null
)`