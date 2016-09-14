HTTP-API for pubsub
===================

# endpoints

# `/floodsub/subscribe?topics=<[array of topics]>&feed=<bool>`

tells the node to subscribe to a topic or an array of topics

the `feed(bool)` tells the node to keep the request socket open and send all the messages in that topic

# `/floodsub/unsubscribe?topics=<[array of topics]`

tells the node to unsubscribe to a topic of an array of topics

# `/floodsub/feed?topics=<array of topics>&time-cache=<bool>`

opens a socket to get the messages on a certain topic

`time-cache(bool)` tells the node to also send all the messages it has on its time-cache

# `/floodsub/peers`

returns the peers we have in our peerSet and the subscriptions we know from them

# `/floodsub/topics`

returns the topics we are subscribed
