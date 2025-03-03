package response

type ApiResponse struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Errors    interface{} `json:"errors"`
	RequestId string      `json:"request_id"`
}

type ApiResponseBuilder struct {
	apiResponse ApiResponse
}

func NewApiResponseBuilder() *ApiResponseBuilder {
	return &ApiResponseBuilder{}
}

func (b *ApiResponseBuilder) SetStatus(status int) *ApiResponseBuilder {
	b.apiResponse.Status = status
	return b
}

func (b *ApiResponseBuilder) SetMessage(message string) *ApiResponseBuilder {
	b.apiResponse.Message = message
	return b
}

func (b *ApiResponseBuilder) SetData(data interface{}) *ApiResponseBuilder {
	b.apiResponse.Data = data
	return b
}

func (b *ApiResponseBuilder) SetErrors(errors interface{}) *ApiResponseBuilder {
	b.apiResponse.Errors = errors
	return b
}

func (b *ApiResponseBuilder) SetRequestId(requestId string) *ApiResponseBuilder {
	b.apiResponse.RequestId = requestId
	return b
}

func (b *ApiResponseBuilder) Build() ApiResponse {
	return b.apiResponse
}
