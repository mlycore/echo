FROM scratch
MAINTAINER mlycore <maxwell92@126.com>
ADD echo /echo
ENTRYPOINT ["/echo"]

