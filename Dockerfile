FROM alpine
ADD logtest-srv /logtest-srv
ENTRYPOINT [ "/logtest-srv" ]
