FROM ruby:2.7

# throw errors if Gemfile has been modified since Gemfile.lock
RUN bundle config --global frozen 1

WORKDIR /usr/src/api

COPY Gemfile Gemfile.lock ./
RUN bundle install -j4 --deployment

COPY . .

EXPOSE 4567
CMD ["bundle", "exec", "ruby", "api.rb"]
