-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `admin_groups` (
    id integer primary key auto_increment,
    name varchar(255) NOT NULL UNIQUE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `admin_groups`;
