package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BundleDto struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Titile        string `json:"title"`
	Icon          string `json:"icon"`
	SignatureInfo string `json:"signatureInfo"`
}

func ListBundles(ctx *gin.Context) {

	bundles := []BundleDto{{Code: "code", Name: "name"}}

	ctx.JSON(http.StatusOK, gin.H{"data": bundles})
}

func GetBundles(ctx *gin.Context) {

	code := ctx.Param("code")
	bundle := BundleDto{Code: code, Name: "myname"}

	ctx.JSON(http.StatusOK, gin.H{"data": bundle})
}
