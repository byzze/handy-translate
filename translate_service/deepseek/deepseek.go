package deepseek

import (
	"context"
	"log"
	"log/slog"
	"sync"

	"handy-translate/config"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

const Way = "deepseek"

var (
	once sync.Once
	llm  *openai.LLM
)

type Deepseek struct {
	config.Translate
}

type TranslationPayload struct {
	Source    []string `json:"source"`
	TransType string   `json:"trans_type"`
	RequestID string   `json:"request_id"`
	Detect    bool     `json:"detect"`
}

type TranslationResponse struct {
	Target []string `json:"target"`
}

func (c *Deepseek) GetName() string {
	return Way
}

func (c *Deepseek) GetLLM() *openai.LLM {
	once.Do(func() {
		var err error
		llm, err = openai.New(
			openai.WithToken(config.Data.Translate[Way].Key),
			openai.WithModel("deepseek-chat"),
			openai.WithBaseURL("https://api.deepseek.com"),
		)
		if err != nil {
			log.Fatal(err)
		}
	})
	return llm
}

func (c *Deepseek) PostQuery(query, fromLang, toLang string) ([]string, error) {
	// Initialize the OpenAI client with Deepseek model

	// 定义模板
	promptTemplate := prompts.NewPromptTemplate(
		"You are a professional translator.\n"+
			"Please translate the following text accurately and naturally.\n"+
			"Keep the original meaning, tone, and formatting.\n"+
			"Do not explain or add anything else.\n\n"+
			"If the text is Chinese, translate to English.\n"+
			"If the text is English, translate to Chinese.\n\n"+
			"Text:\n{{.text}}",
		[]string{"text"},
	)

	// 构建输入
	promptValue, err := promptTemplate.Format(map[string]any{
		"text": query,
	})
	if err != nil {
		panic(err)
	}

	// 调用 LLM
	resp, err := llms.GenerateFromSinglePrompt(context.Background(), c.GetLLM(), promptValue)
	if err != nil {
		panic(err)
	}

	slog.Info(resp)

	return []string{resp, ""}, nil
}

// PostQueryStream 流式翻译
func (c *Deepseek) PostQueryStream(query, fromLang, toLang string, callback func(chunk string)) error {
	// 定义模板
	promptTemplate := prompts.NewPromptTemplate(
		"You are a professional translator.\n"+
			"Please translate the following text accurately and naturally.\n"+
			"Keep the original meaning, tone, and formatting.\n"+
			"Do not explain or add anything else.\n\n"+
			"If the text is Chinese, translate to English.\n"+
			"If the text is English, translate to Chinese.\n\n"+
			"Text:\n{{.text}}",
		[]string{"text"},
	)

	// 构建输入
	promptValue, err := promptTemplate.Format(map[string]any{
		"text": query,
	})
	if err != nil {
		return err
	}

	// 流式调用 LLM
	ctx := context.Background()
	_, err = c.GetLLM().GenerateContent(ctx, []llms.MessageContent{
		{
			Parts: []llms.ContentPart{
				llms.TextPart(promptValue),
			},
			Role: llms.ChatMessageTypeHuman,
		},
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		// 每次接收到数据块时调用回调函数
		if len(chunk) > 0 {
			callback(string(chunk))
		}
		return nil
	}))

	return err
}
