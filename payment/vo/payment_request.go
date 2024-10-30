package vo

type (
	XenditPaymentRequestPayload struct {
		Currency      string        `json:"currency"`
		Amount        int           `json:"amount"`
		PaymentMethod PaymentMethod `json:"payment_method"`
		Metadata      Metadata      `json:"metadata"`
	}

	PaymentMethod struct {
		PaymentMethodType string         `json:"type"`
		Reusability       string         `json:"reusability"`
		ReferenceId       string         `json:"reference_id"`
		VirtualAccount    VirtualAccount `json:"virtual_account"`
	}

	VirtualAccount struct {
		ChannelCode       string            `json:"channel_code"`
		ChannelProperties ChannelProperties `json:"channel_properties"`
	}

	ChannelProperties struct {
		CustomerName string `json:"customer_name"`
	}

	Metadata struct {
		Sku string `json:"sku"`
	}
)
