DROP DATABASE IF EXISTS db_test;
CREATE DATABASE db_test;

\c db_test;

CREATE TYPE person_type AS ENUM('pf', 'pj', 'foreigner');

CREATE TABLE IF NOT EXISTS tbl_debtods(
  id             UUID UNIQUE NOT NULL,
  name           VARCHAR(255) NOT NULL,
  surname        VARCHAR(255) NOT NULL,
  bussiness_name VARCHAR(255),
  observation    TEXT,
  cpf_cnpj       VARCHAR(14) NOT NULL,
  person_type    person_type NOT NULL DEFAULT 'pf'::person_type,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS tbl_debtods_contact_numbers(
  number VARCHAR(14) NOT NULL,
  fk_debtod_id UUID NOT NULL,
  FOREIGN KEY (fk_debtod_id)
  REFERENCES tbl_debtods(id),
  PRIMARY KEY (number, fk_debtod_id)
);

CREATE TABLE IF NOT EXISTS tbl_debtods_contact_emails(
  email VARCHAR(255) NOT NULL,
  fk_debtod_id UUID NOT NULL,
  FOREIGN KEY (fk_debtod_id)
  REFERENCES tbl_debtods(id),
  PRIMARY KEY (email, fk_debtod_id)
);

CREATE TABLE IF NOT EXISTS tbl_debtods_addresses(
  street VARCHAR(255) NOT NULL,
  number INTEGER NOT NULL,
  zipcode INTEGER NOT NULL,
  neighborhood VARCHAR(50),
  observation TEXT,
  description TEXT,
  fk_debtod_id UUID NOT NULL,
  FOREIGN KEY (fk_debtod_id)
  REFERENCES tbl_debtods(id),
  PRIMARY KEY(fk_debtod_id, street, number)
);

CREATE TABLE IF NOT EXISTS tbl_debtods_invoices(
  id                     UUID UNIQUE NOT NULL,
  invoice_identification TEXT,
  invoice_status         BOOLEAN NOT NULL,
  delay_notified         BOOLEAN,
  dont_notify_until      DATE,
  observation            TEXT,
  description            TEXT,
  expired_date           DATE,
  fk_debtod_id           UUID NOT NULL,
  FOREIGN KEY (fk_debtod_id)
  REFERENCES tbl_debtods(id),
  PRIMARY KEY(id)
);