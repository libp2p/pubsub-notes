
type PumSub struct {
  // implements ps.PubSub

  peers peerQueueMap
  subs map[ps.Topic]Subscription
}

// small type to go over the channels.
type msgErr struct {
  M   ps.Message
  Err error
}

type Subscription struct {
  // implements ps.Subscription

  topic Topic

  recv <-chan msgErr // incoming
  send chan<- msgErr // outgoing
}


