// Package cmd
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     ssl_socket_server.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-14 22:23:03
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gorilla/websocket"
	"math/big"
	"os"
	"time"
)

func GenerateCert(serverName string) (crtPath, keyPath string) {
	var ctx = gctx.New()
	crtPathConfig := fmt.Sprintf("server.%s.crtPath", serverName)
	keyPathConfig := fmt.Sprintf("server.%s.keyPath", serverName)
	gvar1, err1 := g.Cfg().Get(ctx, crtPathConfig)
	if err1 != nil {
		panic(fmt.Sprintf("%s does not have SSL certificate 'crtPath' configured, eg: 'crtPath: xx/xx/server.crt'", serverName))
	} else {
		crtPath = gvar1.String()
	}
	gvar2, err2 := g.Cfg().Get(ctx, keyPathConfig)
	if err2 != nil {
		panic(fmt.Sprintf("%s does not have SSL certificate 'keyPath' configured, eg: 'keyPath: xx/xx/server.key'", serverName))
	} else {
		keyPath = gvar2.String()
	}

	if !gfile.Exists(crtPath) || !gfile.Exists(keyPath) {
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)
		template := x509.Certificate{
			SerialNumber: big.NewInt(1),
			NotBefore:    time.Now(),
			NotAfter:     time.Now().Add(8760 * time.Hour),
		}
		derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)

		certOut, _ := os.Create(crtPath)
		err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
		if err != nil {
			str := fmt.Sprintf("SSL certificate file generation failed, %v", err)
			panic(str)
		}

		keyOut, _ := os.Create(keyPath)
		err = pem.Encode(keyOut, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		})
		if err != nil {
			str := fmt.Sprintf("SSL RSA private key file generation failed, %v", err)
			panic(str)
		}
	}
	return
}

func SSLSocketServer(ctx context.Context) *ghttp.Server {
	crtPath, keyPath := GenerateCert(consts.SSLSocketService)
	ser := g.Server(consts.SSLSocketService)
	ser.SetLogger(g.Log(consts.SSLSocketService))
	// Bind WebSocket handler to / endpoint
	ser.BindHandler("/", func(r *ghttp.Request) {
		ws, err := WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			r.Response.Write(err.Error())
			return
		}

		defer func(ws *websocket.Conn) {
			err = ws.Close()
			if err != nil {
				g.Log(consts.SSLSocketService).Error(ctx, err)
			}
		}(ws)

		for {
			msgType, msg, err1 := ws.ReadMessage()
			if err1 != nil {
				break
			}
			g.Log(consts.SSLSocketService).Infof(ctx, "received message: %s", msg)
			if err = ws.WriteMessage(msgType, msg); err != nil {
				break
			}
		}
		g.Log(consts.SSLSocketService).Info(ctx, "websocket connection closed")
	})
	ser.EnableHTTPS(crtPath, keyPath)
	ser.SetGraceful(true)
	ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28443)
	return ser
}
