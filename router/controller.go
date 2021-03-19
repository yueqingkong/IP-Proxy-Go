package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"net/http"
)

var lastIpId int64

func IP(context *gin.Context) {
	ip, err := db.XORM().GetIPById(lastIpId)

	if err != nil || ip.Address == "" {
		ip, err = db.XORM().GetIPById(0)
		lastIpId = ip.Id

		context.JSON(http.StatusOK, gin.H{"status": "200", "body": ip})
	} else {
		lastIpId = ip.Id
		context.JSON(http.StatusOK, gin.H{"status": "200", "body": ip})
	}
}
