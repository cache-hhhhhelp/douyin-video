create table video
(
    id             bigint auto_increment,
    author_id      bigint       not null,
    title          varchar(255) not null,
    created_at     bigint       not null,
    primary key (id)
);



