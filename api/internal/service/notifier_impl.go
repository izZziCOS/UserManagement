// Package service implements an AMQP-based notifier for user events
// It publishes user lifecycle events (create/update/delete) to a RabbitMQ queue
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/izzzicos/UserManagement/api/internal/contracts"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/rabbitmq/amqp091-go"
)

// AMQPNNotifier implements contracts.Notifier using AMQP/RabbitMQ
type AMQPNNotifier struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   string
}

// Ensure AMQPNNotifier implements contracts.Notifier at compile time
var _ contracts.Notifier = (*AMQPNNotifier)(nil)

// NewAMQPNotifier creates a new AMQP notifier connected to the specified queue
func NewAMQPNotifier(amqpURL, queue string) (*AMQPNNotifier, error) {
	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare the queue to ensure it exists
	_, err = ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	return &AMQPNNotifier{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}, nil
}

// NotifyUserCreated publishes a user creation event to the queue
func (n *AMQPNNotifier) NotifyUserCreated(ctx context.Context, user *models.User) error {
	return n.publish(ctx, "user.created", user)
}

// NotifyUserUpdated publishes a user update event to the queue
func (n *AMQPNNotifier) NotifyUserUpdated(ctx context.Context, oldUser, newUser *models.User) error {
	payload := struct {
		Old models.User `json:"old"`
		New models.User `json:"new"`
	}{*oldUser, *newUser}
	return n.publish(ctx, "user.updated", payload)
}

// NotifyUserDeleted publishes a user deletion event to the queue
func (n *AMQPNNotifier) NotifyUserDeleted(ctx context.Context, userID string) error {
	return n.publish(ctx, "user.deleted", map[string]string{"id": userID})
}

// publish is the internal method that handles the actual message publishing
func (n *AMQPNNotifier) publish(ctx context.Context, eventType string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		return err
	}

	log.Printf("Publishing message to queue '%s': %s", n.queue, string(body))

	err = n.channel.PublishWithContext(ctx,
		"",      // exchange
		n.queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Type:        eventType,
			Body:        body,
		})

	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	}
	return err
}

// StartTestConsumer starts a temporary consumer for testing purposes
// Note: This is for development/testing only, not for production use
func (n *AMQPNNotifier) StartTestConsumer() {
	msgs, err := n.channel.Consume(
		n.queue, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		log.Printf("Failed to register consumer: %v", err)
		return
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", string(msg.Body))
		}
	}()
}

// Close cleanly shuts down the AMQP connection and channel
func (n *AMQPNNotifier) Close() error {
	var err error

	if n.channel != nil {
		if cerr := n.channel.Close(); cerr != nil {
			err = fmt.Errorf("channel close error: %w", cerr)
		}
	}

	if n.conn != nil {
		if cerr := n.conn.Close(); cerr != nil {
			if err != nil {
				err = fmt.Errorf("%v, connection close error: %w", err, cerr)
			} else {
				err = fmt.Errorf("connection close error: %w", cerr)
			}
		}
	}

	return err
}
