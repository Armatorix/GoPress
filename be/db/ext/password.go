package ext

import (
	"database/sql/driver"
	"fmt"

	"github.com/Armatorix/GoPress/be/x/xaes"
)

var encClient *xaes.Crypto

func Init(key []byte) {
	encClient = xaes.New(key)
}

type Password string

func (s Password) String() string {
	return string(s)
}

func (e Password) IsEmpty() bool {
	return len(e) == 0
}

func (e Password) Value() (driver.Value, error) {
	if e == "" {
		return "", nil
	}
	return encClient.Encrypt(string(e))
}

func (e *Password) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var s string
	switch src := src.(type) {
	case []byte:
		s = string(src)
	case string:
		s = src
	default:
		return fmt.Errorf("incompatible type %T for password", src)
	}

	if s == "" {
		*e = ""
		return nil
	}

	dec, err := encClient.Decrypt(s)
	if err != nil {
		return fmt.Errorf("error decrypting password: %v", err)
	}
	*e = Password(dec)
	return nil
}
