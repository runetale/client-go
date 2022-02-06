CREATE TABLE IF NOT EXISTS jobs (
  id INT AUTO_INCREMENT NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  user_id INT NOT NULL,
  org_id INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (org_id) REFERENCES orgs(id),
  PRIMARY KEY(id)
);

CREATE INDEX idx_jobs_org_id ON jobs(org_id);
CREATE INDEX idx_jobs_user_id ON jobs(user_id);
CREATE INDEX idx_jobs_name ON jobs(name);
