FROM mysql

ADD ./my.cnf /etc/mysql/conf.d/my.cnf
ADD ./hosts /etc/hosts

RUN ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime &&\
    chmod 600 /etc/mysql/conf.d/my.cnf

EXPOSE 33060
EXPOSE 3306
CMD [ "mysqld" ]
