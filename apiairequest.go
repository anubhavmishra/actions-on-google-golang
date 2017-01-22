package actions_on_google_golang

import "time"

type ApiAiRequest struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Result    struct {
		Source           string `json:"source"`
		ResolvedQuery    string `json:"resolvedQuery"`
		Speech           string `json:"speech"`
		Action           string `json:"action"`
		ActionIncomplete bool   `json:"actionIncomplete"`
		Parameters       struct {
		} `json:"parameters"`
		Contexts []interface{} `json:"contexts"`
		Metadata struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech   string `json:"speech"`
			Messages []struct {
				Type   int    `json:"type"`
				Speech string `json:"speech"`
			} `json:"messages"`
		} `json:"fulfillment"`
		Score float64 `json:"score"`
	} `json:"result"`
	Status struct {
		Code      int    `json:"code"`
		ErrorType string `json:"errorType"`
	} `json:"status"`
	SessionID       string `json:"sessionId"`
	OriginalRequest struct {
		Source string `json:"source"`
		Data   struct {
			Inputs []struct {
				Arguments []struct {
					RawText   string `json:"raw_text"`
					TextValue string `json:"text_value"`
					Name      string `json:"name"`
				} `json:"arguments"`
				Intent    string `json:"intent"`
				RawInputs []struct {
					Query     string `json:"query"`
					InputType int    `json:"input_type"`
				} `json:"raw_inputs"`
			} `json:"inputs"`
			User struct {
				UserID string `json:"user_id"`
			} `json:"user"`
			Conversation struct {
				ConversationToken string `json:"conversation_token"`
				ConversationID    string `json:"conversation_id"`
				Type              int    `json:"type"`
			} `json:"conversation"`
		} `json:"data"`
	} `json:"originalRequest"`
}