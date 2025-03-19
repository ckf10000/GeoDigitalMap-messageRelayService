// Package utils
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     config.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 16:10:46
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package utils

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"net"
)

func GetServiceEnableStatus(ctx context.Context, serviceName string) bool {
	isEnableKey := fmt.Sprintf("server.%s.isEnable", serviceName)
	gvar, err := g.Cfg().Get(ctx, isEnableKey)
	if err != nil {
		return true
	}
	return gvar.Bool()
}

func CheckFederateLocalIP(ctx context.Context) bool {
	loaclIP := GetFederateLocalIP(ctx)
	if len(loaclIP) > 0 {
		currentHostAddress := GetCurrentHostAddress(ctx)
		return Contains(currentHostAddress, loaclIP)
	}
	return false

}

func GetCurrentHostAddress(ctx context.Context) []string {
	address := make([]string, 0, consts.FederatePeerCount)
	interfaces, err := net.Interfaces()
	if err != nil {
		g.Log(consts.FederateLogger).Errorf(ctx, "get host network interface address failed: %+v", err)
		return address
	}

	for _, iface := range interfaces {
		addrs, err1 := iface.Addrs()
		if err1 != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() { // 过滤掉 127.0.0.1
				ip := ipNet.IP.String()
				g.Log(consts.FederateLogger).Infof(ctx, "get host network interface: %s, ip: %s", iface.Name, ip)
				address = append(address, ip)
			}
		}
	}
	return address
}

func GetFederateLocalIP(ctx context.Context) (localIP string) {
	localIPKey := fmt.Sprintf("server.%s.localIP", consts.FederateService)
	gvar1, err := g.Cfg().Get(ctx, localIPKey)
	if err != nil {
		str := fmt.Sprintf("%s does not have 'LocalIP' configured, eg: 'LocalIP: 192.168.1.1'", consts.FederateService)
		g.Log(consts.FederateLogger).Info(ctx, str)
		return
	}
	localIP = gvar1.String()
	//if err = g.Validator().Data(localIP).Rules("ip").Run(ctx); err != nil {
	//	g.Log(consts.FederateLogger).Errorf(ctx, "localIP: %s Format verification failed", localIP)
	//	return "", err
	//}
	return
}

func GetFederateRemoteAddress(ctx context.Context) (hostAddrs []*dto.HostAddress) {
	federateRemoteAddress := fmt.Sprintf("server.%s.federateRemoteAddress", consts.FederateService)
	gvar1, err := g.Cfg().Get(ctx, federateRemoteAddress)
	if err != nil {
		return
	}
	addressAll := gvar1.Strings()
	for _, v := range addressAll {
		if ok := gstr.ContainsAny(v, ":"); ok {
			vSlice := gstr.SplitAndTrim(v, ":")
			if len(vSlice) == 2 {
				hostAddress := &dto.HostAddress{
					IP:   vSlice[0],
					Port: gconv.Uint32(vSlice[1]),
				}
				if err = g.Validator().Data(hostAddress).Run(ctx); err != nil {
					g.Log(consts.FederateLogger).Error(ctx, err)
					continue
				}
				hostAddrs = append(hostAddrs, hostAddress)
			} else {
				g.Log(consts.FederateLogger).Errorf(ctx, "Invalid address format: %s", v)
			}
		} else {
			g.Log(consts.FederateLogger).Errorf(ctx, "Address is missing port: %s", v)
		}
	}
	return
}
