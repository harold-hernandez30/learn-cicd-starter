package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey_Empty(t *testing.T) {
	fmt.Println("========================")
	fmt.Println("TestGetAPIKey_Empty")
	fmt.Println("========================")
	noResponseHeader := http.Header{}

	res, err := GetAPIKey(noResponseHeader)

	if err == ErrNoAuthHeaderIncluded && res == "" {
		fmt.Println("PASSED! Returns error: no authorization header included")
	} else {
		t.Errorf("empty header failed to return error.")
	}
}

func TestGetAPIKey_Not_ApiKey(t *testing.T) {
	fmt.Println("========================")
	fmt.Println("TestGetAPIKey_Not_ApiKey")
	fmt.Println("========================")
	noResponseHeader := http.Header{}
	noResponseHeader.Add("Authorization", "Faki 12313123123123")

	res, err := GetAPIKey(noResponseHeader)

	if err != nil && res == "" {
		fmt.Printf("PASSED! Returns error: %s\n", err)
	} else {
		t.Errorf("ApiKey not used")
	}
}

func TestGetAPIKey_ProperApiKey(t *testing.T) {
	fmt.Println("========================")
	fmt.Println("TestGetAPIKey_ProperApiKey")
	fmt.Println("========================")
	noResponseHeader := http.Header{}
	noResponseHeader.Add("Authorization", "ApiKey 12313123123123")

	res, err := GetAPIKey(noResponseHeader)

	if err == errors.New("malformed authorization header") && res == "" {
		t.Errorf("Authorization malformed")
	} else {

		if res == "12313123123123" {
			fmt.Println("PASSED! Returns proper API key")
		} else {
			t.Errorf("Authorization malformed: %s", res)
		}
	}
}
