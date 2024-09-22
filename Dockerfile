FROM ubuntu:22.04

RUN apt-get update
USER root

# Update the package lists and install the 'tzdata' package
RUN apt-get update && \
    DEBIAN_FRONTEND="noninteractive" apt-get install -y tzdata

# Set the timezone to Africa/Johannesburg
ENV TZ=Africa/Johannesburg
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
#

RUN apt-get clean
RUN apt-get autoclean
RUN apt-get update -y
RUN apt-get install -y alien
RUN apt-get install -y lftp

## Update and install necessary packages
RUN apt-get update && apt-get install -y \
    # Install your dependencies here if needed \
    && rm -rf /var/lib/apt/lists/*


WORKDIR /app
COPY programfile .


EXPOSE 8080

CMD ["./programfile"]