package service

type Payment struct {
	xendit PaymentGatewayProvider
}

type PaymentGatewayProvider interface {
	SendPaymentRequest() (paymentID string, err error)
}

func (p *Payment) Pay() (err error) {
	// Create transaction ID
	// Insert into postgres
	// Construct payment request payload
	// Run Payment Request Validation Logic
	// Call Third Party API
	_, err = p.xendit.SendPaymentRequest()
	if err != nil {
		return err
	}

	return
}
