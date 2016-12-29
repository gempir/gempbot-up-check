/home/gempir/gempbot-up-check/gempbot-up-check

if [ $? -ne 0 ]; then
	echo "$(date) restarting gempbot" >> /var/log/gempbot-up-check.log
	sudo systemctl gempbot.service restart
fi

