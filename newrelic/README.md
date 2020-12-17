# AMP Solutions - FC Integrated with NewRelic
## Background
Under the general trend of cloud native, application containerization and serverlessization are being quickly accepted and selected by developers. New Relic is paying attention to the evolution of cloud native technology and new challenges in the field of observability.
From traditional servers, virtual machines, containers, to serverless functions, developers’ operation and maintenance burdens are gradually reduced or even eliminated. However, the demand for observable and observable insights continues to increase:
* Link tracking: fine-grained single responsibility interconnection of multiple functions, remote function calls to access databases and other cloud services require distributed link tracking
* Instance-level monitoring: Serverless black-boxes the concept of instances, but developers need to see the impact of events such as cold starts caused by instance changes on business performance.
* Richer indicators: some business key indicators such as CPU, memory, network and other cloud services have not yet been provided, and the business side collects time and effort by themselves
* Metrics, logs, link tracking and correlation to solve difficult distributed application problems.
Alibaba Cloud Function Compute supports integrating new relic. 


## Solutions to integrate with Tingyun
STEP1:  Clone repo
```
FROM golang:alpine

RUN mkdir /go/src/fc-demo
ADD ./ /go/src/fc-demo
RUN apk add git && go get github.com/newrelic/go-agent
RUN cd /go/src/fc-demo && go build -o fc-demo ./main_fc_golang.go

WORKDIR /go/src/fc-demo

ENTRYPOINT [“./fc-demo”]
```

STEP2:  deploy to Function Compute
```
fun deploy
```

STEP3：invoke to produce some metrics
```
while true; do curl https://{accound-id}.{region}.fc.aliyuncs.com/2016-08-15/proxy/test-apm/newrelic-golang/invoke; done
```

### View dashboard 
![](https://img.alicdn.com/imgextra/i1/O1CN019NiOo31jibQ3vcqqF_!!6000000004582-2-tps-2838-1520.png)
![](https://img.alicdn.com/imgextra/i1/O1CN01Nq9y6A1OGZF54kfgn_!!6000000001678-2-tps-2854-1528.png)
![](https://img.alicdn.com/imgextra/i3/O1CN01vyQ16z259GP8xBeBR_!!6000000007483-2-tps-2858-1444.png)
![](https://img.alicdn.com/imgextra/i4/O1CN01oqgdwe22NyHDg4nYn_!!6000000007109-2-tps-2868-1478.png)

