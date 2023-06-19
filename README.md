# Extensible Web3 API Proxy

### ✅ This is a web3 api proxy POC, it's designed to be completely extendable, you can add other EVM methods and also other EVM and non-EVM networks easily.

## 📜 Description

This Project Implemented Based on Clean Architecture in Golang.

🔰 Rule of Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

📚 More at [Uncle Bob clean-architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

📚 More at [Martin Fowler PresentationDomainDataLayering](https://martinfowler.com/bliki/PresentationDomainDataLayering.html)

### 🗺 The Diagram: 
![clean architecture](https://github.com/apm-dev/vending-machine/blob/main/clean-arch.png)


### 🏃🏽‍♂️ How To Run This Project
⚠️ Since the project already use Go Module, I recommend to put the source code in any folder except GOPATH.

#### 🧪 Run the Testing

```bash
$ make test
```

#### 🐳 Run the Applications
Here is the steps to run it with `docker-compose`

```bash
# move to directory
$ cd workspace
# Clone it
$ git clone https://github.com/apm-dev/web3-api-proxy.git
# move to project
$ cd web3-api-proxy
# (optional) Build the docker image first
$ make docker
# Run the application
$ make run
# check if the containers are running
$ docker ps
# See the logs
docker logs --follow web3_api_proxy
# test eth_getBalance method
$ curl localhost:8000/eth/balance/0x4a6330220914727a2456a8A059F1ac5b5A1E5b6a
# see the metrics
$ curl localhost:8000/metrics
# Stop
$ make stop
```

### 🛠 Tools Used:
In this project, I use some tools listed below. But you can use any similar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need. 

- All libraries listed in [`go.mod`](https://github.com/apm-dev/web3-api-proxy/blob/develop/go.mod)
- ["github.com/vektra/mockery".](https://github.com/vektra/mockery) To Generate Mocks for testing needs.
