# ---- MESSAGE ----

# ---- BUILD ----
FROM init_server_image AS build
RUN  go build -o ./build/message_server ./message


# ---- RELEASE ---- (replace golang with alpine or anything lighter please)
FROM alpine as release 

COPY --from=build /server/build .

RUN chmod 700 ./message_server
EXPOSE 50053
CMD [ "./message_server" ]