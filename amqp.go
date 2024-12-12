// Copyright (c) 2024 Daniel Pickens
// This file is part of lorna which is released under MIT license.
// See file LICENSE for full license details.

package lorna

import (
	"log"

	"github.com/streadway/amqp"
)

// deliveryAck acknowledges delivery message with retries on error
func deliveryAck(delivery amqp.Delivery) {
	var err error
	for retryCount := 3; retryCount > 0; retryCount-- {
		if err = delivery.Ack(false); err == nil {
			break
		}
	}
	if err != nil {
		log.Printf("amqp_backend: failed to acknowledge result message %+v: %+v", delivery.MessageId, err)
	}
}
