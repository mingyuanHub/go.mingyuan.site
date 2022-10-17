package net

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

func GetClientIP(ctx *gin.Context) string {
	clientIP := ctx.Request.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(ctx.Request.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(ctx.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}