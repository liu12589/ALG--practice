package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

// Encode 编码，开辟一个缓冲区，将数据填入
func Encode(massage string) ([]byte, error) {
	massageLength := int32(len(massage))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, massageLength)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(massage))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码，前四位为该消息的长度，并返回消息
func Decode(reader *bufio.Reader) (string, error) {
	peek, err := reader.Peek(4)
	if err != nil {
		return "", err
	}
	lengthBuff := bytes.NewBuffer(peek)

	var length int32
	err = binary.Read(lengthBuff, binary.LittleEndian, &length)

	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < 4+length {
		return "", fmt.Errorf("err")
	}
	p := make([]byte, int(4+length))
	_, err = reader.Read(p)
	if err != nil {
		return "", err
	}
	return string(p[4:]), nil
}
