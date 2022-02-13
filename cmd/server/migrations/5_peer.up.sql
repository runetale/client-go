CREATE TABLE IF NOT EXISTS peers (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  user_id INT NOT NULL,
  setup_key_id INT NOT NULL UNIQUE,
  organization_id INT NOT NULL,
  user_group_id INT NOT NULL,
  client_pub_key VARCHAR(800) NOT NULL,
  network_id INT NOT NULL,
  ip VARCHAR(18) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (setup_key_id) REFERENCES setup_keys(id),
  FOREIGN KEY (organization_id) REFERENCES organizations(id),
  FOREIGN KEY (user_group_id) REFERENCES user_groups(id),
  FOREIGN KEY (network_id) REFERENCES networks(id)
);
