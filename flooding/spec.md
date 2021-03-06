flood-pubsub:
- keep a peerset, like bitswap
- use a topic object
  - name: a string name (or can be some randomness, uuid4 is fine)
  - authentication modes:
    - any: anyone can publish
    - key: a list of keys to trust, only messages signed by those keys are accepted
    - wot: a web of trust, like key, but publishers can add other publishers by signing a certificate
  - encryption modes:
    - none: not encrypted
    - shared key: hash of the key is on topic object. messages are encrypted with the shared key.
  - the "topic id" is the hash of the topic object
- use topic subscriptions
  - the "topic id" is the hash of the topic object
  - can subscribe to many topics
- topic subscriptions advertised to direct peers
  - uses deltas (sub and unsub).
  - like bitswap wantlist
- uses one stream per peer
  - all topics' messages
  - control messages too (subslist)
- messages
  - have a publisher peer.ID
  - have a unique msgID (peer.ID + seqno)
  - have a list of topics
  - have a payload
  - (maybe) are authenticated (signed)
- incoming messages
  - discard messages for topics we are not subscribed to
  - discard duplicates based on msgID (SeenFilter)
  - then give to clients (through subscription object)
  - then fwd to all other peers subscribed to topic
    - this is flooding
- outgoing messages
  - never send msg to peer who sent it to us (SeenFilter per peer)
- client interface
  - subscribe to topic
    - get subscription object
    - send/recv messages on that
    - could set a handler function too
pubsub topic + peer discovery:
- floodsub uses dht for discovery 
  - providers on a TopicDescriptor object
- floodsub add RPC "get peers" returns PeerInfos (a discovery protocol)
- floodsub considers peers who publish messages
- dont connect to ALL found peers, only consider them
  - we should look at keeping a max number of peers per topic
  - if max: we should look at topology forming to avoid partitions

libp2p pubsub:
- it's a pubsub

libp2p peer discovery:
- can discover peers through messages received
- dont connect to all, but consider them

libp2p peer routing:
- can publish a message asking for a peer
- that peer or peers who know that peer can respond directly

libp2p content routing:
- can publish a getproviders message
- peers with matching records can respond directly (or publish them)

libp2p transport:
- can use pubsub as an ethernet.
- horrible perf but may be useful for some use cases
