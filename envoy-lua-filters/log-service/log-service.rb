require 'sinatra'
require 'json'

set :bind, '0.0.0.0'

get '/' do
  status 200
  [200]
end

post '/log' do
  request.body.rewind
  body = request.body.read
  logger.info JSON.pretty_generate(body)
  [200]
end
