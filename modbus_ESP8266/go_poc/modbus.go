package main

import (
	"os"

	// "github.com/goburrow/modbus"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	modbusAddress := os.Getenv("MODBUS_ADDRESS")
	modbusPort := os.Getenv("MODBUS_PORT")

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Starting Modbus TCP / IP client:")
	logrus.Debug("address: ", modbusAddress, ", port: ", modbusPort)

	// client := modbus.TCPClient(modbusAddress+":"+modbusPort)
	
	logrus.Debug("Writting data on a register:")
	// results, err := client.WriteSingleRegister(address, values)
	

}
