// Package utils
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     validator.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 01:54:07
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package utils

// Contains 判断字符串 a 是否在列表 b 中
func Contains(b []string, a string) bool {
	for _, item := range b {
		if item == a {
			return true
		}
	}
	return false
}
