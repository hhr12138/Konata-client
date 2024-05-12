package utils

import (
	"encoding"
	"fmt"
	"github.com/google/uuid"
	"github.com/hhr12138/Konata-client/consts"
	"hash/fnv"
	"strconv"
	"strings"
)

// 后续优化成雪花算法
func GetReqId() string {
	return uuid.NewString()
}

func GetAddrIdx(key string, length int) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % length
}

func BuildMsg(args ...interface{}) (string, error) {
	var b strings.Builder
	b.WriteRune(consts.RESPArrays)
	b.WriteString(strconv.FormatInt(int64(len(args)), 10))
	b.WriteRune('\r')
	b.WriteRune('\n')

	for _, arg := range args {
		if err := appendVal(&b, arg); err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func appendVal(b *strings.Builder, val interface{}) error {
	switch v := val.(type) {
	case nil:
		appendString(b, "")
	case string:
		appendString(b, v)
	case []byte:
		appendBytes(b, v)
	case int:
		appendString(b, formatInt(int64(v)))
	case int8:
		appendString(b, formatInt(int64(v)))
	case int16:
		appendString(b, formatInt(int64(v)))
	case int32:
		appendString(b, formatInt(int64(v)))
	case int64:
		appendString(b, formatInt(v))
	case uint:
		appendString(b, formatUint(uint64(v)))
	case uint8:
		appendString(b, formatUint(uint64(v)))
	case uint16:
		appendString(b, formatUint(uint64(v)))
	case uint32:
		appendString(b, formatUint(uint64(v)))
	case uint64:
		appendString(b, formatUint(v))
	case float32:
		appendString(b, formatFloat(float64(v)))
	case float64:
		appendString(b, formatFloat(v))
	case bool:
		if v {
			appendString(b, "1")
		} else {
			appendString(b, "0")
		}
	default:
		if bm, ok := val.(encoding.BinaryMarshaler); ok {
			bb, err := bm.MarshalBinary()
			if err != nil {
				return err
			}
			appendBytes(b, bb)
		} else {
			return fmt.Errorf(
				"RESP: can't marshal %T (Maybe you are passing a slice without using ... to expand arguments by mistake? Or just consider implementing encoding.BinaryMarshaler)", val)
		}
	}
	return nil
}

func appendString(b *strings.Builder, s string) {
	b.WriteRune(consts.RESPString)
	b.WriteString(strconv.FormatInt(int64(len(s)), 10))
	b.WriteRune('\r')
	b.WriteRune('\n')
	b.WriteString(s)
	b.WriteRune('\r')
	b.WriteRune('\n')
}

func appendBytes(b *strings.Builder, p []byte) {
	b.WriteRune(consts.RESPString)
	b.WriteString(strconv.FormatInt(int64(len(p)), 10))
	b.WriteRune('\r')
	b.WriteRune('\n')
	b.Write(p)
	b.WriteRune('\r')
	b.WriteRune('\n')
}

func formatInt(n int64) string {
	return strconv.FormatInt(n, 10)
}

func formatUint(u uint64) string {
	return strconv.FormatUint(u, 10)
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
