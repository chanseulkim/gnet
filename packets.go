package gnet

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const HEADER_LENGTH = 2

type HeaderType byte

const (
	TYPE_HEADER_NONE HeaderType = 0
	TYPE_HEADER_CMD  HeaderType = iota
	TYPE_HEADER_SYNC HeaderType = iota
	TYPE_HEADER_MSG  HeaderType = iota
)

type PacketType byte

const (
	TYPE_PACKET_WHOLE    PacketType = 5 + iota
	TYPE_PACKET_CONTINUE PacketType = 5 + iota
	TYPE_PACKET_END      PacketType = 5 + iota
)

type CommandType byte

const (
	TYPE_COMMAND_NONE  CommandType = 0
	TYPE_COMMAND_START CommandType = iota
	TYPE_COMMAND_MOVE  CommandType = iota
	TYPE_COMMAND_ENTER CommandType = iota
	TYPE_COMMAND_LEAVE CommandType = iota
)

// |header type|packet type |   command type    | data
// |  0 0 0 0  |  0 0 0 0   | 0 0 0 0 | 0 0 0 0 | 0 0 0 0 ...
type GPacket struct {
	HeaderType HeaderType
	PacketType PacketType
	Command    CommandType
	SyncInfo   byte

	//header_type byte
	header_buff [HEADER_LENGTH]byte
	data        []byte
	data_size   uint32
}

func (p *GPacket) GetBytes() []byte {
	return p.data
}
func (p *GPacket) GetBytesLength() uint32 {
	return p.data_size
}

func NewGPacket(header_type byte, pack_type byte, cmd_type byte, data []byte) *GPacket {
	p := &GPacket{}
	p.header_buff[0] = 0
	p.header_buff[0] |= header_type << 4
	p.header_buff[0] |= pack_type

	p.header_buff[1] = cmd_type
	p.data = data
	return p
}

func ParsePacketHeader(buff []byte) *GPacket {
	p := &GPacket{}
	p.HeaderType = HeaderType((buff[0] & 0xF0) >> 4)
	p.PacketType = PacketType(buff[0] & 0x0F)

	if p.HeaderType == TYPE_HEADER_CMD {
		p.Command = CommandType(buff[1])
	}
	p.data = buff[HEADER_LENGTH:]
	return p
}

func NewMovePacket(pack_type byte, user_id string, to *Vector2) *GPacket {
	p := &GPacket{}
	p.header_buff[0] = 0
	p.header_buff[0] |= byte(TYPE_HEADER_CMD) << 4
	p.header_buff[0] |= pack_type
	p.header_buff[1] = byte(TYPE_COMMAND_MOVE)

	p.data = []byte(user_id + ";" + v2Str(*to))
	return p
}

func ParseCommandData(data []byte) (string, Vector2) {
	str := string(data)
	var pos Vector2
	var user_id string
	for i := 0; ; i++ {
		token_pos := strings.Index(str, ";")
		if token_pos == -1 {
			return user_id, pos
		}
		switch i {
		case 0:
			user_id = string(str[0:token_pos])
		case 1: // pos
			pos, _ = posStr2V2(string(str[0:token_pos]))
		case 3:
		default:
			return user_id, pos
		}
		str = str[token_pos+1:]
	}
	return user_id, pos
}
func NewSyncPacket(pack_type PacketType, objects *list.List) *GPacket {
	p := &GPacket{}
	p.header_buff[0] = 0
	p.header_buff[0] |= byte(TYPE_HEADER_SYNC) << 4
	p.header_buff[0] |= byte(pack_type)
	p.header_buff[1] = byte(TYPE_COMMAND_NONE)
	for e := objects.Front(); e != nil; e = e.Next() {
		obj := e.Value.(*GObject)
		if obj != nil {
			data, data_size := obj.Serialize()
			// set data size
			size_data := make([]byte, unsafe.Sizeof(data_size))
			size_data[0] |= byte(data_size & 0xFF000000)
			size_data[1] |= byte(data_size & 0x00FF0000)
			size_data[2] |= byte(data_size & 0x0000FF00)
			size_data[3] |= byte(data_size & 0x000000FF)

			// set data
			p.data = append(p.data, size_data...)
			p.data = append(p.data, data...)
			p.data_size += uint32(unsafe.Sizeof(data_size)) + data_size
		}
	}
	return p
}
func ParseSyncData(data []byte) []*GObject {
	var objects []*GObject
	var obj *GObject
	var data_offset uint32 = 0
	var data_size uint32 = 0
	for {
		data_size |= uint32((data[data_offset] & 0xFF) << 8 * 3)
		data_size |= uint32((data[data_offset+1] & 0xFF) << 8 * 2)
		data_size |= uint32((data[data_offset+2] & 0xFF) << 8)
		data_size |= uint32(data[data_offset+3] & 0xFF)
		if data_size <= 0 {
			return objects
		}
		data_offset += 4
		objects = append(objects, obj.Deserialize(data[data_offset:]))
		data_offset += data_size
	}
	return objects
}

//////////////////////////////////////////////////////////////
// "(40, 40)" -> x:40, y:40 int Vector2
func posStr2V2(str string) (Vector2, error) {
	str = strings.Trim(str, "()")
	tok := ", "
	p := strings.Index(str, tok)
	if p == -1 {
		return Vector2{}, fmt.Errorf("invalid value " + str)
	}
	x, _ := strconv.ParseFloat(str[:p], 32)
	y, _ := strconv.ParseFloat(str[p+len(tok):], 32)
	v := Vector2{int(x), int(y)}
	return v, nil
}
func ToPosString(x int, y int) string {
	return "(" + strconv.Itoa(int(x)) + ", " + strconv.Itoa(int(y)) + ")"
}
func v2Str(v Vector2) string {
	return "(" + strconv.Itoa(int(v.X)) + ", " + strconv.Itoa(int(v.Y)) + ")"
}
