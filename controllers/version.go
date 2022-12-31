package controllers

import (
	"net/http"

	"github.com/entgigi/bundle-rest-app/version"
	"github.com/gin-gonic/gin"
)

type Version struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Release   string `json:"release"`
}

func GetVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"version": Version{version.BuildTime, version.Commit, version.Release}})
}
