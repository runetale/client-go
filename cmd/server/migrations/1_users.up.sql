CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  provider_id VARCHAR(255) NOT NULL,
  admin_network_id INT NOT NULL,
  network_id INT NOT NULL,
  user_group_id INT NOT NULL,
  role_id INT NOT NULL,
  provider VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  FOREIGN KEY (admin_network_id) REFERENCES admin_networks(id),
  FOREIGN KEY (network_id) REFERENCES networks(id),
  FOREIGN KEY (user_group_id) REFERENCES user_groups(id)
  FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE INDEX idx_users_provider ON users(provider);
CREATE INDEX idx_users_provider_id ON users(provider_id);
CREATE INDEX idx_users_admin_network_id ON users(admin_network_id);
