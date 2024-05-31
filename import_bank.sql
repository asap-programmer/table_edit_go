
CREATE DATABASE IF NOT EXISTS bank;

CREATE TABLE individuals (
  id int(11) NOT NULL,
  first_name varchar(255) NOT NULL,
  second_name varchar(255) NOT NULL,
  passport varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=cp1251;

INSERT INTO individuals (id, first_name, second_name, passport) VALUES
(17, 'Ivan', 'Ivanov', '3848 672821'),
(23, 'Ilvir', 'Ilvirov', "2032 203033"),
(43, 'Kirill', 'Kirillov', '1233 2367372'),
(48, 'Andrew', 'Andrewov', '8972 273772'),
(50, 'Antonio', 'Antonionov', '666 666666'),

ALTER TABLE individuals
  ADD PRIMARY KEY (id);

ALTER TABLE individuals
  MODIFY id int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=100;

COMMIT;

