package model

import "toko/utils"

type Member struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LoggedIn  int    `json:"logged_in"`
}

type MemberLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReqBuy struct {
	Qty int `json:"qty"`
}

type Purchase struct {
	Book       *Book
	Qty        int `json:"qty"`
	TotalPrice int `json:"total_price"`
}

type IMemberRepository interface {
	AddMember(member *Member) (*Member, utils.MessageErr)
	UpdateStatus(member *Member, status int) utils.MessageErr
	FindByEmail(memberLogin *MemberLogin) (*Member, utils.MessageErr)
}

type IMemberService interface {
	SignUp(member *Member) (*Member, utils.MessageErr)
	SignIn(member *MemberLogin) utils.MessageErr
	BuyBook(id int, bookBuy *ReqBuy) (*Purchase, utils.MessageErr)
}
