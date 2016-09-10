
// MsgID is a value that uniquely identifies a message,
// and its publisher. Nodes SHOULD assume that two messages
// with the same MessageID are the same, even if their payloads
// are not identical. (The burden is on the publishers to )
type MsgID []byte

// NewMsgID constructs a MsgID from given values. The format is:
//   <varint-msgid-length><peer.ID-multihash><seqno>
func NewMsgID(publisher peer.ID, seqno []byte) MsgID {}

// Publisher returns the peer.ID of the author of this message.
// (it is a public key or a hash of a public key)
func (mid MsgID) Publisher() peer.ID {}

// SeqNo is a publisher-unique "sequence number" for this message.
// SeqNo need not be sequential (maybe rename this?), it only
// needs to be unique. It could be generated from an incrementing
// sequence number, a UUID4, a hash, a datetime + sequence number,
// etc. It is up to the publisher to enforce that SeqNo is unique
// for ever (local dates and uuid4s can work well). Receivers will keep a
// cache of undetermined length, so publishers which reuse SeqNos will
// perceive undefined behavior. (subscribers may supress or receive the
// message depending on the size of their caches, or their recency).
//
// Note that SeqNo is more useful than enforcing unique payloads, as
// the sender may wish to send identical payloads twice as two different
// messages, and these should not be supressed automatically.
func (mid MsgID) SeqNo() []byte {}


// Message is an ipld object.
// it carries a payload (an embedded ipld object)
//
// the format of this message is:
//   ---
//   # pubsub message
//   id: <MsgID-value>
//   t: [<Topic>, ...]
//   p: <Payload>


type Message interface {
  Publisher() peer.ID

  // Valid checks the integrity of this message. Invalid messages
  // (false) MUST be discarded. Valid checks the signature.
  Valid() bool

  // MsgID returns the unique value for this message.
  // It is used to discard duplicates.
  // Nodes SHOULD discard messages with the same MsgID().
  MsgID() MsgID

  // Publisher returns the publisher of this message.
  Publisher() peer.ID

  // Topics returns the topics this message is going to.
  Topics() []Topic

  // Payload returns an ipld.Object descriptor. This is the message
  // published through pubsub.
  // it may be the same memory as Message (i.e. an embedded buffer).
  Payload() ipld.Object
}

// NewMessage constructs a message from given ipld.Object.
// object may be nil (in which case this allocates a new
// object).
func NewMessage(o ipld.Object) (Message, error) {

}

// SignedMessage is a signed Message.
// it carries a payload (an embedded ipld object)
//
// the format of this message is:
//   ---
//   # SignedMessage
//   # wrapper object is a keychain.SignedObject
//   key: <keychain.Key>
//   sig: <keychain.Signature>
//   p:
//     # the pubsub message
//     id: <MsgID-value>
//     t: [<Topic>, ...]
//     p: <Payload>


type SignedMessage interface {
  Message

  // Signature returns a cryptographic signature on this object.
  // this signature must be derived from peer.ID. It does not
  // need to be the same key, but needs to trace back to it.
  // (i.e. keychain key derivation).
  Signature() keychain.Signature
}
