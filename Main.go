package main

import (
	"fmt"
	"log"

	"./modbusclient"
)

func main() {
	modbusDevice, err := modbusclient.NewModbusTCP("127.0.0.1:502", 1) // 모드버스 TCP 서버의 주소로 변경해주세요.
	if err != nil {
		log.Fatal("모드버스 TCP 서버에 연결할 수 없습니다:", err)
	}
	defer modbusDevice.Close()

	// 데이터 읽기 예제
	results, err := modbusDevice.ReadHoldingRegisters(0, 10) // 시작 주소와 읽을 레지스터 수
	if err != nil {
		log.Fatal("데이터를 읽을 수 없습니다:", err)
	}
	fmt.Println("읽은 데이터:", results)

	// 데이터 쓰기 예제
	data := []byte{0x01, 0x02, 0x03} // 쓸 데이터
	address := uint16(0)             // 쓸 주소
	_, err = modbusDevice.WriteMultipleRegisters(address, data)
	if err != nil {
		log.Fatal("데이터를 쓸 수 없습니다:", err)
	}
	fmt.Println("데이터를 성공적으로 썼습니다.")
}
