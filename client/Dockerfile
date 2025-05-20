FROM node:18-alpine

WORKDIR /app

COPY package*.json ./

# Install dependencies
RUN npm install

COPY . .

# Build the application
RUN npm run build

EXPOSE 3000

# Start the production server
CMD ["npm", "start"]