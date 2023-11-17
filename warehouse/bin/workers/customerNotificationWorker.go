package main

import "docker/warehouse/notifications"

func main() {
	worker := notifications.NewRabbitMQChannel()

	worker.ProcessNotifications()
}
