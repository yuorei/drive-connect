package db

import "drive-connect/db/model"

func (d *DB) GetUserById(id string) (*model.User, error) {
	var user *model.User
	if err := d.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DB) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User
	if err := d.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DB) CreateUser(user *model.User) error {
	if err := d.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateUser(id string, user *model.User) error {
	if err := d.Conn.Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteUser(id string) error {
	if err := d.Conn.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
