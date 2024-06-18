package cybersource

type Client struct {
	merchantKeyId     string
	merchantSecretKey string
	merchantId        string
	requestHost       string
	isDebugPrint      bool
}

type Params struct {
	MerchantKeyId     string
	MerchantSecretKey string
	MerchantId        string
	RequestHost       string
	IsDebugPrint      bool // 是否打印出网络请求包用语debug
}

func NewClient(params Params) *Client {
	return &Client{
		merchantKeyId:     params.MerchantKeyId,
		merchantSecretKey: params.MerchantSecretKey,
		merchantId:        params.MerchantId,
		requestHost:       params.RequestHost,
		isDebugPrint:      params.IsDebugPrint,
	}
}
