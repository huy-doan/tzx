package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/sqs"
)

func Map[T1, T2 any](s []T1, fn func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = fn(v)
	}
	return r
}

func ArrayToMap[T1 any, T2 comparable](v []T1, fnc func(T1) T2) map[T2]T1 {
	m := map[T2]T1{}
	for _, item := range v {
		m[fnc(item)] = item
	}
	return m
}

func Chunk[T any, U ~[]T](s U, size int) []U {
	if size <= 0 {
		panic("chunk size must be greater than 0")
	}
	num := len(s) / size
	if len(s)%size != 0 {
		num++
	}
	r := make([]U, 0, num)
	for i := range num {
		last := min((i+1)*size, len(s))
		r = append(r, s[i*size:last])
	}
	return r
}

func ToSQSMessageInput[T model.BaseSQSMessage](v T) (*sqs.SendMessageInput, error) {
	body := v.GetMessageBody()
	messageAttribute := v.GetMessageAttributes()

	result, err := json.Marshal(body)
	var messageBody string

	if err != nil {
		return nil, err
	}
	messageBody = string(result)

	messageAttributes := map[string]types.MessageAttributeValue{}

	messageAttributes["MessageType"] = types.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(string(messageAttribute.MessageType)),
	}

	if messageAttribute.Paymethod != nil {
		messageAttributes["Paymethod"] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(string(*messageAttribute.Paymethod)),
		}
	}

	return &sqs.SendMessageInput{
		MessageAttributes: messageAttributes,
		MessageBody:       &messageBody,
	}, nil
}

func ToSQSMessageBatchInput[T model.BaseSQSMessage](v []T) (*sqs.SendMessageBatchInput, error) {
	var entries []types.SendMessageBatchRequestEntry

	for _, item := range v {
		message, err := ToSQSMessageInput(item)
		if err != nil {
			return nil, err
		}
		entry := types.SendMessageBatchRequestEntry{
			Id:                aws.String(uuid.New().String()),
			MessageBody:       message.MessageBody,
			MessageAttributes: message.MessageAttributes,
		}

		entries = append(entries, entry)
	}

	messages := &sqs.SendMessageBatchInput{
		Entries: entries,
	}

	return messages, nil
}

func RedactedSensitiveKeys(bodyBytes []byte, sensitiveKeys []string) (body string, err error) {
	if len(bodyBytes) == 0 {
		return "", nil
	}

	var bodyMap map[string]any
	err = json.Unmarshal(bodyBytes, &bodyMap)
	if err != nil {
		return string(bodyBytes), err
	}

	for k := range bodyMap {
		lowerK := strings.ToLower(k)
		if IsContainsKeyword(lowerK, sensitiveKeys) {
			bodyMap[k] = "[REDACTED]"
		}
	}

	convertBodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		return string(bodyBytes), err
	}

	return string(convertBodyBytes), nil
}

func IsContainsKeyword(targetStr string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(targetStr, keyword) {
			return true
		}
	}
	return false
}

func GenerateSignature(body, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	signature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

func ParseStringToInt64(s *string) int64 {
	if s == nil || *s == "" {
		return 0
	}

	val, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}
