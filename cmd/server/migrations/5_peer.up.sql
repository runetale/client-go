CREATE TABLE IF NOT EXISTS peers (
  id INT AUTO_INCREMENT NOT NULL UNIQUE,

  user_id INT NOT NULL,
  org_id INT NOT NULL,
  user_group_id INT NOT NULL,
  network_id INT NOT NULL,

  ip VARCHAR(18) NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (org_id) REFERENCES orgs(id),
  FOREIGN KEY (user_group_id) REFERENCES user_groups(id),
  FOREIGN KEY (network_id) REFERENCES networks(id),
  PRIMARY KEY(id)
);
