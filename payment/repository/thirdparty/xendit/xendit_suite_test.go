package xendit_test

import (
	"context"
	"errors"
	"net/http"
	"test_cases/payment/repository/thirdparty/xendit"
	"test_cases/payment/repository/thirdparty/xendit/mock"
	"test_cases/payment/vo"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type XenditPaymentTestSuite struct {
	suite.Suite
	ctx            context.Context
	httpClientMock *mock.MockHttpConnector
	xenditClient   xendit.XenditPayment
	xenditHost     string
	xenditAuthKey  string
}

// SetupSuite() - before test
// SetupTest() - before test cases
// TearDownTest() - after each test cases
// TearDownSuiteT() - after test

func TestXenditPayment(t *testing.T) {
	suite.Run(t, new(XenditPaymentTestSuite))
}

func (s *XenditPaymentTestSuite) SetupSuite() {
	// println("1, this line is executed on setup suite")
	s.ctx = context.Background()
	ctrl := gomock.NewController(s.T())
	httpClientMock := mock.NewMockHttpConnector(ctrl)

	s.xenditHost = "http://mock.server"
	s.xenditAuthKey = "supersecret"

	xenditClient := xendit.NewXenditClient(httpClientMock, s.xenditHost, s.xenditAuthKey)

	s.httpClientMock = httpClientMock
	s.xenditClient = xenditClient
}

func (s *XenditPaymentTestSuite) SetupTest() {
	// Reset a property value of XenditPaymentRequestPayload struct
}

func (s *XenditPaymentTestSuite) TestXenditPayment_SendPaymentRequest_EmptyAuthHeader() {
	_, err := s.xenditClient.SendPaymentRequest(s.ctx, vo.XenditPaymentRequestPayload{})
	s.Error(err, "it should return error due to empty auth key")

}
func (s *XenditPaymentTestSuite) TestXenditPayment_SendPaymentRequest_WithEmptyPayload() {
	_, err := s.xenditClient.SendPaymentRequest(s.ctx, vo.XenditPaymentRequestPayload{})
	s.Error(err, "it should return error due to empty payment request payload")
}

func (s *XenditPaymentTestSuite) TestXenditPayment_SendPaymentRequest_ErrorWhileSendingHttpRequest() {
	// s.httpClientMock.EXPECT().Do(gomock.Any()).Return(&http.Response{status:200}, nil) -- Positive Test cases
	s.httpClientMock.EXPECT().Do(gomock.Any()).Return(&http.Response{}, errors.New("http failure"))

	paymentRequest := vo.XenditPaymentRequestPayload{
		Currency: "IDR",
		PaymentMethod: vo.PaymentMethod{
			PaymentMethodType: "VIRTUAL_ACCOUNT",
			ReferenceId:       uuid.NewString(),
		},
	}
	_, err := s.xenditClient.SendPaymentRequest(s.ctx, paymentRequest)
	s.Error(err, "it should return an error due to http failure")
	s.ErrorContains(err, "failure")
}
