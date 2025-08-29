import { MD5Hash, Score } from 'story-hero'

function allSongs(): Array<MD5Hash> {
  return story().groups.flatMap((g) => fixDeep(g, 'songs'))
}

function bestAmount(songID: MD5Hash, mapper: (Score) => number): number {
  const play = plays(songID)
  if (!play) {
    return 0
  }

  return Math.max(...Object.values(play.scores).map(mapper))
}

export const allCompleted: SongsJudge = (songIDs: Array<MD5Hash>|undefined) => !!songIDs && songIDs.every((songID) => plays(songID)?.playCount > 0)
export const enoughStars: StoryJudge = (stars: number) => (_) => allSongs().reduce((_, songID: MD5Hash) => bestAmount(songID, ofStars), 0) >= stars

export function previousGroupMeets(judge: SongsJudge): UnlockFunc {
  return function() {
    const group = this
    
    const groupIndex = story().groups.findIndex((g) => fixDeep(g, 'title') == group.title)
    const previousGroup = story().groups[groupIndex - 1]

    const songs = fixDeep(previousGroup, 'songs')
    return judge(songs)
  }
}

export const lastAreEncores: UnlockFuncFactory = (n: number, encoreJudge: SongsJudge, others?: UnlockFunc) => {
  return function (songID: MD5Hash): boolean {
    const group = this

    const songs = fixDeep(group, 'songs')
    const encoreAfter = songs.length - n - 1
    if (songs.indexOf(songID) <= encoreAfter) {
      if (!others) {
        return true
      }
      return others.bind(this)(songID)
    }

    return encoreJudge(songs.slice(0, -n))
  }
}

// A convenience method that returns the number of songs with best stars (across instruments) equal to (exactly: true) or equal/greater than (exactly: false) the number provided
export function countMeeting(needed: number, mapper: (Score) => number, exactly?: boolean): number {
  return allSongs().reduce((_, songID) => {
    const amount = bestAmount(songID, mapper)
    if (exactly && amount == needed || !exactly && amount >= needed) {
      return 1
    } else {
      return 0
    }
  }, 0)
}

export const ofStars = (s: Score): number => fixDeep(s, 'stars')
export const ofPercentage = (s: Score): number => fixDeep(s, 'percentage')
export const ofScore = (s: Score): number => fixDeep(s, 'score')

// TODO: This won't be necessary once deep mapping works. See https://github.com/go-viper/mapstructure/pull/53
function fixDeep(obj, key) {
  if (!obj.hasOwnProperty(key)) {
    key = key.replace(/^./, char => char.toUpperCase())
  }
  return obj[key]
}