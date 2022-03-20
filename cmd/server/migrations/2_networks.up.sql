CREATE TABLE IF NOT EXISTS networks (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  admin_network_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  ip VARCHAR(15) NOT NULL,
  cidr INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (admin_network_id) REFERENCES admin_networks(id)
);

CREATE INDEX idx_networks_name ON networks(name);

