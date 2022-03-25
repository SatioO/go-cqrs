package message

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// HandlerFunc's are executed parallel when multiple messages was received
// (because msg.Ack() was sent in HandlerFunc or Subscriber supports multiple consumers).
type HandlerFunc func(msg *Message) ([]*Message, error)

// NoPublishHandlerFunc is HandlerFunc alternative, which doesn't produce any messages.
type NoPublishHandlerFunc func(msg *Message) error

type Router struct {
	handlers map[string]*handler
}

func NewRouter() (*Router, error) {
	return &Router{
		handlers: map[string]*handler{},
	}, nil
}

func (r *Router) AddHandler(
	handlerName string,
	subscribeTopic string,
	subscriber Subscriber,
	handlerFunc HandlerFunc,
) *Handler {
	logrus.Info("Adding handler", map[string]any{
		"handler_name": handlerName,
		"topic":        subscribeTopic,
	})

	newHandler := &handler{
		name: handlerName,

		subscriber:     subscriber,
		subscribeTopic: subscribeTopic,

		handlerFunc: handlerFunc,
	}

	r.handlers[handlerName] = newHandler

	return &Handler{
		router:  r,
		handler: newHandler,
	}
}

func (r *Router) AddNoPublisherHandler(
	handlerName string,
	subscribeTopic string,
	subscriber Subscriber,
	handlerFunc NoPublishHandlerFunc,
) *Handler {
	handlerFuncAdapter := func(msg *Message) ([]*Message, error) {
		return nil, handlerFunc(msg)
	}
	return r.AddHandler(handlerName, subscribeTopic, subscriber, handlerFuncAdapter)
}

// Handler handles Messages.
type Handler struct {
	router  *Router
	handler *handler
}

type handler struct {
	name string

	subscriber     Subscriber
	subscribeTopic string
	subscriberName string

	handlerFunc HandlerFunc

	messagesCh <-chan *Message
}

func (r *Router) Run(ctx context.Context) (err error) {
	if err := r.RunHandlers(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Router) RunHandlers(ctx context.Context) error {
	for _, h := range r.handlers {
		h := h
		_, err := h.subscriber.Subscribe(ctx, h.subscribeTopic)
		if err != nil {
			return errors.Wrapf(err, "cannot subscribe topic %s", h.subscribeTopic)
		}

		logrus.Info("Subscribing to topic", map[string]any{
			"subscriber_name": h.name,
			"topic":           h.subscribeTopic,
		})

		go func() {
			h.run(ctx)
		}()
	}

	return nil
}

func (h *handler) run(ctx context.Context) {
	for msg := range h.messagesCh {
		go h.handleMessage(msg, h.handlerFunc)
	}
}

func (h *handler) handleMessage(msg *Message, handler HandlerFunc) {
	logrus.Trace("Received message", msg.UUID)
}
