-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `users_admin_groups` (
    id integer primary key auto_increment,
    user_id bigint,
    admin_group_id bigint
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users_admin_groups`;
