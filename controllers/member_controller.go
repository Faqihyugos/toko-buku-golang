package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	domain "toko/domain/model"
	"toko/services"
	"toko/utils"
)

type memberController struct {
	MemberService domain.IMemberService
}

const (
	SIGN_UP_PATH     = "/member/signup"
	SIGN_IN_PATH     = "/member/signin"
	BUYS_BOOK_PATH   = "/member/book/buys/:memberId"
	GET_HISTORY_PATH = "/member/history/:id"
	ACTIVATED_PATH   = "/member/activated/:memberId"
)

func NewMemberController(db *sql.DB, r *gin.RouterGroup) {
	Controller := memberController{MemberService: services.NewMemberService(db)}
	r.POST(SIGN_UP_PATH, Controller.SignUpMember)
	r.POST(SIGN_IN_PATH, Controller.SignInMember)
	r.POST(BUYS_BOOK_PATH, Controller.Buys)
	r.GET(GET_HISTORY_PATH, Controller.HistoryTrx)
	r.PUT(ACTIVATED_PATH, Controller.ActivatedMember)
}

func (m *memberController) HistoryTrx(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError("Internal server error"))
	}

	histories, errget := m.MemberService.GetHistoryTrxMember(id)
	if errget != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError("Internal server error"))
		return
	}
	c.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", histories))
}

func (m *memberController) Buys(c *gin.Context) {
	param := c.Param("memberId")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Internal Server Error"})
	}
	var requestBook domain.RequestBuyBooks
	var buys []domain.Buy

	errBind := c.ShouldBindJSON(&requestBook)
	if errBind != nil {
		theErr := utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	for _, v := range requestBook.Buys {
		buys = append(buys, v)
	}

	//fmt.Println("controller: ", buys)
	purchases, errPurchase := m.MemberService.Buys(buys, id)
	if errPurchase != nil {
		c.JSON(http.StatusBadRequest, errPurchase)
	} else {
		c.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", purchases))
	}
}

func (m *memberController) SignInMember(c *gin.Context) {
	var member domain.MemberLogin
	errBind := c.ShouldBindJSON(&member)
	if errBind != nil {
		theErr := utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	_, err := m.MemberService.SignIn(&member)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please validate email or password correct!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Verified"})

}

func (m *memberController) SignUpMember(c *gin.Context) {
	var member domain.Member
	errBind := c.ShouldBindJSON(&member)
	if errBind != nil {
		theErr := utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	MemberNew, err := m.MemberService.SignUp(&member)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "Member registration successfully", MemberNew))
}

func (m *memberController) ActivatedMember(c *gin.Context) {
	param := c.Param("memberId")
	id, errParse := strconv.Atoi(param)
	if errParse != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Internal Server Error"})
	}

	err := m.MemberService.ActivatedMember(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Member activated successfully"})
}
