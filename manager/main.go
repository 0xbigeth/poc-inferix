package main

import (
	"encoding/json"
	"fmt"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

type w3bstreamPayload struct {
	DeviceId string      `json:"deviceId"`
	Type     string      `json:"type"`
	Payload  interface{} `json:"payload"`
}

//export start
func start(rid uint32) int32 {
	return handle_data(rid)
}

//export handle_result
func handle_data(rid uint32) int32 {
	log.Log(fmt.Sprintf("starting handle data rid: %d", rid))

	payloadBytes, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: cannot get data" + err.Error())
		return -1
	}
	log.Log(fmt.Sprintf("received payload: %v", string(payloadBytes)))

	payload := w3bstreamPayload{}
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Log("error: cannot parsing payload")
		return -1
	}
	return 0
}

func handle_device_registered(deviceId string) {
	log.Log(fmt.Sprintf("handle_device_registered: deviceId %v", deviceId))
}

func handle_device_binding(deviceId, owner string) {
	log.Log(fmt.Sprintf("handle_device_binding: deviceId %v, owner %v", deviceId, owner))
}
