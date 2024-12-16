FROM docker.io/node:lts-alpine as frontend

WORKDIR /webapp

COPY webapp .

RUN --mount=type=cache,target=/webapp/.npm \
    npm set cache /webapp/.npm && \
    npm ci && \
    npm run build

FROM docker.io/golang:alpine as backend

WORKDIR /app

COPY . .
COPY --from=frontend /webapp/dist ./webapp/dist

RUN --mount=type=cache,target=/go/pkg/mod go generate && go build

FROM scratch

COPY --from=backend /app/go-chef /bin/go-chef
# Atlas requires a tmp directory
COPY --from=backend --chmod=1777 /tmp /tmp
# Needed to apply migrations
COPY --from=docker.io/arigaio/atlas /atlas /bin/atlas

ENV PATH="$PATH:/bin"

ENTRYPOINT ["go-chef"]