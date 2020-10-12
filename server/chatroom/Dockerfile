# ---- CHATROOM ----

# ---- BUILD ----
FROM init_server_image AS build
RUN  go build -o ./build/chatroom_server ./chatroom


# ---- RELEASE ---- (replace golang with alpine or anything lighter please)
FROM alpine as release 

COPY --from=build /server/build .

RUN chmod 700 ./chatroom_server
EXPOSE 50052
CMD [ "./chatroom_server" ]