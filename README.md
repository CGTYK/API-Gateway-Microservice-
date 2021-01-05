# API Gateway Microservice Using NGINX

## Requirements

**Docker**:  https://docs.docker.com/install/

**Golang**: https://golang.org/dl/ 



## Project Structure

```
.
./auth/                 # Our service for authorization
./coffee/               # Our service for delivering coffee
./tea/                  # Our service for delivering tea
./nginx/                # Files for configuration of our NGINX instance
./docker-compose.yml
```

With our docker-compose file consisting of four services: 
- The NGINX gateway/proxy
- The Coffee & Tea services
- The Authorisation service


Let's restart our services with `docker-compose up` and try to call our service:

``` 
$ curl https://localhost/tea -H "Authorization: CSlkjdfj3423lkj234jj==" -k
> Your Coffee has been served by - 668360b94fa8

```

NGINX: 
* https://www.nginx.com/
* https://www.nginx.com/resources/library/complete-nginx-cookbook/
* https://www.nginx.com/resources/library/designing-deploying-microservices/

Golang: https://golang.org/

Docker: https://www.docker.com/

Building Microservices: https://www.amazon.com/Building-Microservices-Designing-Fine-Grained-Systems/dp/1491950358/ref=sr_1_1?ie=UTF8&qid=1527804394&sr=8-1&keywords=building+microservices
