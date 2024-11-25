package model

type Balance struct {
	Owner     User
	Current   Money
	Withdrawn Money
}
