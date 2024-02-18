package handler

import (
	"net/http"
	"reblog/internal/query"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		删除文章
//	@Description	根据slug删除文章
//	@Tags			文章管理
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string		true	"文章的slug"
//	@Success		200		{object}	common.Resp	"删除成功"
//	@Failure		400		{object}	common.Resp	"缺少必要参数"
//	@Failure		404		{object}	common.Resp	"未知的文章"
//	@Security		ApiKeyAuth
//	@Router			/articles/{slug} [delete]
func ArticlesDelete(router fiber.Router) {
	router.Delete("/:slug", func(c fiber.Ctx) error {
		a := query.Article

		slug := c.Params("slug")

		if common.CheckEmpty(slug) {
			return common.RespMissingParameters(c)
		}

		article, err := a.Where(a.Slug.Eq(slug)).First()

		if article == nil {
			return common.RespFail(c, http.StatusNotFound, "未知的文章", nil)
		}

		if err != nil {
			return common.RespServerError(c, err)
		}

		_, err = a.Delete(article)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "删除成功", nil)

	}, common.Auth())
}
