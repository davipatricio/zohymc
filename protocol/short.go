package protocol

func DecodeShort(data []byte) int16 {
	return int16(data[0])<<8 | int16(data[1])
}
