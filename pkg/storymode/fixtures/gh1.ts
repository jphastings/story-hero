/// <reference path="../index.d.ts" />

import { MD5Hash, Story, UnlockFunc } from ".."

// Unlock functions

const isGroupCompleted: UnlockFunc = (groupTitle: string): boolean =>
  group(groupTitle).songs.every((songID) => plays(songID)?.playCount)

const isSongPurchased: UnlockFunc = (songID: MD5Hash): boolean => getState().purchasedSongs[songID]

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
  // Black Label Society - Fire it Up
  "85ff914f1d4a5ff1397c2131fa7ead9f": 300,
  // Graveyard BBQ - Cheat on the Church
  "b9adf49b89045e44d81cb2613c027a09": 300,
  // The Bags - Cavemen Rejoice
  "e7dccf97d331192881606de4eabd13e8": 300,
  // The Upper Crust - Eureka, I've Found Love
  "442602f2650596ea97e17ff20dfc4c1d": 300,
  // Shaimus - All of This
  "dbc36e4f24445c28a567732080b22581": 250,

  // TODO: Check ordering from here on

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
  // The Model Sons - Story of My Love
  "2911ebab0b318fbf84e2729e728ec309": 300,
}

// Track definitions

export default {
  "title": "Guitar Hero",
  "groups": [
    {
      "title": "Opening Licks",
      "songs": [
        // Joan Jett & the Blackhearts - I Love Rock 'n Roll
        "806a20bbb49cafb16275c29a6090c99a",
        // Ramones - I Wanna Be Sedated
        "4776137b3a1a4ff1481c39bd7035d27a",
        // White Zombie - Thunder Kiss '65
        "ed3ac5473b8ce5afe02774b6ecba10ad",
        // Deep Purple - Smoke on the Water
        "422c6ecb99bea19b6467c3f8c1e0219a",
        // Bad Religion - Infected
        "d46b1259fc3278ee69cd0c62732668b4"
      ]
    },
    {
      "title": "Axe-Grinders",
      "songs": [
        // Black Sabbath - Iron Man
        "15cb2d11e87f0cffdb0b04f291779117",
        // Boston - More Than a Feeling
        "cc7e9d33fd6a4eb39ac64a5d7def7df8",
        // Judas Priest - You've Got Another Thing Comin'
        "b78ed8db5858aaefd0545adfc6d46e6b",
        // Franz Ferdinand - Take Me Out
        "9d2dbec92e065700867c169c16b4fce3",
        // ZZ Top - Sharp Dressed Man
        "778eb140327326172d256528bbdec9ff"
      ],
      "isUnlocked": (_) => isGroupCompleted('Opening Licks')
    },
    {
      "title": "Thrash and Burn",
      "songs": [
        // Queen - Killer Queen
        "6bc631e3ae627c6d4e6778e203a12919",
        // The Exies - Hey You
        "a4f556bb54478392c547b02e36115fd5",
        // Incubus - Stellar
        "60db66fa53e3b81bc31733a2c06de835",
        // Burning Brides - Heart Full of Black
        "0085da2cd3526eb741fd71e4ac41bb9c",
        // Megadeth - Symphony of Destruction
        "7ac8287a899ebbf095a836b97968c185"
      ],
      "isUnlocked": (_) => isGroupCompleted('Axe-Grinders')
    }, {
      "title": "Return of the Shred",
      "songs": [
        // David Bowie - Ziggy Stardust
        "ebdfd37ca2bf9bbeb1ec9f741209ba41",
        // Sum 41 - Fat Lip
        "3e4076e8528bf8b507c87b431eafe80a",
        // Audioslave - Cochise
        "3f26d1899cf3d51ebff3513f4f304bd4",
        // The Donnas - Take it Off
        "c0b687867d2ab5222032ea3c52c5af3e",
        // Helmet - Unsung
        "90d19bede0572f4b53800eaa9a5f07dd"
      ],
      "isUnlocked": (_) => isGroupCompleted('Thrash and Burn')
    }, {
      "title": "Fret-Burners",
      "songs": [
        // The Jimi Hendrix Experience - Spanish Castle Magic
        "2aef4a167321b49e21ab1ba976cfb8bf",
        // Red Hot Chili Peppers - Higher Ground
        "71935951b4fe9245eb954656dd77d298",
        // Queens of the Stone Age - No One Knows
        "075e517a6b6d8b6f537ba458f5ed45e9",
        // Motörhead - Ace of Spades
        "43bdc0fc38c07e73a19779aa090ff8d3",
        // Cream - Crossroads
        "847fc809b488b6fe7d2dd24ae71081ed"
      ],
      "isUnlocked": (_) => isGroupCompleted('Return of the Shred')
    }, {
      "title": "Face Melters",
      "songs": [
        // Blue Öyster Cult - Godzilla
        "66ac6bfceb61ab2b74507c60dff8a294",
        // Stevie Ray Vaughan - Texas Flood
        "5568c8ded979b487165995287ee9ae40",
        // The Edgar Winter Group - Frankenstein
        "b869fe0ba0cbfabedb53f1fa95b5784f",
        // Pantera - Cowboys from Hell
        "0023ed150f6e8c83da6fcde3add8f13a",
        // Ozzy Osbourne - Bark at the Moon
        "9334d685bc36a1b75d741355c8378f23"
      ],
      "isUnlocked": (_) => isGroupCompleted('Fret-Burners')
    },
    {
      "title": "Bonus Tracks",
      "songs": Object.keys(songShop),
      "isUnlocked": isSongPurchased
    },
    {
      "title": "Hidden Tracks",
      "songs": [
        // Andraleia Buch - Trippolette
        "783f5e0eb7a332364b3ed4170f341318",
        // Gurney - Graveyard Shift
        "c1e8fcc0c538594cdae40a4bc3bbca70"
      ],
      "isUnlocked": (_) => countMeetingStars(5, true) >= 47
    }
  ]
} as Story
