// Copyright © 2017 Touhid Arastu <touhid.arastu@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package likelo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

type Likelo struct {
	config *viper.Viper
	oauth  *oauth
	client *twitter.Client
}

type oauth struct {
	config *oauth1.Config
	token  *oauth1.Token
}

func (l *Likelo) Run() {
	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		go func(tweet *twitter.Tweet) {
			time.Sleep(l.config.GetDuration("delay"))

			fave := &twitter.FavoriteCreateParams{
				ID: tweet.ID,
			}
			_, _, err := l.client.Favorites.Create(fave)

			if err != nil {
				log.Warning(err)
			}

			log.Infof("tweet: %s user: %s is liked!", tweet.Text, tweet.User.Name)
		}(tweet)
	}

	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}

	demux.Event = func(event *twitter.Event) {
		log.Infof("new event occurred, event: %v", event)
	}

	fmt.Println("Starting Stream...")

	// FILTER
	filterParams := &twitter.StreamUserParams{
		With:          "followings",
		StallWarnings: twitter.Bool(true),
	}
	stream, err := l.client.Streams.User(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()

}

func (l *Likelo) SetConfig(config *viper.Viper) {
	l.config = config
	consumerKey := config.GetString("twitter.consumer-key")
	consumerSecret := config.GetString("twitter.consumer-secret")
	accessToken := config.GetString("twitter.access-token")
	accessSecret := config.GetString("twitter.access-secret")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	l.oauth = &oauth{
		config: oauth1.NewConfig(consumerKey, consumerSecret),
		token:  oauth1.NewToken(accessToken, accessSecret),
	}

	httpClient := l.oauth.config.Client(oauth1.NoContext, l.oauth.token)

	// Twitter Client
	l.client = twitter.NewClient(httpClient)
}
