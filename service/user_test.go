package service

import "testing"

func TestSaveUser(t *testing.T) {
	SaveUser()
}

func TestSaveUserBath(t *testing.T) {
	SaveUserBatch()
}

func TestSaveUserUpdate(t *testing.T) {
	UpdateUser()
}
func TestSaveUserCount(t *testing.T) {
	CountUser()
}
