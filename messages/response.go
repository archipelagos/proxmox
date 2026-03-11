package messages

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Response struct {
	Data map[string](any) `json:"data"`
}

// TODO: Implement via Unmarshaller interface.
func ParseResponse(bytes []byte) (Response, error) {
	var response Response
	unmarshalErr := json.Unmarshal(bytes, &response)
	if unmarshalErr != nil {
		return Response{}, unmarshalErr
	}

	for key, value := range response.Data {
		if response.Data[key] == nil {
			return Response{}, fmt.Errorf("Structure different than expected, no: '%v' was found\n", key)
		}

		_ = value
		switch key {
		case "ticket":
			response.Data[key] = value
		case "username":
			response.Data[key] = value
		case "CSRFPreventionToken":
			response.Data[key] = value
		case "cap":
			capValue, ok := value.(map[string](any))
			if !ok {
				return Response{}, fmt.Errorf("Type of cap different than expected: %v\n", reflect.TypeOf(value))
			}
			/*
				var responseData ResponseData

				responseData, parseResponseAccessErr := ParseResponseData([]byte(capValue))
				if parseResponseAccessErr != nil {
					return Response{}, fmt.Errorf("Failed to parse: %v\n", parseResponseAccessErr)
				}

				response.Data[key] = responseData
			*/
			response.Data[key] = capValue
		default:
		}
	}

	return response, nil
}
