package modbusGo

import "fmt"

func (h *TCPHandler) ReadRegisters(slaveID, functionCode, address, quantity uint16) ([]byte, error) {
	// Формирование запроса
	request := make([]byte, 12)
	request[0] = 0 // Transaction ID
	request[1] = 0
	request[2] = 0 // Protocol ID
	request[3] = 0
	request[4] = 0 // Length
	request[5] = 6
	request[6] = byte(slaveID)      // Unit ID
	request[7] = byte(functionCode)
	request[8] = byte(address >> 8)
	request[9] = byte(address)
	request[10] = byte(quantity >> 8)
	request[11] = byte(quantity)

	_, err := h.conn.Write(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	response := make([]byte, 9+2*int(quantity))
	_, err = h.conn.Read(response)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return response[9:], nil
}
