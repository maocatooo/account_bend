package model

import "gorm.io/gorm"

func ErrRecordNotFound(err error) error {
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}
