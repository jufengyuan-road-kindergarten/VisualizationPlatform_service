FROM ubuntu
MAINTAINER mzz m@mzz.pub
WORKDIR /VisualizationPlatform_service
ADD VisualizationPlatform_service ./
RUN mkdir config
ADD config/app.json ./config/
EXPOSE 8127
ENTRYPOINT ["./VisualizationPlatform_service"]