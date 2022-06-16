FROM node

WORKDIR /app/frontend

COPY . .

RUN npm install

CMD ["npm", "run", "dev", "--", "--host"]
EXPOSE 3000
