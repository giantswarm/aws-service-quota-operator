FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./aws-service-quota-operator /aws-service-quota-operator

ENTRYPOINT ["/aws-service-quota-operator"]
