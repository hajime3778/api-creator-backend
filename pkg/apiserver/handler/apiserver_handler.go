package handler

import (
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/usecase"

	"github.com/gin-gonic/gin"
)

// APIServerHandler APIServerAPIに対するリクエストハンドラ
type APIServerHandler struct {
	usecase usecase.APIServerUsecase
}

// NewAPIServerHandler APIServerHandlerを作成します
func NewAPIServerHandler(engine *gin.Engine, u usecase.APIServerUsecase) {
	handler := &APIServerHandler{
		usecase: u,
	}
	engine.Any("/*proxyPath", handler.RequestDocumentServer)
}

// RequestDocumentServer リクエスト情報からAPIServerを特定し、ドキュメントに対してCRUDします
func (h *APIServerHandler) RequestDocumentServer(c *gin.Context) {

	httpMethod := c.Request.Method
	// 最初の文字は/なので削除する
	url := c.Param("proxyPath")[1:]

	// buf := make([]byte, 2048)
	// num, _ := c.Request.Body.Read(buf)
	// body := string(buf[0:num])

	body, _ := c.GetRawData()
	log.Println(string(body))

	method, param, err := h.usecase.RequestDocumentServer(httpMethod, url, body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"param":  param,
		"method": method,
	})
}
