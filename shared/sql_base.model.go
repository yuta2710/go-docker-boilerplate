package shared

import "time"

type BaseSQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (m *BaseSQLModel) Mask(dbType DbType) {
	uid := NewUID(uint32(m.Id), int(dbType), 1)
	m.FakeId = &uid
}
