package data

// the required sql queries
const (
	InsertUser    = "insert into users(email,username,firstname,lastname,password) values($1,$2,$3,$4,$5) on conflict(email) do update set firstname=excluded.firstname returning userid, case when xmax::text::int > 0 then 'updated' else 'inserted' end"
	CheckPassword = "select password,userid from users where email=$1"
	InsertIssue   = "insert into issue(title,description,assigned_to,created_by,status) values($1,$2,$3,$4,$5) returning id"
)
