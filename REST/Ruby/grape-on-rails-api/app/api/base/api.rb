module Base
  class Api < Grape::API
	  mount V1::Root
  end
end
