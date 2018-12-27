package data

// the required sql queries
const (
	InsertUser    = "insert into users(email,username,firstname,lastname,password,access_token) values($1,$2,$3,$4,$5,$6) on conflict(email) do update set firstname=excluded.firstname returning userid "
	CheckPassword = "select password from users where email=$1"
)
