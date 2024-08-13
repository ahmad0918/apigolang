package nulltype

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	mysql.NullTime
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}
