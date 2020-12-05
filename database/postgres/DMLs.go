package postgres

const addClient = `insert into clients( name, surname, login, password, age, gender, phone )
	values( $name, $surname, $login, $password, $age, $gender, $phone )`

const AddAccount = `insert into clientsaccounts( client_id, account_number, balance, status, card_number )
	values( $client_id, $account_number, $balance, $status, $card_number )`

const updateLimit = `update clientsaccounts
	set limit_transfer = $limit_transfer, limit_payment = $limit_payment where id = $accountId`

const AddATMs = `insert into atms( address, status )
	values( $address, $status )`

const changeStatusATM = `update atms
	set status = $status where id = $atmId`

const addService  = `insert into services( name )
	values( $serviceName )`

const getAllClients = `select *from clients`

const GetAllAccounts = `select * from accounts a left join clients c on a.client_id = c.id`

const GetAllATMs = `select *from ATMs`

const LoginSQL = `select login, password from clients where login = ?`

const SearchClientByLogin = `select id, surname from clients where login = ?`

const searchAccountByIDSql = `select id, name, accountNumber, balance, locked from accounts where locked = true and user_id = ?`

const getAllServices = `select id, name, price from services`

