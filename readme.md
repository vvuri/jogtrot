## jogtrot

QA DevOps scheduler for posting information in hidden Telegram channel.
- Git commits 
- Jenkins tasks
- Jira issue status changes

### Steps
1. Create hiden Telegram channal - in Web UI
    https://t.me/joinchat/AAAA...
2. Go to BotFather in Web UI @botfather
    /start
    /newbot
    enter name: JogTrotBot
    enter user name: JTUserBot
3. Manual in Web UI add user JTUserBot to hidden chanel as Admin with only post message.
4. Secret param data from ENV + .env file for easy load
    M_BOT_API_KEY = ####
    TM_CHANNEL_ID = -100####
    DEBUG_MODE = false
    SCHEDULER_TIMER_SEC = 60
    PORT = 5000
5. Send test message 
6. Add requirenment
    $ go mod init jogtrot
7. $ go mod vendor


#### Future
Bot - for get information about workin process
- kubectl get pod
- ping resource
- state Jenkins
- work staging services

#### Steps
- img -> sendPhoto
- Scheduler
    - goCron  https://github.com/go-co-op/gocron
    - горутин запуск процеса
    - конфиг - да нет для запуска
    - простой способ https://stackoverflow.com/questions/53057237/how-to-schedule-a-task-at-go-periodically
    - разобрать библиотеку github.com/prprprus/scheduler
- Рефакторинг
    - разделение на модули
    - шедулер
    - каждое действие как отдельный блок
- Bitbacket API
- Jira API
- Jenkins API
- Add Logs
- Test

#### Links
- https://core.telegram.org/api
- https://core.telegram.org/bots
- Telegram bot golang
    https://github.com/go-telegram-bot-api/telegram-bot-api


