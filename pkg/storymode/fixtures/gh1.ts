/// <reference path="../index.d.ts" />

import { MD5Hash, UnlockFunc } from ".."

interface State {
  cash: number,
  purchasedSongs: Record<MD5Hash, boolean>,
}

function updateState(previousState: State | null): State {
  const purchasedSongs = previousState?.purchasedSongs || {}

  const cash = story.groups.reduce(
    (totalCash, group) => totalCash + group.songs.reduce(
      (groupCash, songID) => groupCash + cashForSong(songID)
    , 0)
  , 0)

  return { cash, purchasedSongs }
}

function cashForSong(songID: MD5Hash): number {
  const play = plays(songID)
  if (!play) {
    return 0
  }

  const stars = Object.values(play.scores).reduce(
    (currentMaxStars, score) => Math.max(currentMaxStars, score.stars)
  , 0)

  return payouts[stars] || 0
}

const getState = useState(updateState)

export const isGroupCompleted: UnlockFunc = (groupTitle: string): boolean =>
  group(groupTitle).songs.every((songID) => plays(songID)?.playCount)

export const isSongPurchased: UnlockFunc = (songID: MD5Hash): boolean => getState().purchasedSongs[songID]

const payouts: Record<number,number> = {
  // What about other star ratings?
  4: 300,
  5: 650,
}

const songShop: Record<MD5Hash, number> = {
  // Shaimus - All of This
  "": 250,
  // Anarchy Club - Behind the Mask
  "": 250,
  // Artillery - The Breaking Wheel
  "": 250,
  // The Acro-Brats - Callout
  "": 250,
  // Drist - Decontrol
  "": 250,
  // The Slip - Even Rats
  "": 250,
  // Made in Mexico - Farewell Myth
  "": 250,
  // Din - Fly on the Wall
  "": 250,
  // Freezepop - Get Ready 2 Rokk
  "": 250,
  // Monkey Steals the Peach - Guitar Hero -[
  "": 250,
  // Honest Bob and the Factory-to-Dealer Incentives - Hey
  "": 250,
  // Count Zero - Sail Your Ship By
  "": 250,
  // The Bags - Cavemen Rejoice
  "": 300,
  // Cheat on the Church - Graveyard BBQ
  "": 300,
  // The Upper Crust - Eureka, I've Found Love
  "": 300,
  // Black Label Society - Fire it Up
  "85ff914f1d4a5ff1397c2131fa7ead9f": 300,
  // The Model Sons - Story of My Love
  "": 300,
}
