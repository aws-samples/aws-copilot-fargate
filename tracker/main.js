const express = require('express');
const EventSource = require('eventsource');
const app = express();
const port = 3000;
const maxResponseSize = 50;

const emojis = new Map();
const strEmoji = (name) => String.fromCodePoint(...name.split("-").map(i => parseInt(i, 16)));

const source = new EventSource("https://stream.emojitracker.com/subscribe/eps");
source.onmessage = (event) => {
  const updates = JSON.parse(event.data);
  for (const [k, v] of Object.entries(updates)) {
    const char = strEmoji(k);
    if (char in emojis) {
      emojis[char] += v;
    } else {
      emojis[char] = v;
    }
  }
};

app.get('/_healthcheck', (req, res) => {
  res.send('healthcheck okay!')
});

app.get('/', (req, res) => {
  const sortable = [];
  for (const [k, v] of Object.entries(emojis)) {
    sortable.push([k, v]);
  }
  sortable.sort((a, b) => b[1] - a[1]);

  const topK = sortable.slice(0, maxResponseSize);
  const response = {
    emojis: [],
  };
  for (const [k, v] of topK) {
    response.emojis.push({
      emoji: k,
      score: v,
    });
  }
  res.json(response);
});

app.listen(port, () => console.log(`Example app listening at http://localhost:${port}`))