create table issue(id serial primary key,title text not null,description text,assigned_to text,created_by text,status bool);

create table users(userid serial primary key,email text unique not null,username text unique,firstname text not null,lastname text,password text not null);

