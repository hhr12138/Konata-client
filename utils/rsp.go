package utils

import (
	"errors"
	"fmt"
	"github.com/hhr12138/Konata-client/consts"
	"strconv"
	"strings"
)

func ReadIntReply(line string) (int64, error) {
	switch line[0] {
	case consts.RESPError:
		return 0, ParseErrorReply(line)
	case consts.RESPInteger:
		return parseInt(line[1:], 10, 64)
	default:
		return 0, fmt.Errorf("redis: can't parse int reply: %.100q", line)
	}
}

func ParseErrorReply(line string) error {
	return errors.New(line[1:])
}

func parseInt(b string, base int, bitSize int) (int64, error) {
	return strconv.ParseInt(b[:len(b)-4], base, bitSize)
}

func ReadBytesReply(line string) ([]byte, error) {

	switch line[0] {
	case consts.RESPError:
		return nil, ParseErrorReply(line)
	case consts.RESPBulkString:
		return readTmpBytesValue(line)
	case consts.RESPString:
		return parseStatusValue(line), nil
	default:
		return nil, fmt.Errorf("redis: can't parse string reply: %.100q", line)
	}
}

func readTmpBytesValue(line string) ([]byte, error) {
	if isNilReply(line) {
		return nil, nil
	}

	replyLen, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, fmt.Errorf("read reply timeout")
	}
	startPos := 1 + len(strconv.FormatInt(int64(replyLen), 10)) + len("\r\n") + 1

	return []byte(line[startPos : startPos+replyLen]), nil
}

func isNilReply(b string) bool {
	return len(b) == 3 &&
		(b[0] == consts.RESPBulkString || b[0] == consts.RESPArrays) &&
		b[1] == '-' && b[2] == '1'
}

func ReadFloatReply(line string) (float64, error) {
	var (
		bs  []byte
		err error
	)
	switch line[0] {
	case consts.RESPError:
		err = ParseErrorReply(line)
	case consts.RESPBulkString:
		bs, err = readTmpBytesValue(line)
	case consts.RESPString:
		bs = parseStatusValue(line)
	default:
		bs, err = nil, fmt.Errorf("redis: can't parse string reply: %.100q", line)
	}
	if err != nil {
		return 0, err
	}
	return parseFloat(bs, 64)
}

func parseStatusValue(line string) []byte {
	return []byte(line[1 : len(line)-4])
}

func parseUint(b []byte, base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(string(b), base, bitSize)
}

func parseFloat(b []byte, bitSize int) (float64, error) {
	return strconv.ParseFloat(string(b), bitSize)
}

func ReadArrayReply(val string, parse func(val []string, n int64) (interface{}, error)) (interface{}, error) {
	switch val[0] {
	case consts.RESPError:
		return nil, ParseErrorReply(val)
	case consts.RESPArrays:
		// 分割数组，得到每一个元素
		vals := strings.Split(val, "\\r\\n")
		n, err := parseArrayLen(vals[0])
		if err != nil {
			return nil, err
		}
		// 末尾待遇/r/n的话去掉最后的空元素
		if len(vals[len(vals)-1]) == 0 {
			vals = vals[:len(vals)-1]
		}
		return parse(vals[1:], n)
	default:
		return nil, fmt.Errorf("can't parse array reply: %.100q", val)
	}
}

func ReadArrayParse(vals []string, n int64) (interface{}, error) {
	if len(vals) != int(n) {
		return nil, fmt.Errorf("array cnt err,target cnt=%v,cnt=%v", n, len(vals))
	}
	rsp := make([]interface{}, n)
	for _, val := range vals {
		reply, err := ReadReply(val)
		if err != nil {
			return nil, err
		}
		rsp = append(rsp, reply)
	}
	return rsp, nil
}

func ReadMapParse(vals []string, n int64) (interface{}, error) {
	if len(vals) != int(n) {
		return nil, fmt.Errorf("array cnt err,target cnt=%v,cnt=%v", n, len(vals))
	}
	if n%2 != 0 {
		return nil, fmt.Errorf("array cnt err, map cnt must be even")
	}
	// map对应的array里两行对应一个kv对
	rsp := make(map[string]interface{}, n/2)
	for i := 0; i < len(vals); i += 2 {
		key, err := ReadStringReply(vals[i])
		if err != nil {
			return nil, err
		}
		value, err := ReadReply(vals[i+1])
		if err != nil {
			return nil, err
		}
		rsp[key] = value
	}
	return rsp, nil
}

func parseArrayLen(line string) (int64, error) {
	if isNilReply(line) {
		return 0, nil
	}
	return parseInt(line[1:], 10, 64)
}

func ReadStringReply(line string) (string, error) {
	var (
		bs  []byte
		err error
	)
	switch line[0] {
	case consts.RESPError:
		bs, err = nil, ParseErrorReply(line)
	case consts.RESPBulkString:
		bs, err = readTmpBytesValue(line)
	case consts.RESPString:
		bs, err = parseStatusValue(line), nil
	default:
		bs, err = nil, fmt.Errorf("redis: can't parse string reply: %.100q", line)
	}
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ReadBoolReply(line string) (interface{}, error) {

	switch line[0] {
	case consts.RESPError:
		return nil, ParseErrorReply(line)
	case consts.RESPString:
		return parseStatusValue(line), nil
	case consts.RESPInteger:
		return parseInt(line[1:], 10, 64)
	case consts.RESPBulkString:
		return readTmpBytesValue(line)
	}
	return nil, fmt.Errorf("redis: can't parse %.100q", line)
}

func ReadReply(msg string) (interface{}, error) {

	switch msg[0] {
	case consts.RESPString:
		return ReadStringReply(msg)
	case consts.RESPInteger:
		return ReadIntReply(msg)
	case consts.RESPBulkString:
		return readTmpBytesValue(msg)
	case consts.RESPArrays:
		return ReadArrayReply(msg, ReadArrayParse)
	}
	return nil, fmt.Errorf("redis: can't parse %.100q", msg)
}
