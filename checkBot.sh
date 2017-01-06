/home/gempir/gempbot-up-check/gempbot-up-check

if [ $? -ne 0 ]; then
	echo "$(date) restarting gempbot" >> /var/log/gempbot-up-check.log
	sudo systemctl restart gempbot.service
else
	echo "$(date) gempbot is fine"
fi

