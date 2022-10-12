-- +goose Up
-- +goose StatementBegin
create table user(
   id         integer primary key,
   username   text not null,
   name       text not null,
   created_at timestamp default current_timestamp,
   is_current_payer boolean not null default 0
);

create table user_payment_history(
   id         integer primary key not null,
   created_at timestamp default current_timestamp,
   action     text check(action in ('pay','skip', 'unknown')) not null default 'unknown',
   user_id    integer,
   foreign    key(user_id) references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
drop table user_payment_history;
-- +goose StatementEnd
