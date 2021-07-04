package main

import (
	"encoding/json"
	"fmt"
)

type RebotAction struct {
	webHook string
	data    interface{}
}

func NewRebot(webHook string, data interface{}) *RebotAction {
	return &RebotAction{
		webHook: webHook,
		data:    data,
	}
}

func (r *RebotAction) DoAction() error {
	jsonStr, err := json.Marshal(r.data)
	if err != nil {
		return err
	}
	fmt.Println("jsonStr", string(jsonStr))
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

type Text struct {
	Content             string   `json:"context"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type TextMessage struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}

func NewTextMessage(content string, mentionedList []string, mentionedMobileList []string) *TextMessage {
	return &TextMessage{
		Msgtype: "text",
		Text: Text{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}
}

type Markdown struct {
	Content string `json:"context"`
}

type MarkdownMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

func NewMarkdownMessage(content string) *MarkdownMessage {
	return &MarkdownMessage{
		Msgtype: "markdown",
		Markdown: Markdown{
			Content: content,
		},
	}
}
