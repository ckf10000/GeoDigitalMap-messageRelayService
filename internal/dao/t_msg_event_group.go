// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GeoDigitalMap-messageRelayService/internal/dao/internal"
)

// internalTMsgEventGroupDao is an internal type for wrapping the internal DAO implementation.
type internalTMsgEventGroupDao = *internal.TMsgEventGroupDao

// tMsgEventGroupDao is the data access object for the table t_msg_event_group.
// You can define custom methods on it to extend its functionality as needed.
type tMsgEventGroupDao struct {
	internalTMsgEventGroupDao
}

var (
	// TMsgEventGroup is a globally accessible object for table t_msg_event_group operations.
	TMsgEventGroup = tMsgEventGroupDao{
		internal.NewTMsgEventGroupDao(),
	}
)

// Add your custom methods and functionality below.
