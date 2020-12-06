package postgres

const AddClient = `insert into clients(name, surname, login, password, age, gender, phone , status)
	values($1, $2, $3, $4, $5, $6, $7, $8)`

const AddAccount = `insert into accounts( client_id, account_number, balance, status, card_number )
	values( $1, $2, $3, $4, $5 )`

const updateLimit = `update accounts
	set limit_transfer = $limit_transfer, limit_payment = $limit_payment where id = $accountId`

const AddATMs = `insert into atms( name, status )
	values( $name, $status )`

const changeStatusATM = `update atms
	set status = $status where id = $atmId`

const AddService = `insert into services( name )
	values( $1 )`

const GetAllClients = `select *from clients`

const GetAllAccounts = `select * from accounts a left join clients c on a.client_id = c.id`

const GetAllATMs = `select *from ATMs`

const LoginSQL = `select login, password from clients where login = ($1)`

const SearchClientByLogin = `select id, surname from clients where login = ($1)`

const SearchAccountByID = `select id, name, accountNumber, balance, locked from accounts where locked = true and user_id = ?`

const GetAllServices = `select id, name from services`

