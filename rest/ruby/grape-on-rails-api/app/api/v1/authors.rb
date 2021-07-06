module V1
  class Authors < Grape::API
    resources :authors do
      desc "returns all authors"
      get "/" do
        @authors = Author.all # 最後に評価された値がレスポンスとして返される
        present @authors, with: V1::Entities::AuthorEntity # @authors を V1::Entities::AuthorEntityを使用して返却する
      end

      desc "returns an author"
      params do
        requires :id, type: Integer
      end
      get "/:id" do
        @author = Author.find(params[:id])
        present @author, with: V1::Entities::AuthorEntity
      end

      desc "Create an author"
      params do
        requires :name, type: String
      end
      post "/" do
        @author = Author.new(name: params[:name])

        # 作成の成否によってレスポンスを分ける
        if @author.save
          status 201
          present @author, with: V1::Entities::AuthorEntity
        else
          status 400
          present @author.errors.full_messages
        end
      end

      desc "Delete an author"
      params do
        requires :id, type: Integer
      end
      delete "/:id" do
        @author = Author.find(params[:id])

        # 削除の成否によってレスポンスを分ける
        if @author.destroy
          status 204
          present nil
        else
          status 400
          present @author.errors.full_messages
        end
      end
    end
  end
end
