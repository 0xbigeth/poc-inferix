package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	deviceId := uuid.New().String()
	for {
		_send_data(deviceId)
		time.Sleep(1 * time.Second)
	}

}

type SamplePayload struct {
	DeviceId string     `json:"device_id"`
	Data     SampleData `json:"data"`
}

type SampleData struct {
	ProducedImageBytes string `json:"produced_img_bytes"`
}

func _send_data(deviceId string) {
	url := "https://devnet-prod-api.w3bstream.com/event/eth_0x152c78bd6ceb6f58d58ab0e1168fa9b948f3dcc0_poc_inferix"

	method := "POST"

	payload := SamplePayload{
		DeviceId: deviceId,
		Data: SampleData{
			ProducedImageBytes: "sample bytes",
		},
	}
	payloadString, _ := json.Marshal(payload)
	fmt.Printf("send data %v\n", string(payloadString))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadString)))
	req.Header.Add("Authorization", "w3b_MV8xNzAxMzU3NTE1X2FaNkBIb3lbT28lbA")
	req.Header.Add("Content-Type", "application/octet-stream")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
