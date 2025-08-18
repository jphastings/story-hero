/// <reference path="../index.d.ts" />

import { MD5Hash, UnlockFunc } from ".."

// Unlock functions

export const isGroupCompleted: UnlockFunc = (groupTitle: string): boolean =>
  group(groupTitle).songs.every((songID) => plays(songID)?.playCount)

export const isSongPurchased: UnlockFunc = (songID: MD5Hash): boolean => getState().purchasedSongs[songID]

// State management

interface State {
  cash: number,
  purchasedSongs: Record<MD5Hash, boolean>,
}

const getState = useState((previousState: State | null): State => ({
  purchasedSongs: previousState?.purchasedSongs || {},
  cash: cashIn() - cashOut(Object.keys(previousState?.purchasedSongs || {}))
}))

const cashIn = () => Object.entries(payouts).reduce(
  (sum, [stars, pay]) => sum + countMeetingStars(Number(stars), true) * pay
, 0)

const cashOut = (purchased: Array<MD5Hash>): number =>
  purchased.reduce((sum, songID) => sum + songShop[songID], 0)

// Configuration for this story

const payouts: Record<number,number> = {
  // Stars -> $
  5: 650,
  4: 300,
  3: 100,
}

const songShop: Record<MD5Hash, number> = {
  // Shaimus - All of This
  "dbc36e4f24445c28a567732080b22581": 250,
  // Anarchy Club - Behind the Mask
  "be16e90bb30795ed62209cc8d4ff17e8": 250,
  // Artillery - The Breaking Wheel
  "0ebec2bfea29aeae07eaa1cc44917be7": 250,
  // The Acro-Brats - Callout
  "b61081cd5d5dd7801cd8d62dce8ac279": 250,
  // Drist - Decontrol
  "f6f8a11a2b3bb64bbb56f76c47b52e09": 250,
  // The Slip - Even Rats
  "934e14a4015609d79939f053d61b0562": 250,
  // Made in Mexico - Farewell Myth
  "f843e2bd2706f79d496e4e11a2dee339": 250,
  // Din - Fly on the Wall
  "3f8cf84eedcb278686cfb973a190700c": 250,
  // Freezepop - Get Ready 2 Rokk
  "5810143a0fa57e2b4c3cd2fec25acccb": 250,
  // Monkey Steals the Peach - Guitar Hero
  "86764ca81f7194df0b96900ed28536ad": 250,
  // Honest Bob and the Factory-to-Dealer Incentives - Hey
  "619b15b43c5e60c36285f2a9a50ad786": 250,
  // Count Zero - Sail Your Ship By
  "5849a22be1e7a568edb8d488f02206d4": 250,
  // The Bags - Cavemen Rejoice
  "e7dccf97d331192881606de4eabd13e8": 300,
  // Graveyard BBQ - Cheat on the Church
  "b9adf49b89045e44d81cb2613c027a09": 300,
  // The Upper Crust - Eureka, I've Found Love
  "442602f2650596ea97e17ff20dfc4c1d": 300,
  // Black Label Society - Fire it Up
  "85ff914f1d4a5ff1397c2131fa7ead9f": 300,
  // The Model Sons - Story of My Love
  "2911ebab0b318fbf84e2729e728ec309": 300,
}
