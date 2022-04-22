FROM rabbitmq:3.9.0

RUN apt-get update && \
apt-get install -y curl

RUN curl -L https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.9.0/rabbitmq_delayed_message_exchange-3.9.0.ez > rabbitmq_delayed_message_exchange-3.9.0.ez && \
mv rabbitmq_delayed_message_exchange-3.9.0.ez plugins/

RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange