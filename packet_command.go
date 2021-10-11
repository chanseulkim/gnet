package gnet

// type CmdType byte

// const (
// 	TYPE_COMMAND_START = CmdType(100 + iota)
// 	TYPE_COMMAND_MOVE  = CmdType(iota)
// 	TYPE_COMMAND_ENTER = CmdType(iota)
// 	TYPE_COMMAND_LEAVE = CmdType(iota)
// 	//don't use COMMAND_END
// 	TYPE_COMMAND_END = CmdType(127)
// )

// type CmdPacket struct {
// 	header       *GHeader
// 	command_type byte
// 	buff         []byte
// 	buff_len     int32
// }

// func NewCmdPacket(command_type byte, data []byte, data_size int32) *CmdPacket {
// 	return &CmdPacket{
// 		header:       NewGHeader(TYPE_HEADER_CMD),
// 		command_type: command_type,
// 		buff:         data,
// 		buff_len:     data_size,
// 	}
// }

// // func (p *CmdPacket) Write(data []byte, data_size int32) {
// // 	copy(p.buff[p.buff_len:], data)
// // 	p.buff_len += data_size
// // }
