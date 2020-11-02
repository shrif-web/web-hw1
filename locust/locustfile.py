import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def index_page(self):
        self.client.get("/nodejs/write/?num=1")
        self.client.get("/go/write/?num=1")


    @task
    def index_page(self):
        self.client.post("/nodejs/sha256", json={"num1":"2", "num2":"3"}) 
        self.client.post("/go/sha256", json={"num1":"2", "num2":"3"}) 


    @task(3)
    def view_item(self):
        for item_id in range(10):
            self.client.get(f"/nodejs/write/?num={item_id}", name="/nodejs/write")
            self.client.post("/nodejs/sha256", json={"num1":"{item_id}", "num2":"{item_id}"})
            self.client.post("/nodejs/sha256", json={"num1":"{item_id}", "num2":"{item_id}"}) 
            self.client.get("/go/write/?num={item_id}")
            time.sleep(1)


