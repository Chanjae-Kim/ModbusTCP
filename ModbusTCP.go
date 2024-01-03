package modbusclient

import (
	"time"
)

type ModbusTCP struct {
	handler *modbus.TCPClientHandler
	client  modbus.Client
}

func NewModbusTCP(address string, slaveID uint8) (*ModbusTCP, error) {
	handler := modbus.NewTCPClientHandler(address)
	handler.Timeout = 10 * time.Second
	handler.SlaveId = slaveID

	err := handler.Connect()
	if err != nil {
		return nil, err
	}

	client := modbus.NewClient(handler)
	return &ModbusTCP{
		handler: handler,
		client:  client,
	}, nil
}

func (m *ModbusTCP) Close() error {
	return m.handler.Close()
}

func (m *ModbusTCP) ReadHoldingRegisters(startAddress, quantity uint16) ([]byte, error) {
	return m.client.ReadHoldingRegisters(startAddress, quantity)
}

func (m *ModbusTCP) WriteMultipleRegisters(address uint16, data []byte) (int, error) {
	return m.client.WriteMultipleRegisters(address, data)
}
