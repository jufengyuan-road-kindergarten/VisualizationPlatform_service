FROM ubuntu
MAINTAINER mzz m@mzz.pub
WORKDIR /payment
ADD payment-management_service ./
RUN mkdir config
RUN mkdir -p docs/swagger
ADD config/app.json ./config/
ADD docs/swagger/swagger.json ./docs/swagger
EXPOSE 8127
ENTRYPOINT ["./payment-management_service"]