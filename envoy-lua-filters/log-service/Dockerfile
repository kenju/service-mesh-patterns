FROM ruby:2.6.4

RUN mkdir /app

COPY Gemfile Gemfile.lock /tmp/
RUN cd /tmp && bundle install -j4 --deployment --without 'development test' --path .bundle

WORKDIR /app
COPY . /app
RUN cp -a /tmp/.bundle /app/

EXPOSE 4567

CMD ["bundle", "exec", "ruby", "log-service.rb"]
