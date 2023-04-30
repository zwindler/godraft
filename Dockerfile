FROM scratch

COPY godraft /
COPY static/ /static/
COPY templates /templates

CMD [ "/godraft" ]
