package main

import (
	"encoding/json"
	"fmt"
)

type Action interface {
	DoAction() error
}

type RebotAction struct {
	webHook string
	data    interface{}
}

func (r *RebotAction) DoAction() error {
	jsonStr, err := json.Marshal(r.data)
	if err != nil {
		return err
	}
	fmt.Println("jsonStr", jsonStr)
	/*fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", r.webHook, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))*/

	return nil
}
