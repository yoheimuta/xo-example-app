// Package models contains the types for schema 'test-xo-db'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// User represents a row from 'test-xo-db.users'.
type User struct {
	UserID int `json:"user_id"` // user_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO test-xo-db.users (` +
		`user_id` +
		`) VALUES (` +
		`?` +
		`)`

	// run query
	XOLog(sqlstr, u.UserID)
	_, err = db.Exec(sqlstr, u.UserID)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Update statements omitted due to lack of fields other than primary key

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM test-xo-db.users WHERE user_id = ?`

	// run query
	XOLog(sqlstr, u.UserID)
	_, err = db.Exec(sqlstr, u.UserID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByUserID retrieves a row from 'test-xo-db.users' as a User.
//
// Generated from index 'users_user_id_pkey'.
func UserByUserID(db XODB, userID int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`user_id ` +
		`FROM test-xo-db.users ` +
		`WHERE user_id = ?`

	// run query
	XOLog(sqlstr, userID)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, userID).Scan(&u.UserID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}