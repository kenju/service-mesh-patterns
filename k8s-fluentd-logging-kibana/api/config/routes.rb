Rails.application.routes.draw do
  get 'health', action: :show, controller: 'health'
  get '*unmatched_route', to: 'application#route_not_found'
end
