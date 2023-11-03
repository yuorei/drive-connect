package db

import "drive-connect/db/model"

func (d *DB) GetBoardById(id string) (*model.Board, error) {
	var board *model.Board
	if err := d.Conn.Where("id = ?", id).First(&board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (d *DB) GetBoardList() ([]*model.Board, error) {
	var boards []*model.Board
	if err := d.Conn.Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

func (d *DB) CreateBoard(board *model.Board) error {
	if err := d.Conn.Create(board).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateBoard(id string, board *model.Board) error {
	if err := d.Conn.Where("id = ?", id).Updates(board).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteBoard(id string) error {
	if err := d.Conn.Where("id = ?", id).Delete(&model.Board{}).Error; err != nil {
		return err
	}
	return nil
}
