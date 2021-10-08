package gnet

const (
	TYPE_HEADERTYPE_CMD  byte = 0
	TYPE_HEADERTYPE_SYNC byte = 1
	TYPE_HEADERTYPE_MSG  byte = 2
	//don't use TYPE_HEADERTYPE_END
	TYPE_HEADERTYPE_END byte = 8
)
const (
	HEADER_LENGTH = 2
	//MAX_NAME_LENGTH int = 20
)

// |header type|command type|
// |  0 0 0 0  |  0 0 0 0  |
type GHeader struct {
	header_type byte
	buffer      [HEADER_LENGTH]byte
}

func (h *GHeader) SetHeader(header_type byte) bool {
	if header_type >= TYPE_HEADERTYPE_END {
		return false
	}
	h.header_type = header_type
	h.buffer[0] = header_type
	return true
}
func (h *GHeader) GetBytes() [HEADER_LENGTH]byte {
	return h.buffer
}
