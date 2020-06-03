Rails.application.configure do
  config.lograge.enabled = true

  # this is required for Rails API mode
  # https://github.com/roidrage/lograge#installation
  config.lograge.base_controller_class = 'ActionController::API'

  config.lograge.formatter = Lograge::Formatters::Json.new

  # log errors / execptions in jsonl
  # https://github.com/roidrage/lograge#logging-errors--exceptions
  config.lograge.custom_options = lambda do |event|
    {
      exception: event.payload[:exception], # ["ExceptionClass", "the message"]
      exception_object: event.payload[:exception_object] # the exception instance
    }
  end
end
