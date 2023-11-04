package db

import "drive-connect/db/model"

func (d *DB) GetBoardById(id string) (*model.Board, error) {
	var board *model.Board
	if err := d.Conn.Where("id = ?", id).First(&board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (d *DB) ReadAllBoard() ([]*model.Board, error) {
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

func (d *DB) GetCommentById(id string) (*model.Comment, error) {
	var comment *model.Comment
	if err := d.Conn.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (d *DB) GetCommentList() ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := d.Conn.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (d *DB) CreateComment(comment *model.Comment) error {
	if err := d.Conn.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateComment(id string, comment *model.Comment) error {
	if err := d.Conn.Where("id = ?", id).Updates(comment).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteComment(id string) error {
	if err := d.Conn.Where("id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}
