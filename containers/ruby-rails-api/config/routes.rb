Rails.application.routes.draw do
  get 'health', action: :show, controller: 'health'
end
