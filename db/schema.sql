CREATE TABLE dex (
  id varchar(4) PRIMARY KEY,
  name varchar(15),
  gen int,
  type_1 varchar(10),
  type_2 varchar(10),
  height float(5, 3),
  weight float(10, 3)
);

CREATE TABLE types (
  type varchar(10) UNIQUE
);

CREATE TABLE stats (
  id varchar(4) PRIMARY KEY,
  hp int,
  atk int,
  def int,
  spl_atk int,
  spl_def int,
  speed int,
  total int
);

CREATE TABLE dex_entries (
  id varchar(4) PRIMARY KEY,
  entry1 varchar(300),
);

CREATE TABLE favourites (
  id varchar(4) PRIMARY KEY
);

ALTER TABLE dex ADD FOREIGN KEY (type_1) REFERENCES types (type);

ALTER TABLE dex ADD FOREIGN KEY (type_2) REFERENCES types (type);

ALTER TABLE stats ADD FOREIGN KEY (id) REFERENCES dex (id);

ALTER TABLE dex_entries ADD FOREIGN KEY (id) REFERENCES dex (id);

ALTER TABLE favourites ADD FOREIGN KEY (id) REFERENCES dex (id);
