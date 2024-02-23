# Telegram Gemini Bot
> 一键部署使用Gemini API的Telegram机器人

Looking for an English version? [You can check it here](README.en.md)
## 准备工作
### GEMINI_API_KEY
使用Gemini API时，开发者需要专属的api key才能使用，按照[官方指引](https://ai.google.dev/tutorials/web_quickstart?hl=zh-cn#set-up-project)进行申请
### TGBOT_TOKEN
创建telegram机器人后，需要携带专属token才能使用此机器人发消息或进行其他行为设置。
与[BotFater](https://core.telegram.org/bots/tutorial)对话进行机器人申请，直到获得token。
可以参见[完整指引](https://core.telegram.org/bots/tutorial)。


## 开始部署
1. 点击 [![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fxsymphony%2Ftelegram-gemini-bot&env=GEMINI_API_KEY,TGBOT_TOKEN,DOMAIN&demo-title=Telegram%20Gemini%20Bot&demo-url=https%3A%2F%2Ftelegram-gemini-bot-ten.vercel.app%2F) 
2. 按照vercel指引，完成github账号关联、仓库创建。
3. 需要填充环境变量时，填入自己的`GEMINI_API_KEY`、`TGBOT_TOKEN`
4. 填充环境变量`DOMAIN`。因为本项目使用了telegram机器人webhook的形式接受消息，所以需要设置机器人接收消息后的回调地址为自己部署后的项目地址。
    使用`vercel`部署后，若未更改仓库名称，填入`https://telegram-gemini-bot-{{github user name}}-projects.vercel.app`的形式进行填充。
    若这一步填入错误地址，也可以在部署完成后，修改环境变量，重启任务生效正确地址。
    其他可用的地址：
    + https://telegram-gemini-{{随机后缀}}.vercel.app
    + https://telegram-gemini-bot-{{github user name}}-projects.vercel.app
    + https://telegram-gemini-bot-git-main-{{github user name}}-projects.vercel.app
5. 进入部署成功的项目页面，点击`开启机器人消息回调`，随后页面刷新，展示正确的webhook设置信息即代表生效。
6. 与自己创建的telegram bot对话，验证消息处理是否正确。

