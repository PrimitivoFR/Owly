# ---- AUTH ----

# ---- BUILD ----
FROM init_server_image AS build
RUN  go build -o ./build/auth_server ./auth


# ---- RELEASE ---- (replace golang with alpine or anything lighter please)
FROM golang as release 

COPY --from=build /server/build .

RUN chmod 700 ./auth_server
EXPOSE 50054
CMD [ "./auth_server" ]