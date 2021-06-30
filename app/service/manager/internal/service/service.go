package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	repo biz.Repo
}

func (s *Service) Hello(ctx *gin.Context) {
	str, err := s.repo.Hello()
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, str)
}
