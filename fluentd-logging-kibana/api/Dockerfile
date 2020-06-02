FROM ruby:2.7

# throw errors if Gemfile has been modified since Gemfile.lock
RUN bundle config --global frozen 1

WORKDIR /usr/src/app

COPY Gemfile Gemfile.lock ./
RUN bundle install -j4 --deployment --without 'doc test'

COPY . .

EXPOSE 8080
CMD ["bundle", "exec", "bin/rails", "server", "--using", "puma", "--port", "8080", "--binding", "0.0.0.0"]
