package models

import (
	"fmt"
	"reflect"
)

type Business struct {
	Name           string      `bson:"name,omitempty"`
	Title          string      `bson:"title,omitempty"`
	Verified       bool        `bson:"verified,omitempty"`
	Images         [3]string   `bson:"images,omitempty"`
	SmallImages    [3]string   `bson:"smallImages,omitempty"`
	Rating         float64     `bson:"rating,omitempty"`
	Description    string      `bson:"description,omitempty"`
	Location       interface{} `bson:"location,omitempty"`
	Schedule       interface{} `bson:"schedule,omitempty"`
	Active         bool        `bson:"active,omitempty"`
	Category       string      `bson:"category,omitempty"`
	OffersDelivery bool        `bson:"offersDelivery,omitempty"`
	DeliveryFee    float64     `bson:"deliveryFee,omitempty"`
	AcceptsCard    bool        `bson:"acceptsCard,omitempty"`
	AcceptsCash    bool        `bson:"acceptsCash,omitempty"`
	SchoolCode     string      `bson:"schoolCode,omitempty"`
}

func NewBusiness(business Business) (*Business, error) {
	expectedLocationTypes := []reflect.Type{
		reflect.TypeOf(""),
		reflect.TypeOf([2]float64{}),
	}

	validLocationType := false

	locationType := reflect.TypeOf(business.Location)

	for _, expectedType := range expectedLocationTypes {
		if locationType == expectedType {
			validLocationType = true
			break
		}
	}

	if !validLocationType {
		return nil, fmt.Errorf("Invalid location type. Expected one of %v, but got %v", expectedLocationTypes, locationType)
	}

	return &business, nil
}
