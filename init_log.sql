# 	ID        uint `gorm:"primary_key"`
# 	CreatedAt time.Time
# 	UpdatedAt time.Time
# 	DeletedAt *time.Time `sql:"index"`
CREATE TABLE IF NOT EXISTS log (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  created_at DATETIME     NOT NULL,
  updated_at DATETIME     NOT NULL,
  deleted_at DATETIME,

  subject    VARCHAR(32)  NOT NULL,
  content    VARCHAR(128) NOT NULL
)