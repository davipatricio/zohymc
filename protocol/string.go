package protocol

func EncodeString(value string) (data []byte) {
	data = append(data, EncodeVarInt(int32(len(value)))...)
	data = append(data, []byte(value)...)

	return
}
