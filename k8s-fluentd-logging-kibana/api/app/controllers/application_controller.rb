class ApplicationController < ActionController::API
  def route_not_found
    render 'error_pages/404', status: :not_found
  end
end
