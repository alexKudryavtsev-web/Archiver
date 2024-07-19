package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk
type HexChunk string
type HexChunks []HexChunk

var chunkSize = 8

func (chunk BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(chunk), 2, chunkSize)

	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func (chunks BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(chunks))
	for _, chunk := range chunks {
		res = append(res, chunk.ToHex())
	}
	return res
}

func (bcs HexChunks) ToString() string {
	const sep = " "

	switch len(bcs) {
	case 0:
		return ""
	case 1:
		return string(bcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(bcs[0]))

	for _, chunk := range bcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func Encode(str string) string {
	str = prepateText(str)
	chunks := splitByChunks(encodeBin(str), chunkSize)

	return chunks.ToHex().ToString()
}

// prepareText prepares text to be fit for encode:
// changes upper case latters to: ! + lower case letter
func prepateText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// splitByChunks splits binary string to chunks with given size
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()

		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

// encodeBin encodes str into binary codes string without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

// bin encode rune to bytes
func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]

	if !ok {
		panic("unknown character " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		'e': "101",
		't': "1001",
		'o': "10001",
		'n': "10000",
		'a': "011",
		's': "0101",
		'i': "01001",
		'r': "01000",
		'h': "0011",
		'd': "00101",
		'l': "001001",
		'!': "001000",
		'u': "00011",
		'c': "000101",
		'f': "000100",
		'm': "000011",
		'p': "0000101",
		'g': "0000100",
		'w': "0000011",
		'b': "0000010",
		'y': "0000001",
		'v': "00000001",
		'j': "000000001",
		'k': "0000000001",
		'x': "00000000001",
		'q': "000000000001",
		'z': "000000000000",
	}
}
