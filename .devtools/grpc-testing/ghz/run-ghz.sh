ghz --insecure \
    --call user.UserService.SearchUserByUsername \
    -d '{"username": "user"}' \
    localhost:50051