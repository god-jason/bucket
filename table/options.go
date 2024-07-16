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
