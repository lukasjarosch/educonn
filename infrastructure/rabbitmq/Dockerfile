FROM rabbitmq:3-management
ADD ./plugins/autocluster-0.10.0.ez /usr/lib/rabbitmq/lib/rabbitmq_server-$RABBITMQ_VERSION/plugins/
RUN rabbitmq-plugins enable --offline autocluster
