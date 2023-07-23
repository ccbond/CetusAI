package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ccbond/cetus-ai/internal/model"
	"github.com/ccbond/cetus-ai/internal/util/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (srv *Server) getPoolList(ctx *gin.Context) {
	url := "https://api-sui.cetus.zone/v2/sui/pools_info"
	resp, err := http.Get(url)
	srv.logger.Debug("getPoolList", zap.Any("url", url))
	if err != nil {
		api.ResponseErrors(ctx, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	srv.logger.Debug("get resp", zap.Any("body", body))
	if err != nil {
		api.ResponseErrors(ctx, err)
		return
	}
	var poolList model.Response
	if err := json.Unmarshal(body, &poolList); err != nil {
		api.ResponseErrors(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, poolList)
}
