ghz --insecure \
    --call user.UserService.SearchUserByUsername \
    -d '{"username": "user"}' \
    -O json \
    localhost:50051 | http POST localhost:3000/api/projects/1/ingest