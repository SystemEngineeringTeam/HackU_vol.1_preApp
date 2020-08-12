FROM node:lts-alpine3.12
WORKDIR /app

COPY package.json .
COPY yarn.lock .

# RUN  mkdir ~/.npm-global&&\
#     npm config set prefix '~/.npm-global'

# ENV PATH=~/.npm-global/bin:$PATH

RUN apk update && \
    yarn global add @vue/cli && \
    yarn install

# RUN yarn global add @vue/cli&&\
#     yarn install
# CMD ["yarn","serve"]
