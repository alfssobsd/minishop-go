package entities

import uuid "github.com/satori/go.uuid"

type HttpCustomerRegistrationRequestEntity struct {
	CustomerUsername string `json:"username"`
	CustomerFullName string `json:"full_name"`
}

type HttpCustomerRegistrationResponseEntity struct {
	CustomerId       uuid.UUID `json:"customer_id"`
	CustomerUsername string    `json:"username"`
	CustomerFullName string    `json:"full_name"`
}
