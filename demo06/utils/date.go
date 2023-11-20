package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type GORM_MODEL struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	CreatedAt LocalDateTimeFormat `gorm:"column:created_at;autoCreateTime;type:datetime" json:"created_at"`
	UpdatedAt LocalDateTimeFormat `gorm:"column:updated_at;autoUpdateTime;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index"`
}

type LocalDateTimeFormat time.Time

func (t LocalDateTimeFormat) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t *LocalDateTimeFormat) UnmarshalJSON(data []byte) error {
	tt, err := time.Parse("2006-01-02 15:04:05", string(data))
	if err != nil {
		return err
	}
	*t = LocalDateTimeFormat(tt)
	return nil
}

func (t *LocalDateTimeFormat) String() string {
	return fmt.Sprintf("%s", time.Time(*t).String())
}

// 存入格式化
func (t LocalDateTimeFormat) Value() (driver.Value, error) {
	log.Println("存入:", time.Time(t))
	return time.Time(t).Format("2006-01-02 15:04:05"), nil
}

// 取出格式化
func (t *LocalDateTimeFormat) Scan(v interface{}) error {
	switch value := v.(type) {
	case time.Time:
		*t = LocalDateTimeFormat(value)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t LocalDateTimeFormat) local() time.Time {
	return time.Time(t)
}
