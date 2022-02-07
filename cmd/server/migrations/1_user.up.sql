CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  provider_id VARCHAR(255) NOT NULL,
  provider VARCHAR(255) NOT NULL,
  org_group_id INT NOT NULL,
  network_id INT NOT NULL,
  user_group_id INT NOT NULL,
  permission INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (org_group_id) REFERENCES orgs(id),
  FOREIGN KEY (network_id) REFERENCES networks(id),
  FOREIGN KEY (user_group_id) REFERENCES user_groups(id)
);

CREATE INDEX idx_users_provider ON users(provider);
CREATE INDEX idx_users_provider_id ON users(provider_id);
