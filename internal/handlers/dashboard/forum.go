package dashboard

import (
	"errors"
	"forum/internal/handlers/dashboard/render"
	"forum/internal/handlers/dashboard/request"
	"forum/internal/pkg/constants"
	"forum/internal/pkg/helper"
	"forum/internal/services"
	"forum/pkg/res"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var forumService = services.ForumService

func CreateForum(c *gin.Context) {
	input := &request.Forum{}
	if !helper.Validator(c, input) {
		return
	}

	err := forumService.Create(input)
	if err == nil {
		c.JSON(http.StatusOK, res.JsonSuccess())
		return
	}

	if errors.Is(services.ErrNameDuplicate, err) || errors.Is(services.ErrForumIdDuplicate, err) {
		c.JSON(http.StatusOK, res.JsonErrorMessage(constants.LogicError, err.Error()))
	} else {
		c.JSON(http.StatusOK, res.JsonErrorMessage(constants.InternalError, services.ErrInternal.Error()))
	}
	c.Abort()
}

func UpdateForum(c *gin.Context) {

}

func DeleteForum(c *gin.Context) {
	forum, _ := strconv.ParseUint(c.Param("forum"), 10, 64)
	forumService.Delete(forum)
	c.JSON(http.StatusOK, res.JsonSuccess())
}

func AllForum(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 15
	}
	offset := (page - 1) * limit
	c.JSON(http.StatusOK, res.JsonData(render.CreateAllForum(forumService.All(limit, offset))))
}
