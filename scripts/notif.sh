if [[ "$TRAVIS_BRANCH" == "master" ]]; then
  wget https://raw.githubusercontent.com/DiscordHooks/travis-ci-discord-webhook/master/send.sh
  chmod +x send.sh
  ./send.sh $1 $WEBHOOK_URL
fi
