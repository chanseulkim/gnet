package gnet

const (
	TYPE_HEADER_CMD  byte = 0
	TYPE_HEADER_SYNC byte = 1
	TYPE_HEADER_MSG  byte = 2
	//don't use TYPE_HEADER_END
	TYPE_HEADER_END byte = 8
)
const (
	HEADER_LENGTH = 2
	//MAX_NAME_LENGTH int = 20
)

// |header type|command type|
// |  0 0 0 0  |  0 0 0 0  |
type GHeader struct {
	//header_type byte
	buffer [HEADER_LENGTH]byte
}

func NewGHeader(header_type byte) *GHeader {
	if header_type >= TYPE_HEADER_END {
		return nil
	}
	h := &GHeader{}
	h.buffer[0] = header_type
	return h
}
func ParseHeader(buff []byte) *GHeader {
	htype := buff[0]
	return NewGHeader(htype)
}

func (h *GHeader) GetBytes() [HEADER_LENGTH]byte {
	return h.buffer
}
