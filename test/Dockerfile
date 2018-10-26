FROM golang:1.11.1

RUN apt update \ 
    && apt install \
       locales \
       wget \
       make \
       git \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

