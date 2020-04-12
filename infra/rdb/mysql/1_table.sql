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


