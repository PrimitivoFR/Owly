# ---- USER ----

# ---- BUILD ----
FROM init_server_image AS build
RUN  go build -o ./build/user_server ./user


# ---- RELEASE ---- (replace golang with alpine or anything lighter please)
FROM alpine as release 

COPY --from=build /server/build .

RUN chmod 700 ./user_server
EXPOSE 50051
CMD [ "./user_server" ]