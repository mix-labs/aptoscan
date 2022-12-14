package types

import (
	"gorm.io/gorm"
	"time"
)

type ProcessorStatus struct {
	Name         string
	StartVersion int64
	EndVersion   int64
	Success      bool
	Detail       string

	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime;not null"`
}

func (ProcessorStatus) TableName() string {
	return "processor_statuses"
}

//func ProcessorStatusFromVersions(name string, startVersion, endVersion int64, success bool, detail string) []ProcessorStatus {
//	var status []ProcessorStatus
//	newProcessorStatus := NewProcessorStatus(name, 0, success, detail)
//	for i := startVersion; i < endVersion; i++ {
//		newProcessorStatus.Version = i
//		status = append(status, *newProcessorStatus)
//	}
//	return status
//}

func AutoCreateProcessorStatusTable(db *gorm.DB) error {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&ProcessorStatus{})
	if err != nil {
		return err
	}
	return nil
}
