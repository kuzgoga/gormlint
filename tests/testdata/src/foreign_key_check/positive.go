package foreign_key_check

type PrepTask struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	TaskId      uint
	WorkAreaId  uint
	WorkArea    `gorm:"foreignKey:WorkAreaIds;constraint:OnDelete:CASCADE;"` // want "Foreign key \"WorkAreaIds\" mentioned in tag at field \"WorkArea\" doesn't exist in model \"PrepTask\""
	Deadline    int64
}

type WorkArea struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Performance uint
	PrepTasks   []PrepTask `gorm:"constraint:OnDelete:CASCADE;"`
}

type Shift struct {
	Id            uint `gorm:"primaryKey"`
	Description   string
	ProductAmount uint
	ShiftDate     int64
	WorkAreaId    string   // want "Foreign key should have type like int, not \"string\""
	WorkArea      WorkArea `gorm:"foreignKey:WorkAreaId;"`
}
