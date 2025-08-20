# Story Hero

This is a (work in progress!) Clone Hero companion program that enables "story mode" play throughs.

> [!WARNING]
> This is incomplete & will not function for you today. If you want to help, the [Clone Hero Mod](https://discord.com/channels/424748428451381248/1407613636763193385) is the last step!

## How it works

This app:

1. Reads a "story" definition (a `.story.ts` file, [written in TypeScript](./pkg/storymode/fixtures/gh1.story.ts))
   - eg. a list of the tiers in Guitar Hero, the songs in each tier, and the conditions for unlocking the next tier
2. Watches your score data for changes at `scoredata.bin`, and when there are…
3. Decides which songs should be hidden (ie. are currently locked) because of the story
4. Writes a `hiddensongs.bin` file (next to `songcache.bin`) containing the "Song ID" (the MD5 Hash of the `notes.mid` file, the same ID used in the `songcache.bin`) of each song that should be hidden.

Then, the [unstarted Clone Hero Mod](https://discord.com/channels/424748428451381248/1407613636763193385) will:

1. Watch the `hiddensongs.bin` for changes, and when there are…
2. Hide the specified songs from Clone Hero
