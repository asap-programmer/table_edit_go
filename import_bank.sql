
CREATE DATABASE IF NOT EXISTS bank;

CREATE TABLE IF NOT EXISTS `bank`.`individuals` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NOT NULL,
  `surname` VARCHAR(45) NOT NULL,
  `patronymic` VARCHAR(45) NULL DEFAULT NULL,
  `passport` VARCHAR(45) NOT NULL,
  `inn` VARCHAR(45) NOT NULL,
  `snils` VARCHAR(45) NOT NULL,
  `driver_license` VARCHAR(45) NULL DEFAULT NULL,
  `add_documents` VARCHAR(45) NULL DEFAULT NULL,
  `notice` VARCHAR(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `bank`.`borrowers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `inn` VARCHAR(45) NOT NULL,
  `is_individual` BIT(1) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `sum` DECIMAL(15,2) NOT NULL,
  `terms` VARCHAR(255) NOT NULL,
  `legal_notes` VARCHAR(255) NULL DEFAULT NULL,
  `contracts` VARCHAR(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `bank`.`loans` (
  `id` INT NOT NULL,
  `id_organization` INT NOT NULL,
  `id_individual` INT NOT NULL,
  `sum` DECIMAL(15,2) NOT NULL,
  `period` INT NOT NULL,
  `procent` DECIMAL(5,2) NOT NULL,
  `terms` VARCHAR(255) NOT NULL,
  `notices` VARCHAR(255) NULL DEFAULT NULL,
  `id_borrower` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `id_individual_idx` (`id_individual` ASC) VISIBLE,
  INDEX `id_borrower_idx` (`id_borrower` ASC) VISIBLE,
  CONSTRAINT `id_borrower`
    FOREIGN KEY (`id_borrower`)
    REFERENCES `bank`.`borrowers` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `id_individual_constraint`
    FOREIGN KEY (`id_individual`)
    REFERENCES `bank`.`individuals` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `bank`.`borrowings` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_individual` INT NOT NULL,
  `sum` DECIMAL(15,2) NOT NULL,
  `procent` DECIMAL(5,2) NOT NULL,
  `rate` INT NOT NULL,
  `period` INT NOT NULL,
  `terms` TEXT NOT NULL,
  `notices` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `id_individual_idx` (`id_individual` ASC) VISIBLE,
  CONSTRAINT `id_individual`
    FOREIGN KEY (`id_individual`)
    REFERENCES `bank`.`individuals` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO `bank`.`individuals` (id, first_name, surname, patronymic,passport,inn,snils,driver_license,add_documents,notice) VALUES
(17, 'Ivan', 'Ivanov', 'Ivanovich', '3848 672821', '418309800001', '12345678901', '1234 567890', 'A1B2C3D4E5', 'Pending approval'),
(19, 'Ilvir', 'Ilvirov', 'Ilvirovich', '2032 203033', '591174820009', '98765432109', '2345 678901', 'Z9Y8X7W6V5U4T3S2R1', 'Approved without issues'),
(21, 'Kirill', 'Kirillov', 'Kirillovich', '1233 236772', '737029340005', '13579246800', '3456 789012', 'N7M6L5K4J3H2G1F0E9', 'Data entry error'),

(23, 'Andrew', 'Andrewov', 'Andrewovich', '9219 417730', '848526790003', '24680135706', '4567 890123', 'P2Q3R4S5T6U7V8W9X0', 'Requires additional documentation'),
(25, 'Antonio', 'Antonionov', 'Antonionovich', '3321 777621', '204617530002', '11223344555', '5678 901234', 'L1K2J3H4G5F6D7S8A9', 'Document verified');


COMMIT;

