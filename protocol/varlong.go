package protocol

import (
	"errors"
	"net"
)

func ReadVarLong(conn net.Conn) (result int64, err error) {
	data := make([]byte, 1)
	_, err = conn.Read(data)

	if err != nil {
		return
	}

	return DecodeVarLong(data)
}

func DecodeVarLong(data []byte) (result int64, err error) {
	var numRead int32 = 0
	var read byte
	var value int64

	for {
		if numRead > 8 {
			err = errors.New("VarLong too big")
			return
		}

		read = data[numRead]
		value = int64(read & 0b01111111)
		result |= value << (7 * numRead)

		numRead++
		if read&0b10000000 == 0 {
			break
		}
	}

	return
}

func WriteVarLong(conn net.Conn, value int64) (err error) {
	data := EncodeVarLong(value)
	_, err = conn.Write(data)

	return
}

func EncodeVarLong(value int64) (data []byte) {
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
