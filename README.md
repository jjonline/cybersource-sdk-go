# cybersource-sdk-go

Go client for Cybersource restful api

# usage

go mod get

````
go get github.com/jjonline/cybersource-sdk-go
````

construct client

````
client := cybersource.NewClient(cybersource.Params{
    MerchantKeyId:     "your_merchant_key_id",
    MerchantSecretKey: "your_merchant_secret_key",
    MerchantId:        "your_merchant_id",
    RequestHost:       "apitest.cybersource.com", // for test
    // RequestHost:       "api.cybersource.com", // for prod
})
````
