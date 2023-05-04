### entraffle

### Install Atlas

```
curl -sSf https://atlasgo.sh | sh
```

### Run auto migration with Atlas

```
atlas schema apply \
  --to ent://ent/schema \
  --url "sqlite://ent.db?_fk=1" \
  --dev-url "sqlite://file?mode=memory&_fk=1"
```

This is a small program for running a raffle for all Ent Discord users who retweet this: https://twitter.com/entgo_io/status/1653360590465662978

If you haven't joined our Discord yet, we're waiting for you there: https://discord.gg/qZmPgTE6RX

