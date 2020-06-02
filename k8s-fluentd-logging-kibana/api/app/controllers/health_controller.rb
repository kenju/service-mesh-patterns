class HealthController < ApplicationController
  def show
    render json: {
      log: 'ok'
    }
  end
end
