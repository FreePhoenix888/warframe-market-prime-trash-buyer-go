package models

type UserStatus string

func (s UserStatus) IsOnline() bool {
	return s == "online"
}

func (s UserStatus) IsOffline() bool {
	return s == "offline"
}

func (s UserStatus) IsIngame() bool {
	return s == "ingame"
}