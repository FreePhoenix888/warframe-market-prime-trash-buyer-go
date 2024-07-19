package models

type Platform string

func (p Platform) IsPC() bool {
	return p == "pc"
}