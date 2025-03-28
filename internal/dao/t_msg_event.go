// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GeoDigitalMap-messageRelayService/internal/dao/internal"
)

// internalTMsgEventDao is an internal type for wrapping the internal DAO implementation.
type internalTMsgEventDao = *internal.TMsgEventDao

// tMsgEventDao is the data access object for the table t_msg_event.
// You can define custom methods on it to extend its functionality as needed.
type tMsgEventDao struct {
	internalTMsgEventDao
}

var (
	// TMsgEvent is a globally accessible object for table t_msg_event operations.
	TMsgEvent = tMsgEventDao{
		internal.NewTMsgEventDao(),
	}
)

// Add your custom methods and functionality below.
