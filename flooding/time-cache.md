
pubsub time cache dupe supression

- use a time cache for messages
- msgid := msg.publisher + msg.seqno
- map[hash(msgid)] = time_msg_received
- entry expires after a constant time (eg 30s)
- discard duplicates in that time frame


```
msgid = 128 B

sha256(msgid) = 32 B
time.Time = 8 B
map entry = 8 B
q entry = 24 B
----------------
total = 72B

peers = 200
topics = 5
msgs = 10/s

10,000 msgs/s

* 72B * 30s

60 MB
```

```go
type TimeCache struct {
  Q container.List
  M map[msgid]time.Time
}


tdiff = now - t
if tdiff > 10s {
  metrics.Log()
} else if tdiff > 20s {
  metrics.Log()
}
```
