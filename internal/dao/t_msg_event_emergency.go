// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GeoDigitalMap-messageRelayService/internal/dao/internal"
)

// internalTMsgEventEmergencyDao is an internal type for wrapping the internal DAO implementation.
type internalTMsgEventEmergencyDao = *internal.TMsgEventEmergencyDao

// tMsgEventEmergencyDao is the data access object for the table t_msg_event_emergency.
// You can define custom methods on it to extend its functionality as needed.
type tMsgEventEmergencyDao struct {
	internalTMsgEventEmergencyDao
}

var (
	// TMsgEventEmergency is a globally accessible object for table t_msg_event_emergency operations.
	TMsgEventEmergency = tMsgEventEmergencyDao{
		internal.NewTMsgEventEmergencyDao(),
	}
)

// Add your custom methods and functionality below.
