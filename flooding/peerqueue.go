
// peerQueue keeps a duplex stream per peer.
type peerQueue struct {
  peer     peer.ID

  // supress filters
  allFilter  SeenFilter
  peerFilter SeenFilter // here so is gc-ed with conn

  net      dsnet.Network
  receiver dsnet.Receiver

  // buffers. should be small, but nonzero. avoid one peer
  // hanging others.
  // what to do about slow peers? drop msgs? disconnect? :(
  outgoing chan Message
  incoming chan Message
}

// this queue sends out messages to given peers.
type peerQueueMap map[peer.ID]peerQueue

type netReceiver struct {
  // implements ps.network.Receiver
  psnet     psnet.PubSubNet
  pubsub    PubSub
  allFilter SeenFilter
}


func (nr *netReceiver) ReceiveMessage(
  ctx context.Context,
  sender peer.ID,
  incoming ps.Message) {

  sub := nr.pubsub.SubscriptionFor(incoming.Topic())
  if sub == nil {
    // we're not part of this thing. get outta here...
    // and stop talking to us.
    nr.psnet.SendSubscription()
  }


  k := string(incoming.MsgID())
  if seen := nr.allFilter.Filter(k); seen {
    // supressed.
    // log that we supressed the message.
    // maybe increment some stat.
    return
  }

  // from here, handle the message.
  // network context irrelevant henceforth.
  go nr.pubsub.(incoming)
}

func (nr *netReceiver) ReceiveError(error) {
  // probably stop everything?
  nr.pubsub.
}

// Connected/Disconnected warns bitswap about peer connections
func (nr *netReceiver) PeerConnected(p peer.ID) {
  // setup peerQueue
  // add peerQueue to peerSet
}

func (nr *netReceiver) PeerDisconnected(p peer.ID) {
  // rm peerQueue from peerSet
}
