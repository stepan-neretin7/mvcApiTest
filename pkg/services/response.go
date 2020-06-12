package services

import "encoding/json"

func GenerateResponse(data map[string]interface{}) ([]byte, error){
	resp, err := json.Marshal(data)
	if err != nil{
		return nil, err
	}

	return resp, nil
}