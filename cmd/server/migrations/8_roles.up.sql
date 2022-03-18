CREATE TABLE IF NOT EXISTS roles (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  admin_network_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  -- RWX, RW, R => 0, 1, 2
  permission INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (admin_network_id) REFERENCES admin_networks(id)
);

CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_admin_network_id ON roles(admin_network_id);