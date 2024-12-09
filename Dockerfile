FROM docker.io/node:lts-alpine as frontend

WORKDIR /app

COPY app .

RUN npm ci \
    && npm run build

FROM docker.io/golang:alpine as backend

WORKDIR /app

COPY . .
COPY --from=frontend /app/dist ./app/dist

RUN go generate -v \
    && go build -v

FROM scratch

COPY --from=docker.io/arigaio/atlas /atlas /bin/atlas
COPY --from=backend /app/go-chef /bin/go-chef
COPY --from=backend --chmod=1777 /tmp /tmp

ENV PATH="$PATH:/bin"

ENTRYPOINT ["go-chef"]