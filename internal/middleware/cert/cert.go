// Package cert
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     cert.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-16 23:45:26
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cert

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"math/big"
	"os"
	"time"
)

func GenerateCert(ctx context.Context, serverName string) (crtPath, keyPath string) {
	crtPathConfig := fmt.Sprintf("server.%s.crtPath", serverName)
	keyPathConfig := fmt.Sprintf("server.%s.keyPath", serverName)
	gvar1, err := g.Cfg().Get(ctx, crtPathConfig)
	if err != nil {
		panic(fmt.Sprintf("%s does not have SSL certificate 'crtPath' configured, eg: 'crtPath: xx/xx/server.crt'", serverName))
	} else {
		crtPath = gvar1.String()
	}
	gvar2, err := g.Cfg().Get(ctx, keyPathConfig)
	if err != nil {
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
		err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
		if err != nil {
			str := fmt.Sprintf("SSL certificate file generation failed, %+v", err)
			panic(str)
		}

		keyOut, _ := os.Create(keyPath)
		err = pem.Encode(keyOut, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		})
		if err != nil {
			str := fmt.Sprintf("SSL RSA private key file generation failed, %+v", err)
			panic(str)
		}
	}
	return
}
