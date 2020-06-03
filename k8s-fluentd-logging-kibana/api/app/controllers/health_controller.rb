class HealthController < ApplicationController
  def show
    render json: {
      log: 'ok'
    }
  end

  def error
    raise NotImplementedError
  end
end
