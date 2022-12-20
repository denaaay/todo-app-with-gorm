package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	result := u.db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSession(token string) error {
	result := u.db.Where("token = ?", token).Delete(&model.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	result := u.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(model.Session{Token: session.Token, Username: session.Username, Expiry: session.Expiry})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	result := model.Session{}
	resp := u.db.Raw("SELECT * FROM sessions WHERE username = ?", name).Scan(&result)
	if resp.Error != nil {
		return model.Session{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return model.Session{}, errors.New("record not found")
	}
	return result, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	result := model.Session{}
	resp := u.db.Raw("SELECT * FROM sessions WHERE token = ?", token).Scan(&result)
	if resp.Error != nil {
		return model.Session{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return model.Session{}, errors.New("record not found")
	}
	return result, nil // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}
	// TODO: replace this

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil
}

func (u *SessionsRepository) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
