module V1
  module Entities
    class BookEntity < Grape::Entity
      expose :id
      expose :title
      expose :price
      expose :tax_included_price do |instance, options| # ブロックを渡す事ができる。
        instance.price * 1.08 # 実際は四捨五入したり、消費税率を定数として持っておくなどすべき
      end
      expose :author, using: V1::Entities::AuthorEntity
    end
  end
end
