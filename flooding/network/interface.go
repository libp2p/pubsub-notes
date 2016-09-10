package network

import (
  key "github.com/ipfs/go-ipfs/blocks/key"
  bsmsg "github.com/ipfs/go-ipfs/exchange/bitswap/message"
  protocol "gx/ipfs/QmVL44QeoQDTYK8RVdpkyja7uYcK3WDNoBNHVLonf9YDtm/go-libp2p/p2p/protocol"
  context "gx/ipfs/QmZy2y8t9zQH2a1b8q2ZSLKp17ATuJoCNxxyMFG5qFExpt/go-net/context"
  peer "gx/ipfs/QmbyvM8zRFDkbFdYyt1MnevUMJ62SiSGbfDFZ3Z8nkrzr4/go-libp2p-peer"
)

var ProtocolPubSub protocol.ID = "/libp2p/pubsub/1.0.0"

// PubSubNetwork provides network connectivity for PubSub sessions
type PubSubNetwork interface {

  // SendMessage sends a BitSwap message to a peer.
  SendMessage(
    ctx context.Context,
    receiver peer.ID,
    outgoing bsmsg.BitSwapMessage) error

  // SetDelegate registers the Reciver to handle messages received
  // from the network.
  SetDelegate(Receiver)

  // ConnectTo makes sure the network is connected to given peer.
  ConnectTo(context.Context, peer.ID) error

  Routing
}

// Implement Receiver to receive messages from the BitSwapNetwork
type Receiver interface {
  ReceiveMessage(
    ctx context.Context,
    sender peer.ID,
    incoming ps.Message)

  ReceiveError(error)

  // Connected/Disconnected warns bitswap about peer connections
  PeerConnected(peer.ID)
  PeerDisconnected(peer.ID)
}
