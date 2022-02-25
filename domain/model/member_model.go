package model

import "toko/utils"

type Member struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    int    `json:"Status"`
}

type MemberLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReqBuy struct {
	Qty int `json:"qty"`
}

type Purchase struct {
	Id         int `json:"id"`
	Book       *Book
	Qty        int `json:"qty"`
	TotalPrice int `json:"total_price"`
}

type Buy struct {
	BookID int `json:"bookId"`
	Qty    int `json:"qty"`
}

type RequestBuyBooks struct {
	Buys []Buy `json:"buys"`
}

type ResponseHistoryMember struct {
	User       string `json:"user"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
}

type IMemberRepository interface {
	Find() ([]*Member, utils.MessageErr)
	AddMember(member *Member) (*Member, utils.MessageErr)
	UpdateStatus(member *Member, status int) utils.MessageErr
	FindByEmail(memberLogin *MemberLogin) (*Member, error)

	//	note pointer
	AddBooks(purchases []Purchase, memberId int) ([]Purchase, utils.MessageErr)
	FindHistoryByMember(memberId int) ([]*ResponseHistoryMember, utils.MessageErr)
	FindMemberById(id int) (*Member, utils.MessageErr)
}

type IMemberService interface {
	FindMembers() ([]*Member, utils.MessageErr)
	SignUp(member *Member) (*Member, utils.MessageErr)
	SignIn(member *MemberLogin) (*Member, error)
	ActivatedMember(memberId int) utils.MessageErr
	Buys(buys []Buy, memberId int) ([]Purchase, utils.MessageErr)
	GetHistoryTrxMember(memberId int) ([]*ResponseHistoryMember, utils.MessageErr)
}
