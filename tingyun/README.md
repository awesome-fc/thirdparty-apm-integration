# AMP Solutions - FC Integrated with Tingyun
## Background
Under the general trend of cloud native, application containerization and serverlessization are being quickly accepted and selected by developers. Tingyun, as the only domestic manufacturer to enter the Gartner APM Magic Quadrant, is also always paying attention to the evolution of cloud native technology and new challenges in the field of observability.
From traditional servers, virtual machines, containers, to serverless functions, developers’ operation and maintenance burdens are gradually reduced or even eliminated. However, the demand for observable and observable insights continues to increase:
* Link tracking: fine-grained single responsibility interconnection of multiple functions, remote function calls to access databases and other cloud services require distributed link tracking
* Instance-level monitoring: Serverless black-boxes the concept of instances, but developers need to see the impact of events such as cold starts caused by instance changes on business performance.
* Richer indicators: some business key indicators such as CPU, memory, network and other cloud services have not yet been provided, and the business side collects time and effort by themselves
* Metrics, logs, link tracking and correlation to solve difficult distributed application problems.
Alibaba Cloud Function Computing is the earliest serverless cloud service with the most developers in China. Tingyun and Alibaba Cloud Function Computing have cooperated to release Tingyun APM solutions that support function computing. Function calculations are free of operation and maintenance, pay-as-you-go and other outstanding features combined with the rich and easy-to-use multi-language APM SDK to help you easily gain insights into your applications.

## Solutions to integrate with Tingyun
STEP1:  Clone repo
```
FROM golang:alpine

RUN mkdir /go/src/fc-demo
ADD ./ /go/src/fc-demo
RUN apk add git && go get github.com/TingYunAPM/go && go get github.com/TingYunAPM/go
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
while true; do curl https://{accound-id}.{region}.fc.aliyuncs.com/2016-08-15/proxy/test-apm/tingyun-golang/invoke; done
```

### View dashboard 

![](https://img.alicdn.com/imgextra/i2/O1CN011fuha11izibKhhJvn_!!6000000004484-2-tps-1448-1212.png)
![](https://img.alicdn.com/imgextra/i4/O1CN01mWPElO1TDQbmeVtMr_!!6000000002348-2-tps-1414-1210.png)
![](https://img.alicdn.com/imgextra/i2/O1CN01vYgzuR1PnpcnJoy7a_!!6000000001886-2-tps-1430-1267.png)