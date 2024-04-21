# message-sniffer

Starting Rabbitmq container with management UI:
```
sudo docker run -d \
    --name rabbitmq \
    -p 5672:5672 \
    -p 15672:15672 \
    rabbitmq:3-management
```

http://localhost:15672/

Queue:
```
go run main.go queue \
    --name testing \
    -u guest \
    -p guest \
    --hostname localhost \
    --port 5672
```

Topic:
```
go run main.go topic \
    -n test.exchange \
    -r org.building1.floor2.#
    -u guest \
    -p guest \
    --hostname localhost \
    --port 5672
```