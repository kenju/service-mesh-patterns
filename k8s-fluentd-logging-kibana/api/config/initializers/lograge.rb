Rails.application.configure do
  config.lograge.enabled = true

  # this is required for Rails API mode
  # https://github.com/roidrage/lograge#installation
  config.lograge.base_controller_class = 'ActionController::API'

  config.lograge.formatter = Lograge::Formatters::Json.new
  # config.lograge.formatter = Lograge::Formatters::Logstash.new
end
