## stand-alone
### start
```bash
docker run -d ttbb/rabbitmq:mate
```
### start port expose
```bash
docker run -p 4369:4369 -p 5672:5672 -p 25672:25672 ttbb/rabbitmq:mate
```