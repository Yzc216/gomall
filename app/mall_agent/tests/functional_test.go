package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type QueryRequest struct {
	UserID  string            `json:"user_id"`
	Query   string            `json:"query"`
	Context map[string]string `json:"context,omitempty"`
}

type QueryResponse struct {
	Message      string                 `json:"message"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Success      bool                   `json:"success"`
	ErrorMessage string                 `json:"error_message,omitempty"`
}

const baseURL = "http://localhost:8080/mall"

func TestBasicQuery(t *testing.T) {
	req := QueryRequest{
		UserID: "test_user",
		Query:  "有哪些商品推荐?",
	}
	
	resp, err := sendQuery(req)
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.Message)
}

func TestEmptyUserID(t *testing.T) {
	req := QueryRequest{
		UserID: "",
		Query:  "有哪些商品推荐?",
	}
	
	resp, err := sendQuery(req)
	assert.NoError(t, err)
	assert.False(t, resp.Success)
	assert.Contains(t, resp.ErrorMessage, "用户ID不能为空")
}

func TestEmptyQuery(t *testing.T) {
	req := QueryRequest{
		UserID: "test_user",
		Query:  "",
	}
	
	resp, err := sendQuery(req)
	assert.NoError(t, err)
	assert.False(t, resp.Success)
	assert.Contains(t, resp.ErrorMessage, "查询内容不能为空")
}

func TestHistoryAndClear(t *testing.T) {
	userID := fmt.Sprintf("test_user_%d", time.Now().Unix())
	
	// 发送查询
	req := QueryRequest{
		UserID: userID,
		Query:  "有哪些商品推荐?",
	}
	_, err := sendQuery(req)
	assert.NoError(t, err)
	
	// 获取历史记录
	history, err := getHistory(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, history["history"])
	
	// 清除会话
	success, err := clearSession(userID)
	assert.NoError(t, err)
	assert.True(t, success)
	
	// 再次获取历史记录
	history, err = getHistory(userID)
	assert.NoError(t, err)
	assert.Empty(t, history["history"])
}

func TestStreamResponse(t *testing.T) {
	userID := "test_user"
	query := "我想买一件衬衫"
	
	resp, err := http.Get(fmt.Sprintf("%s/stream?user_id=%s&query=%s", baseURL, userID, query))
	assert.NoError(t, err)
	defer resp.Body.Close()
	
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "text/event-stream", resp.Header.Get("Content-Type"))
	
	// 读取部分响应内容
	buf := make([]byte, 1024)
	n, err := resp.Body.Read(buf)
	assert.True(t, n > 0 || err == io.EOF)
}

func sendQuery(req QueryRequest) (QueryResponse, error) {
	var resp QueryResponse
	
	jsonData, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}
	
	httpResp, err := http.Post(baseURL+"/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()
	
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	return resp, err
}

func getHistory(userID string) (map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/history?user_id=%s", baseURL, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func clearSession(userID string) (bool, error) {
	resp, err := http.Post(fmt.Sprintf("%s/clear?user_id=%s", baseURL, userID), "application/json", nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}
	
	return result["success"].(bool), nil
}