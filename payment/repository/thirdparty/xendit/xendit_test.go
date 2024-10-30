package xendit

import (
	"context"
	"test_cases/payment/repository/thirdparty/xendit/mock"
	"test_cases/payment/vo"
	"testing"

	"github.com/golang/mock/gomock"
)

// NEGATIVE TEST CASES
// sending request with empty auth header - DONE
// sending request with invalid auth header
// sending request with valid auth header
// sending request with empty request body - IN PROGRESS
// "{\"error_code\":\"API_VALIDATION_ERROR\",\"message\":\"amount must be greater than 0\"}\n"
// "{\"error_code\":\"API_VALIDATION_ERROR\",\"message\":\"Only one of 'payment_method' or 'payment_method_id' should be present per request\"}\n"
// sending request with broken request body
// sending request with incomplete request body
// sending request with invalid request body
// sending request with duplicate referenceId

// POSITIVE TEST CASES
//"{\"id\":\"pr-f39025a3-e37a-4131-b89a-be9b180acf9a\",\"country\":\"ID\",\"amount\":100000,\"currency\":\"IDR\",\"business_id\":\"599bd7f1ccab55b020bb1147\",\"reference_id\":\"75d98c92-de9c-4a91-b293-dcf4cdcaf766\",\"payment_method\":{\"id\":\"pm-ec666518-cd69-44bf-a8c3-1a8f4fddde7c\",\"type\":\"VIRTUAL_ACCOUNT\",\"reference_id\":\"pm-level-5e690998-6bb7-418f-a71d-c3187360ddb8\",\"description\":null,\"created\":\"2024-10-08T16:09:25.3504544Z\",\"updated\":\"2024-10-08T16:09:25.566775031Z\",\"card\":null,\"ewallet\":null,\"direct_debit\":null,\"direct_bank_transfer\":null,\"over_the_counter\":null,\"virtual_account\":{\"amount\":100000,\"currency\":\"IDR\",\"channel_code\":\"BRI\",\"channel_properties\":{\"customer_name\":\"John Doe\",\"virtual_account_number\":\"262158018824509\",\"expires_at\":\"2055-10-08T16:09:25.402543Z\"}},\"qr_code\":null,\"metadata\":null,\"billing_information\":{\"city\":null,\"country\":\"\",\"postal_code\":null,\"province_state\":null,\"street_line1\":null,\"street_line2\":null},\"reusability\":\"ONE_TIME_USE\",\"status\":\"PENDING\"},\"description\":null,\"metadata\":{\"sku\":\"ABCDEFGH\"},\"customer_id\":null,\"capture_method\":\"AUTOMATIC\",\"initiator\":null,\"card_verification_results\":null,\"created\":\"2024-10-08T16:09:25.324531973Z\",\"updated\":\"2024-10-08T16:09:25.324531973Z\",\"status\":\"PENDING\",\"actions\":[],\"failure_code\":null,\"channel_properties\":null,\"shipping_information\":null,\"items\":null}"
// sending request with complete and valid request body
// sending request with inactive channel_code: BCA

// EDGE TEST CASES
// TODO: find the edge cases

// func TestXenditPayment_SendPaymentRequest_APIExploration(t *testing.T) {
// 	httpClient := &http.Client{}
// 	hostName := "https://api.xendit.co"
// 	// apiKey := "something"
// 	authKey := "eG5kX2RldmVsb3BtZW50X09vbUFmT1V0aCtHb3dzWTZMZUpPSHpMQ1p0U2o4NEo5a1hEbitSeGovbUhXK2J5aERRVnhoZz09Og=="
// 	xenditClient := NewXenditClient(httpClient, hostName, authKey)

// 	ctx := context.Background()
// 	paymentId, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequestPayload{})
// 	if err != nil {
// 		t.Fatalf("it should not return any error, but got: %s ", err.Error())
// 	}

// 	if paymentId == "" {
// 		t.Errorf("it should not return empty paymentID, but got: %s", paymentId)
// 	}
// }

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_EmptyAuthHeadrr(t *testing.T) {
	httpClientMock := mock.NewMockHttpConnector(gomock.NewController(t))
	hostName := "http://mock.server"
	authKey := ""
	xenditClient := NewXenditClient(httpClientMock, hostName, authKey)

	ctx := context.Background()
	_, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequestPayload{})
	if err == nil {
		t.Fatal("it should return error due to empty auth key")
	}

}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_WithEmtpyPayload(t *testing.T) {
	httpClientMock := mock.NewMockHttpConnector(gomock.NewController(t))
	hostName := "http://mock.server"
	authKey := "supersecret"
	xenditClient := NewXenditClient(httpClientMock, hostName, authKey)

	ctx := context.Background()
	paymentReq := vo.XenditPaymentRequestPayload{}
	_, err := xenditClient.SendPaymentRequest(ctx, paymentReq)
	if err == nil {
		t.Fatal("it should return error due to empty auth key")
	}

}

func TestXenditPayment_SendPaymentRequest_ErrorWhileSendingHttpRequest(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_Got200_ButEmptyResponseBody(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_Got200_ButGotBrokenResponseBody(t *testing.T) {
}

// func TestXenditPayment_SendPaymentRequest(t *testing.T) {
// 	httpClientMock := mock.NewMockHttpConnector(gomock.NewController(t))
// 	host := "http://mock.server"

// 	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))

// 	authKey := "xnd_development_OomAfOUth+GowsY6LeJOHzLCZtSj84J9kXDn+Rxj/mHW+byhDQVxhg=="
// 	xenditClient := NewXenditClient(httpClientMock, host, authKey)

// 	ctx := context.Background()
// 	paymentId, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequestPayload{})
// 	assert.Error(t, err, "it should not return error")

// 	assert.Empty(t, paymentId, "it should return a valid created payment ID")

// }
