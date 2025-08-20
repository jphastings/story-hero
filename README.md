# Story Hero

This is a (work in progress!) Clone Hero companion program that enables "story/career mode" play throughs.

> [!WARNING]
> This is incomplete & will not function for you today. If you want to help, the [Clone Hero Mod](https://discord.com/channels/424748428451381248/1407613636763193385) is the last step!

## How it works

This app:

1. Reads a "story" definition (a `.story.ts` file, [written in TypeScript](pkg/storymode/fixtures/gh1.story.ts))
   - eg. a list of the tiers in Guitar Hero, the songs in each tier, and the conditions for unlocking the next tier
2. Watches your score data for changes at `scoredata.bin`, and when there are…
3. Decides which songs should be hidden (ie. are currently locked) because of the story
4. Writes a `hiddensongs.bin` file (next to `songcache.bin`) containing the "Song ID" (the MD5 Hash of the `notes.mid` file, the same ID used in the `songcache.bin`) of each song that should be hidden.

Then, the [unstarted Clone Hero Mod](https://discord.com/channels/424748428451381248/1407613636763193385) will:

1. Watch the `hiddensongs.bin` for changes, and when there are…
2. Hide the specified songs from Clone Hero

## The details

A `.story.ts` definition file is written in TypeScript, this is so that the logic for whether a group of songs is unlocked can be expressed _in code_, rather than in some newly created and increasingly complex config language. See [an example here](pkg/storymode/fixtures/gh1.story.ts).

It contains a "story definition", which includes the name of the story, the groups of songs, and the conditions for unlocking them.

This program reads those TypeScript files, executes them in a sandboxed VM (with access to [specific functions](pkg/index.d.ts) that allow reading scores, stars, percentages, completeness, etc), and follows the [how it works](#how-it-works) steps above.

## Helpers

If you're hacking about in this space, I've also created (mostly complete) [SynalizeIt!](https://www.synalysis.net/)/[Hexinator](https://hexinator.com) grammars for exploring `songcache.bin` ([grammar here](etc/SongCache.grammar)) and `scoredata.bin` ([grammar here](etc/ScoreData.grammar)).

I'll gladly take suggestions for improving these (just raise an issue here on Github, or send an email to the address in the grammar file); there are a bunch of fields I don't understand yet.

## TODO

- [ ] Read story definition files from within the configured Clone Hero song directories
- [ ] Only enable stories you have the songs for
- [ ] Better logging, feedback & reliability
- [ ] OMG tests.
  - Something simple to cover the TypeScript <-> Go conversion would be good
- [ ] Get Actions working (eg. buying songs)
- [ ] A UI to be able to perform Actions on songs that have them defined
- [ ] Get Github Action-built applications working for macOS, Windows & Linux
- [ ] Clear any other `TODO` comments
- [ ] Make the current locale available in the story definition (for translations)
