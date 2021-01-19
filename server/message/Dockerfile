# ---- MESSAGE ----
ARG INIT_SERVER_IMAGE
# ---- BUILD ----
FROM $INIT_SERVER_IMAGE AS build
RUN  go build -o ./build/message_server ./message


# ---- RELEASE ---- (replace golang with alpine or anything lighter please)
FROM golang as release 

COPY --from=build /server/build .

RUN chmod 700 ./message_server
EXPOSE 50053
CMD [ "./message_server" ]