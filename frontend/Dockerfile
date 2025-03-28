# Use a lightweight Node.js image
FROM node:20-alpine

WORKDIR /app

# Install dependencies first for caching
COPY package*.json ./
RUN npm install

# Copy the rest of the code
COPY . .

EXPOSE 3000

# Use development server for local testing (for production consider building and serving statically)
CMD ["npm", "run", "dev"]
