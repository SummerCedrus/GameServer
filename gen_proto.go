package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	IN_CMD_FILE_PATH = "protocol/cmd.txt"
	OUT_CMD_FILE_PATH = "protocol/cmd.go"
	IN_CONST_FILE_PATH = "protocol/const.txt"
	OUT_CONST_FILE_PATH = "protocol/const.go"
	IN_ERROR_FILE_PATH = "protocol/error.txt"
	OUT_ERROR_FILE_PATH = "protocol/error.go"
	DECODE_FILE_PATH = "protocol/decode.go"
)
func main(){
	err := parse_cmd()
	if nil != err && io.EOF != err{
		fmt.Printf("parse cmd error %v", err)
	}
	err = parse_common(IN_CONST_FILE_PATH, OUT_CONST_FILE_PATH)
	if nil != err && io.EOF != err{
		fmt.Printf("parse const error %v", err)
	}
	err = parse_common(IN_ERROR_FILE_PATH, OUT_ERROR_FILE_PATH)
	if nil != err && io.EOF != err{
		fmt.Printf("parse errorcode error %v", err)
	}
}


func parse_cmd() error{
	fIn, err := os.Open(IN_CMD_FILE_PATH)
	if nil != err{
		return err
	}
	inBuf := bufio.NewReader(fIn)
	fOut,err := os.OpenFile(OUT_CMD_FILE_PATH, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if nil != err{
		return err
	}
	fDecode, err := os.OpenFile(DECODE_FILE_PATH, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if nil != err {
		return err
	}
	defer func() {
		fOut.WriteString(")")
		fDecode.WriteString("default:\n" +
			"fmt.Printf(\"Error Cmd [%d]\", Cmd)\n" +
			"return nil, errors.New(\"Error Cmd\")\n" +
			"}}\n")

	}()

	fDecode.WriteString("package protocol\n" +
		"import (\n" +
		"\"github.com/golang/protobuf/proto\"\n" +
		"\"fmt\"\n" +
		"\"errors\"\n" +
		")\n")

	fDecode.WriteString("func ReflectMessage(Cmd uint32) (proto.Message, error){\n" +
		"switch Cmd {\n")
	fOut.WriteString("package protocol\n")
	fOut.WriteString("const (\n")

	//
	//	return &ItemInfo{}, nil


	for {
		line, err := inBuf.ReadString('\n')
		if nil != err{
			return err
		}
		if line == "" || line == "\n"||line == "\r\n"{
			continue
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//"){
			continue
		}
		vStr := strings.Split(line, "#")

		if len(vStr) != 4{
			return errors.New(fmt.Sprintf("Invalid format [%s]\n", line))
		}
		fOut.WriteString(fmt.Sprintf("%s = %s //%s\n",vStr[1], vStr[0], vStr[3]))
		fDecode.WriteString(fmt.Sprintf("case %s:\n", vStr[1]))
		fDecode.WriteString(fmt.Sprintf("return &%s{}, nil\n",vStr[2]))
	}

	return nil
}

func parse_common(inFile string, outFile string) error{
	fIn, err := os.Open(inFile)
	if nil != err{
		return err
	}
	inBuf := bufio.NewReader(fIn)
	fOut,err := os.OpenFile(outFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if nil != err{
		return err
	}
	defer fOut.WriteString(")")
	fOut.WriteString("package protocol\n")
	fOut.WriteString("const (\n")
	for {
		line, err := inBuf.ReadString('\n')
		if nil != err{
			return err
		}
		if "" == line || "\n" == line || "\r\n" == line{
			continue
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//"){
			continue
		}
		vStr := strings.Split(line, "#")
		if len(vStr) != 3{
			return errors.New(fmt.Sprintf("Invalid format %s\n", line))
		}
		fOut.WriteString(fmt.Sprintf("%s = %s //%s\n",vStr[1], vStr[0], vStr[2]))


	}

	return nil
}