# DDL for the dashboard database

# Data table, holds all of the specified fields
DROP TABLE IF EXISTS `Data`
CREATE TABLE Data (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `team_id` INTEGER(3) NOT NULL,
  `status` INTEGER(3) NOT NULL,
  `acceleration` INTEGER(10) NOT NULL,
  `position` INTEGER(10) NOT NULL,
  `velocity` INTEGER(10) NOT NULL,
  `battery_voltage` INTEGER(10) ,
  `battery_current` INTEGER(10) ,
  `battery_temperature` INTEGER(10) ,
  `pod_temperature` INTEGER(10) ,
  `stripe_count` INTEGER(10)
);

# Update table
# Keeps track of all of the updates in sequence
CREATE TABLE `Update` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `last_update_start` INTEGER(10) NOT NULL,
  `last_update_end` INTEGER(10) NOT NULL,
  FOREIGN KEY(`last_update_start`) REFERENCES `Data`(`id`),
  FOREIGN KEY(`last_update_end`) REFERENCES `Data`(`id`)
);
