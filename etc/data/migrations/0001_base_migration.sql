-- +goose Up
create table safety_domain_blocklist
(
    id         bigint auto_increment
        primary key,
    domain     varchar(255) not null,
    created_at datetime     not null,
    updated_at datetime     null,
    constraint safety_domain_blocklist__domain
        unique (domain)
)
    collate = utf8mb4_unicode_ci;

-- +goose Down
drop table safety_domain_blocklist;
