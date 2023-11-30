package main

import (
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func start(rid uint32) int32 {
	return handle_data(rid)
}

//export handle_result
func handle_data(rid uint32) int32 {
	log.Log(fmt.Sprintf("start rid: %d", rid))

	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}
	res := string(message)
	log.Log(fmt.Sprintf("message received: %v", res))
	return 0
}

func handle_device_registered(deviceId string) {
	log.Log(fmt.Sprintf("handle_device_registered: deviceId %v", deviceId))
}

func handle_device_binding(deviceId, owner string) {
	log.Log(fmt.Sprintf("handle_device_binding: deviceId %v, owner %v", deviceId, owner))
}
