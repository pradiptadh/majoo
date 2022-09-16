
-- +migrate Up
CREATE TABLE IF NOT EXISTS `majoo`.`users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `user_name` VARCHAR(45) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB ; 

CREATE TABLE IF NOT EXISTS `majoo`.`merchants` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `merchant_name` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_merchants_users_idx` (`user_id` ASC),
  CONSTRAINT `fk_merchants_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `majoo`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `majoo`.`outlets` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `merchant_id` INT UNSIGNED NOT NULL,
  `outlet_name` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_outlets_merchants1_idx` (`merchant_id` ASC),
  CONSTRAINT `fk_outlets_merchants1`
    FOREIGN KEY (`merchant_id`)
    REFERENCES `majoo`.`merchants` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `majoo`.`transactions` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `merchant_id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `bill_total` DOUBLE NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_transactions_merchants1_idx` (`merchant_id` ASC),
  INDEX `fk_transactions_outlets1_idx` (`outlet_id` ASC),
  CONSTRAINT `fk_transactions_merchants1`
    FOREIGN KEY (`merchant_id`)
    REFERENCES `majoo`.`merchants` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_transactions_outlets1`
    FOREIGN KEY (`outlet_id`)
    REFERENCES `majoo`.`outlets` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- +migrate Down
DROP TABLE IF EXISTS `majoo`.`users`;
DROP TABLE IF EXISTS `majoo`.`merchants`;
DROP TABLE IF EXISTS `majoo`.`outlets`;
DROP TABLE IF EXISTS `majoo`.`transactions`;