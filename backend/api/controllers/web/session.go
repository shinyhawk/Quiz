package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ip-05/quizzus/api/middleware"
	"github.com/ip-05/quizzus/entity"
)

type ISessionService interface {
	GetSessions(userId, limit int) (*[]entity.GameSession, *[]entity.Leaderboard)
	NewSession(id, userId, instId int) uint
	EndSession(id, userId, instId, place, questions, players int, points float64) uint
}

type SessionController struct {
	session ISessionService
}

func NewSessionController(session ISessionService) *SessionController {
	return &SessionController{session: session}
}

func (s *SessionController) GetSessions(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}

	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	sessions, _ := s.session.GetSessions(int(user.Id), limit)

	c.JSON(http.StatusOK, sessions)
	return
}
