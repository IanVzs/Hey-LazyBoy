package models

type LLMRequest struct {
	Content string `json:"content"`
}

// OpenAI
type PayloadOpenAI struct {
	Messages         []MessageOpenAI `json:"messages"`
	MaxTokens        int64           `json:"max_tokens"`
	Temperature      float64         `json:"temperature"`
	TopP             int64           `json:"top_p"`
	FrequencyPenalty int64           `json:"frequency_penalty"`
	PresencePenalty  int64           `json:"presence_penalty"`
	Model            string          `json:"model"`
}

// ChatCompletion 是JSON顶层对象对应的Go结构体

type ResponseChatCompletionOpenAI struct {
	ID                  string                     `json:"id"`
	Object              string                     `json:"object"`
	Created             int64                      `json:"created"`
	Model               string                     `json:"model"`
	PromptFilterResults []PromptFilterResultOpenAI `json:"prompt_filter_results"`
	Choices             []ChoiceOpenAI             `json:"choices"`
	Usage               UsageOpenAI                `json:"usage"`
}

// PromptFilterResultOpenAI 是JSON中prompt_filter_results字段对应的Go结构体

type PromptFilterResultOpenAI struct {
	PromptIndex          int                        `json:"prompt_index"`
	ContentFilterResults ContentFilterResultsOpenAI `json:"content_filter_results"`
}

// ContentFilterResults 是JSON中content_filter_results字段对应的Go结构体

type ContentFilterResultsOpenAI struct {
	Hate     FilterResultOpenAI `json:"hate"`
	SelfHarm FilterResultOpenAI `json:"self_harm"`
	Sexual   FilterResultOpenAI `json:"sexual"`
	Violence FilterResultOpenAI `json:"violence"`
}

// FilterResultOpenAI 是JSON中过滤结果字段（如hate, self_harm等）对应的Go结构体

type FilterResultOpenAI struct {
	Filtered bool   `json:"filtered"`
	Severity string `json:"severity"`
}

// ChoiceOpenAI 是JSON中choices字段对应的Go结构体

type ChoiceOpenAI struct {
	FinishReason         string                     `json:"finish_reason"`
	Index                int                        `json:"index"`
	Message              MessageOpenAI              `json:"message"`
	ContentFilterResults ContentFilterResultsOpenAI `json:"content_filter_results"`
	Logprobs             interface{}                `json:"logprobs"`
}

// MessageOpenAI 是JSON中message字段对应的Go结构体

type MessageOpenAI struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// UsageOpenAI 是JSON中usage字段对应的Go结构体

type UsageOpenAI struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
