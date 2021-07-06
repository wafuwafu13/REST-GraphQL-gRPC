#!/usr/bin/env python3

from concurrent.futures import ThreadPoolExecutor
import json

import grpc
import user_pb2
import user_pb2_grpc

with open("../data/users.json") as fp:
    users = json.load(fp)


class UserManager(user_pb2_grpc.UserManagerServicer):
    def getUser(self, request, context):
        user_id = request.id

        if str(user_id) not in users:
            return user_pb2.UserResponse(error=True,
                                         message="not found")
        user = users[str(user_id)]

        result = user_pb2.User()
        result.id = user["id"]
        result.name = user["name"]

        return user_pb2.UserResponse(error=False,
                                     user=result)


def main():
    server = grpc.server(ThreadPoolExecutor(max_workers=2))

    user_pb2_grpc.add_UserManagerServicer_to_server(UserManager(), server)

    server.add_insecure_port('[::]:1234')

    server.start()

    server.wait_for_termination()


if __name__ == '__main__':
    main()
