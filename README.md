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

```
go run main.go queue \
    --name testing \
    -u guest \
    -p guest \
    --hostname localhost \
    --port 5672
```