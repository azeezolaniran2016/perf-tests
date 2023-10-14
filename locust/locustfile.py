from locust import HttpUser, task

class FibonnacciServer(HttpUser):
    @task
    def fibonacci(self):
        self.client.get("", name="fibbonaci")
