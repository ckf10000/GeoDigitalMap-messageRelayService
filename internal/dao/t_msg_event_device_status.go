// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GeoDigitalMap-messageRelayService/internal/dao/internal"
)

// internalTMsgEventDeviceStatusDao is an internal type for wrapping the internal DAO implementation.
type internalTMsgEventDeviceStatusDao = *internal.TMsgEventDeviceStatusDao

// tMsgEventDeviceStatusDao is the data access object for the table t_msg_event_device_status.
// You can define custom methods on it to extend its functionality as needed.
type tMsgEventDeviceStatusDao struct {
	internalTMsgEventDeviceStatusDao
}

var (
	// TMsgEventDeviceStatus is a globally accessible object for table t_msg_event_device_status operations.
	TMsgEventDeviceStatus = tMsgEventDeviceStatusDao{
		internal.NewTMsgEventDeviceStatusDao(),
	}
)

// Add your custom methods and functionality below.
