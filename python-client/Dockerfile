# 1. Load base image
FROM python:latest

# 2. Set working dir of host/server where the code will be executed
WORKDIR /usr/app/src

# 3. Copy the code over to the dir
COPY . .

# 4. Install dependencies
RUN pip3 install --no-cache-dir -r requirements.txt

# 5. Run the application

CMD ["python3", "main.py"]