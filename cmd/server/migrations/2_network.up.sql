CREATE TABLE IF NOT EXISTS networks (
  id INT AUTO_INCREMENT NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  cidr VARCHAR(3) NOT NULL,
  dns varchar(253) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE INDEX idx_networks_name ON networks(name);

