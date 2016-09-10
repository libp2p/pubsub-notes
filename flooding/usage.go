
n := libp2p.Node(...)
ps := newPubSub(n)
sub, err := ps.Subscribe(SubOpts{
  Topic: "foo",
  Buffer: make(chan Message, 10)
})
