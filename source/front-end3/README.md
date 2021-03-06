Run with:

```sh
make run
```

# Known issues

Currently, for some reason when sending a message to the websocket, i.e. pulsar's websocket handler, pulsar
disconnects. The log says

```
08:25:01.826 [pulsar-web-68-14] INFO  org.apache.pulsar.websocket.AbstractWebSocketHandler - [/172.19.0.1:34354] WebSocket error on topic persistent://public/default/zombie : null
08:25:01.826 [pulsar-web-68-14] INFO  org.apache.pulsar.websocket.AbstractWebSocketHandler - [/172.19.0.1:34354] Closed WebSocket session on topic persistent://public/default/zombie. status: 1011 - reason: NullPointerException
08:25:01.826 [pulsar-web-68-14] WARN  org.apache.pulsar.websocket.ConsumerHandler - [persistent://public/default/zombie] Failed to remove consumer handler
08:25:01.826 [pulsar-io-50-21] INFO  org.apache.pulsar.broker.service.ServerCnx - [/127.0.0.1:53584] Closing consumer: 6
08:25:01.826 [pulsar-io-50-21] INFO  org.apache.pulsar.broker.service.AbstractDispatcherSingleActiveConsumer - Removing consumer Consumer{subscription=PersistentSubscription{topic=persistent://public/default/zombie, name=mysub2}, consumerId=6, consumerName=6566c, address=/127.0.0.1:53584}
08:25:01.826 [pulsar-io-50-21] INFO  org.apache.pulsar.broker.service.ServerCnx - [/127.0.0.1:53584] Closed consumer Consumer{subscription=PersistentSubscription{topic=persistent://public/default/zombie, name=mysub2}, consumerId=6, consumerName=6566c, address=/127.0.0.1:53584}
08:25:01.826 [pulsar-client-io-95-1] INFO  org.apache.pulsar.client.impl.ConsumerImpl - [persistent://public/default/zombie] [mysub2] Closed consumer
08:25:01.827 [pulsar-client-io-95-1] INFO  org.apache.pulsar.websocket.ConsumerHandler - [persistent://public/default/zombie/mysub2] Consumer was closed while receiving msg from broker
```

Since I've decided to try Kinesis instead, I'm not pursuing this now.
