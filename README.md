# likelo
Twitter auto like bot

> Where does likelo's idea come from?
>
> I have a friend named [Arash](https://github.com/ara4sh)([AraaX on twitter](https://twitter.com/Ara4Sh)) who follows me in twitter, 
> every time I post something or even reply to my friend,
> after a short period of time, I get a notification which says Arash liked my post or reply :)
>
> Likelo simulates Arash for your followers :))


## How To Start To Liking, Like A Bot :)


### Dependency
Note: You will need **[Go 1.7](https://golang.org/dl/)** or newer.


### Create a twitter application
While signed in under your Twitter account, visit [https://apps.twitter.com](https://apps.twitter.com).
Click Create an application.
Assign a name, description, and URL to the application. If you don't know the URL of your application yet, that's fine, you can change it later when you know (actually this is not required for likelo). Click the Yes, I agree check box, fill out the CAPTCHA, and click Create your Twitter application.
Once the application has successfully created, visit the Settings tab for the application. Select the Read and Write radio button and click Update this Twitter application's settings. This sets the proper permissions for the application to query and post new tweets to the account.
Visit the Keys and Access Tokens tab. Take note of the consumer key, consumer secret, access token, and access secret — you'll need these for likelo bot. If the access token/secret are not shown, click Create my access token at the bottom of the page.

> #### consumer key, consumer secret, access token, and access secret is required for using in likelo config.


### Download likelo 
download likelo with go get

```bash
go get github.com/arastu/likelo
```

Copy sample config file(.likelo.config.sample.yaml) from likelo folder to your home folder and rename it to .likelo.yaml
```bash
cd $GOPATH/src/github.com/arastu/likelo
cp .likelo.config.sample.yaml $HOME/.likelo.yaml
```
    Then edit .likelo.yaml and add consumer key, consumer secret, access token, and access secret.


### Building from Source
```bash
cd $GOPATH/src/github.com/arastu/likelo
make likelo
```
This will create likelo binary executable file in $GOPATH/src/github.com/arastu/likelo folder.


### Run likelo
```bash
./likelo
Using config file: /Users/arastu/.likelo.yaml
Starting Stream...
```
