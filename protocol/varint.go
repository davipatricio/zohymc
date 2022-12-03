package protocol

import (
	"errors"
	"net"
)

func ReadVarInt(conn net.Conn) (result int32, err error) {
	data := make([]byte, 1)
	_, err = conn.Read(data)

	if err != nil {
		return
	}

	return DecodeVarInt(data)
}

func DecodeVarInt(data []byte) (result int32, err error) {
	var numRead int32 = 0
	var read byte
	var value int32

	for {
		if numRead > 4 {
			err = errors.New("VarInt too big")
			return
		}

		read = data[numRead]
		value = int32(read & 0b01111111)
		result |= value << (7 * numRead)

		numRead++
		if read&0b10000000 == 0 {
			break
		}
	}

	return
}

func WriteVarInt(conn net.Conn, value int32) (err error) {
	data := EncodeVarInt(value)
	_, err = conn.Write(data)

	return
}

func EncodeVarInt(value int32) (data []byte) {
	for {
		temp := byte(value & 0b01111111)
		value >>= 7
		if value != 0 {
			temp |= 0b10000000
		}

		data = append(data, temp)

		if value == 0 {
			break
		}
	}

	return
}
