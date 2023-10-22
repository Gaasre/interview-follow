package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"interview-follow/config"
	"log"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type ParsedEmail struct {
	Company   string    `json:"company"`
	Title     string    `json:"title"`
	Stage     string    `json:"stage"`
	Summary   string    `json:"summary"`
	Interview string    `json:"interview"`
	Type      string    `json:"type"`
	Date      time.Time `json:"date"`
	Links     []string  `json:"links"`
}

var openAIClient *openai.Client

func InitOpenAI() {
	openAIClient = openai.NewClient(config.Config("OPENAI_SECRET"))
}

func ParseEmail(from string, subject string, body string) ParsedEmail {
	resp, err := openAIClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Input: You will be given an email coming from a recruiter. The context of the email is an interview process.\nOutput:I want you to return a json object that consist of (company:the company name if found, title: the job title if found, stage: the upcoming hiring process stage and if it's a rejection email specify it, summary: the summary of the email in the form of a task, interview: whether or not the next step is a scheduled interview/assessment, type: if it's a scheduled interview/assessment specify the type of the interview/assessment, date: the date of the interview/assessment if it's scheduled, links: if there are any useful link for the interview in the email, only include actual/relevant links)",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("From:%s Subject:%s Body:%s", from, subject, body),
				},
			},
		},
	)

	if err != nil {
		log.Printf("OpenAI error: %v\n", err)
		return ParsedEmail{}
	}

	if len(resp.Choices) == 0 {
		return ParsedEmail{}
	}

	data := ParsedEmail{}
	json.Unmarshal([]byte(resp.Choices[0].Message.Content), &data)

	return data
}
