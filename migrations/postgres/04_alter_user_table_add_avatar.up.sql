alter table users add column avatar varchar(150);
alter table users add deleted_at timestamp;
ALTER TABLE users ADD CONSTRAINT phone_number UNIQUE(phone_number);