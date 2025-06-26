# ใช้ Node.js LTS เป็น base image
FROM node:18-alpine

# สร้าง working directory
WORKDIR /app

# Copy package.json และ package-lock.json
COPY package*.json ./

# ติดตั้ง dependencies
RUN npm install

# Copy ไฟล์ทั้งหมดเข้าไปใน container
COPY . .

# เปิดพอร์ตที่ใช้สำหรับ development server (ปกติคือ 3000)
EXPOSE 3000

# รัน development server
CMD ["npm", "run", "dev"]
