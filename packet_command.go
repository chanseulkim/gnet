package gnet

const (
	TYPE_COMMAND_MOVE  byte = 0
	TYPE_COMMAND_ENTER byte = 1
	TYPE_COMMAND_LEAVE byte = 2
	//don't use COMMAND_END
	TYPE_COMMAND_END byte = 8
)

type CmdPacket struct {
	header       *GHeader
	command_type byte
	buff         []byte
	buff_len     int32
}

func NewCmdPacket(command_type byte, data []byte, data_size int32) *CmdPacket {
	return &CmdPacket{
		header:       &GHeader{header_type: TYPE_HEADER_CMD},
		command_type: command_type,
		buff:         data,
		buff_len:     data_size,
	}
}

func (p *CmdPacket) Write(data []byte, data_size int32) {
	copy(p.buff[p.buff_len:], data)
	p.buff_len += data_size
}
