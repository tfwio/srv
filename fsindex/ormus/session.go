package ormus

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Session represents users who are logged in.
type Session struct {
	ID      int64     `gorm:"auto_increment;unique_index;primary_key;column:id"`
	UserID  int64     `gorm:"column:user_id"` // [users].[id]
	Port    int       `gorm:"column:port"`    // running multiple server instance/port(s)?
	Created time.Time `gorm:"not null;column:created"`
	Expires time.Time `gorm:"not null;column:expires"`
	Keep    *bool     `gorm:"column:keep"`
	SessID  string    `gorm:"not null;column:sessid"`
}

// TableName Set User's table name to be `users`
func (Session) TableName() string {
	return "sessions"
}

// EnsureTableSessions creates table [sessions] if not exist.
func EnsureTableSessions() {
	var s Session
	db, err := gorm.Open("sqlite3", datasource)
	defer db.Close()
	if err != nil {
		fmt.Printf("error: ensuring database: sessions\n")
	} else {
		if !db.HasTable(s) {
			db.CreateTable(s)
		}
	}
}

// SessionFind attempts to find a session from SessionID
func SessionFind(sessid string) (Session, User) {
	var s Session
	var u User
	db, err := gorm.Open("sqlite3", datasource)
	defer db.Close()
	if err != nil {
		db.Where("[sessid] = ?", sessid).First(&s)
		db.Where("[id] = ?", s.UserID).First(&u)
		// Where("")
	} else {
		println("coundn't find session")
	}
	return s, u
}
