# Telegram Gemini Bot
> One-click deployment of a Telegram bot using the Gemini API

## Preparation
### GEMINI_API_KEY
When using the Gemini API, developers need an exclusive API key to utilize its functionalities. Follow the [official instructions](https://ai.google.dev/tutorials/web_quickstart?hl=zh-cn#set-up-project) to obtain one.

### TGBOT_TOKEN
After creating a Telegram bot, you need to obtain a token specific to that bot in order to send messages or configure other bot behaviors. Initiate a conversation with [BotFather](https://core.telegram.org/bots/tutorial) to request a token. You can refer to the [complete guide](https://core.telegram.org/bots/tutorial) for more details.

## Getting Started with Deployment
1. Click on [![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fxsymphony%2Ftelegram-gemini-bot&env=GEMINI_API_KEY,TGBOT_TOKEN,DOMAIN&demo-title=Telegram%20Gemini%20Bot&demo-url=https%3A%2F%2Ftelegram-gemini-bot-ten.vercel.app%2F).
2. Follow the Vercel instructions to link your GitHub account and create a repository.
3. When prompted to fill in environment variables, provide your `GEMINI_API_KEY` and `TGBOT_TOKEN`.
4. Fill in the `DOMAIN` environment variable. Since this project uses Telegram bot webhook for message reception, you need to set the callback URL for your bot to the deployed project's URL. If deploying with Vercel and not changing the repository name, fill in `https://telegram-gemini-bot-{{github user name}}-projects.vercel.app`.
   If you input an incorrect address at this step, you can modify the environment variable after deployment and restart the task to apply the correct address.
   Other available addresses include:
   + https://telegram-gemini-{{random suffix}}.vercel.app
   + https://telegram-gemini-bot-{{github user name}}-projects.vercel.app
   + https://telegram-gemini-bot-git-main-{{github user name}}-projects.vercel.app
5. Once the deployment is successful, go to the project page and click on `Enable bot message callbacks`. After the page refreshes, if it displays the correct webhook setup information, it means it's effective.
6. Initiate a conversation with your created Telegram bot to verify if message processing is functioning correctly.
