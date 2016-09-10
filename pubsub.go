
// Topic is a string identifier for subscriptions
type Topic string

// Subscription represents a live subscription to a pubsub
// system. You get messages out of this.
type Subscription interface {
  // PubSub returns the PubSub system.
  PubSub() PubSub

  // Topic returns the topic this subscription is for.
  Topic() string

  // Receive receives a message to m. (like io.Read)
  // It should be reentrant safe.
  Receive(m Message) (error)

  // Publish sends out given message. (like io.Write)
  // The message will be adjusted (Publisher and Topic set).
  // It should be reentrant safe.
  Publish(m Message) (error)

  // PublishPayload is a convenience methos that creates
  // the message from given payload, and then calls Publish on it.
  PublishPayload(p ipld.Payload) (error)
}

// MessageForSub is a convenience function that creates a
// message for given subscription. It sets the publisher,
// the
func MessageForSub(s Subscription, payload ipld.Object) (m Message, error) {
  //
}

// SubOpts represents a set of common subscription options
type SubOpts struct {
  Topic Topic
}

// PubSub represents a pubsub system.
type PubSub interface {

  // Peer is the current peer.ID.
  Peer() peer.ID

  // Peers is the set of peers directly connected to this node
  // in the pubsub system.
  Peers() peer.Set

  // Subscribe subscribes this node to a topic.
  // Use the subscription to receive messages on.
  // Behaves similar to a stream.
  Subscribe(SubOpts) (Subscription, error)

  // Publish sends out given message on given topic.
  // It must already be prepared
  Publish(m Message, t Topic) (error)
}


type MessageHandler func(s Subscription, m Message) (error)

// HandleSubscription is a function that sequentially calls handler
// on all messages being received from given subscription.
func HandleSubscription(s Subscription, handler MessageHandler) error {
  for {
    var m Message
    err := s.Receive(m)
    if err != nil {
      return err
    }

    // call handler.
    // synchronous. user can make async if she wishes.
    handler(m)
  }
}
