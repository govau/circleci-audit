package main

import (
	circleci "github.com/jszwedko/go-circleci"
)

//Circle client instance
type Circle struct {
	Client circleci.Client
}

//NewCircle Create new Circle instance. The circleci token is tested and any error is returned.
func NewCircle(circleToken string) (*Circle, error) {
	circle := &Circle{}

	circle.Client = circleci.Client{Token: circleToken}

	return circle, nil
}
