package table

type SnapshotOptions struct {
	Crontab string `json:"crontab,omitempty"`
	Table   string `json:"table,omitempty"`
}

type DeleteOptions struct {
	Backup      bool   `json:"backup,omitempty"`
	BackupTable string `json:"backup_table,omitempty"`
}

type HistoryOptions struct {
	Backup      bool   `json:"backup,omitempty"`
	BackupTable string `json:"backup_table,omitempty"`
}

type BackupOptions struct {
	Deleted     bool   `json:"deleted,omitempty"`
	DeleteTable string `json:"deleted_table,omitempty"`
	Updated     bool   `json:"updated,omitempty"`
	UpdateTable string `json:"updated_table,omitempty"`
}
