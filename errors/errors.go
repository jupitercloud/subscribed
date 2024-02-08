package errors

import (
	"github.com/gorilla/rpc/v2/json2"
)

func Unauthenticated() *json2.Error {
    return &json2.Error{
        Code: -1001,
        Message: "Authorization requred",
    }
}

func InvalidVendorIdClaim() *json2.Error {
    return &json2.Error{
        Code: -1002,
        Message: "Invalid vendorId claim",
    }
}

func JwtError(cause error) *json2.Error {
    return &json2.Error{
        Code: -1003,
        Message: "Invalid JWT",
        Data: map[string]interface{}{
            "details": cause.Error(),
        },
    }
}
