package gnet

type SyncPacket struct {
	header      *GHeader
	object_type byte
	buff        []byte
	buff_len    int32
}

func NewSyncPacket(object_type byte, data []byte, data_size int32) *SyncPacket {
	return &SyncPacket{
		header:      &GHeader{header_type: TYPE_HEADERTYPE_SYNC},
		object_type: object_type,
		buff:        data,
		buff_len:    data_size,
	}
}
