require 'sinatra'
require 'json'
require 'logger'

# The default value for :bind in the development environment is 'localhost'
# so that it won't work within docker network
set :bind, ENV.fetch('BIND', '0.0.0.0')

# Put logs both file and stdout/stderr
# http://recipes.sinatrarb.com/p/middleware/rack_commonlogger
::Logger.class_eval { alias_method :write, :<< }
file = File.new('/var/log/sinatra.log', 'a+')
logger = ::Logger.new(file)

configure do
  use Rack::CommonLogger, logger
end

before do
  env['rack.errors'] = logger
  env['rack.logger'] = logger
end

get '/' do
  { log: 'ok', stream: 'stdout' }.to_json
end

get '/error' do
  { log: 'error', stream: 'stderr' }.to_json
end
