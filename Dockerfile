FROM scratch

MAINTAINER Richard Drew <richardkdrew@gmail.com>

# copy the files into place
ADD hello-from /hello-from

# expose the hosting ports
ENV HOST_PORT 8080
EXPOSE 8080

# set the run command
CMD ["/hello-from"]