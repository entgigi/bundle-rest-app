package controllers

import (
	"net/http"

	"github.com/entgigi/bundle-rest-app/clients"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/rest"
)

type BundleDto struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Titile        string `json:"title"`
	Icon          string `json:"icon"`
	SignatureInfo string `json:"signatureInfo"`
}

type BundleCtrl struct {
	apiClient *clients.BundleV1Alpha1Client
}

func NewBundleCtrl(config *rest.Config) (*BundleCtrl, error) {
	s, err := clients.NewForConfig(config)
	return &BundleCtrl{s}, err
}

func (bc *BundleCtrl) ListBundles(ctx *gin.Context) {

	bundles := []BundleDto{{Code: "code", Name: "name"}}

	ctx.JSON(http.StatusOK, gin.H{"data": bundles})
}

func (bc *BundleCtrl) GetBundle(ctx *gin.Context) {

	code := ctx.Param("code")
	bundle := BundleDto{Code: code, Name: "myname"}

	ctx.JSON(http.StatusOK, gin.H{"data": bundle})
}
