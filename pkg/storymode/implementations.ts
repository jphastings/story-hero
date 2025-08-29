import { MD5Hash, Score } from 'story-hero'

function allSongs(): Array<MD5Hash> {
  // TODO: Fix Songs -> songs when deep mapping is complete
  return story().groups.flatMap((g) => g.Songs)
}

function bestAmount(songID: MD5Hash, mapper: (Score) => number): number {
  const play = plays(songID)
  if (!play) {
    return 0
  }

  return Math.max(...Object.values(play.scores).map(mapper))
}

export const allCompleted: SongsJudge = (songIDs: Array<MD5Hash>|undefined) => !!songIDs && songIDs.every((songID) => (plays(songID)?.playCount || 0) > 0)
export const enoughStars: StoryJudge = (stars: number) => (_) => allSongs().reduce((_, songID: MD5Hash) => bestAmount(songID, ofStars), 0) >= stars

export function previousGroupMeets(judge: SongsJudge): UnlockFunc {
  return function() {
    const group = this
    
    // TODO: Fix Title -> title when deep mapping is complete
    const groupIndex = story().groups.findIndex((g) => g.Title == group.title)
    const previousGroup = story().groups[groupIndex - 1]

    // TODO: Fix Songs -> songs when deep mapping is complete
    return judge(previousGroup.Songs)
  }
}

export const lastAreEncores: UnlockFuncFactory = (n: number, encoreJudge: SongsJudge, others?: UnlockFunc) => {
  return function (songID: MD5Hash): boolean {
    const group = this

    const encoreAfter = group.songs.length - n - 1
    if (group.songs.indexOf(songID) <= encoreAfter) {
      return others?.bind(this)(songID) || true
    }

    // TODO: Fix Songs -> songs when deep mapping is complete
    return encoreJudge(group.Songs.slice(0, -n))
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

export const ofStars = (s: Score): number => s.stars
export const ofPercentage = (s: Score): number => s.percentage
export const ofScore = (s: Score): number => s.score
