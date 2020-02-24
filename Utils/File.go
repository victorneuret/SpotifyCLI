package Utils

//
//import (
//	"bytes"
//	"encoding/binary"
//	"fmt"
//	"golang.org/x/oauth2"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func WriteFile(fileName string, token *oauth2.Token) {
//	var buf bytes.Buffer
//
//	err := binary.Write(&buf, binary.BigEndian, token)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	f, err := os.Create(fileName)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	_, err = f.Write(buf.Bytes())
//	if err != nil {
//		fmt.Println(err)
//		f.Close()
//		return
//	}
//	err = f.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
//
//func ReadFile(fileName string) *oauth2.Token {
//	var buf bytes.Buffer
//	var finalData oauth2.Token
//
//	fileData, err := ioutil.ReadFile(fileName)
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//	fmt.Println(string(fileData))
//	buf = *bytes.NewBuffer(fileData)
//	err = binary.Read(&buf, binary.BigEndian, &finalData)
//	if err != nil {
//		fmt.Println("failed to Read:", err)
//		return nil
//	}
//	return &finalData
//}
