DROP SCHEMA IF EXISTS letustalk;
CREATE SCHEMA letustalk;
USE letustalk;

CREATE TABLE `invoice_detail` (
    `invoice_id` VARCHAR(255) NOT NULL,
    `amount` bigint(20) NOT NULL,
    `paid` boolean NOT NULL default false,
    PRIMARY KEY (`invoice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `invoice_detail` (`invoice_id`, `amount`) VALUES ("f0c28384-3aa4-3f87-9fba-66a0aa62c504", 5000);

CREATE TABLE `account_credentials` (
    `account_id` VARCHAR(255) NOT NULL,
    `mail` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `position` VARCHAR(255) NOT NULL default 'general',
    PRIMARY KEY (`account_id`),
    UNIQUE KEY (`mail`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `account_credentials` (`account_id`, `mail`, `password`) VALUES ("f0c28384-3aa4-3f87-9fba-66a0aa62c504", "bambootuna@gmail.com", "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1");

CREATE TABLE `account_detail` (
    `account_id` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `introduction` VARCHAR(2000) NOT NULL,
    PRIMARY KEY (`account_id`),
    FOREIGN KEY(`account_id`) REFERENCES `account_credentials`(`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `account_detail` (`account_id`, `name`, `introduction`) VALUES ("f0c28384-3aa4-3f87-9fba-66a0aa62c504", "たけちゃ", "自称エンジニア");


CREATE TABLE `reservation` (
    `reservation_id` VARCHAR(255) NOT NULL,
    `child_account_id` VARCHAR(255) NOT NULL,
    `invoice_id` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`reservation_id`),
    FOREIGN KEY(`child_account_id`) REFERENCES `account_credentials`(`account_id`),
    FOREIGN KEY(`invoice_id`) REFERENCES `invoice_detail`(`invoice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `reservation` (`reservation_id`, `child_account_id`, `invoice_id`) VALUES ("1", "f0c28384-3aa4-3f87-9fba-66a0aa62c504", "f0c28384-3aa4-3f87-9fba-66a0aa62c504");

CREATE TABLE `schedule` (
    `schedule_id` VARCHAR(255) NOT NULL,
    `parent_account_id` VARCHAR(255) NOT NULL,
    `from` datetime NOT NULL,
    `to` datetime NOT NULL,
    `reservation_id` VARCHAR(255),
    PRIMARY KEY (`schedule_id`),
    FOREIGN KEY(`parent_account_id`) REFERENCES `account_credentials`(`account_id`),
    FOREIGN KEY(`reservation_id`) REFERENCES `reservation`(`reservation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `schedule` (`schedule_id`, `parent_account_id`, `from`, `to`, `reservation_id`) VALUES ("1", "f0c28384-3aa4-3f87-9fba-66a0aa62c504", "2020-04-20 00:00:00", "2020-04-20 00:30:00", "1");

CREATE TABLE `schedule_detail` (
    `schedule_id` VARCHAR(255) NOT NULL,
    `unit_price` bigint(20) NOT NULL,
    PRIMARY KEY (`schedule_id`),
    FOREIGN KEY(`schedule_id`) REFERENCES `schedule`(`schedule_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `schedule_detail` (`schedule_id`, `unit_price`) VALUES ("1", 5000);
