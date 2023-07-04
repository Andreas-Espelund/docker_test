import requests

print("Hello docker!")


for i in range(5):
    response = requests.get("http://localhost:8080")

    if response.status_code == 200:
        print(response.text)
    else:
        print("Request failed with code:", response.status_code)

        



