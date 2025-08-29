defineStory({
  title: "Band Hero",
  groups: [
    {
      title: "Mall Of Fame Tour",
      songs: [
        // Cold War Kids - Hang Me Up to Dry
        "843f25a5dfb40ee6e1f8810ce9777f36",
        // Culture Club - Do You Really Want To Hurt Me
        "f719d4b1e9f12e740acca0c4f93b2ebd",
        // The Turtles - Happy Together
        "744954dc4d81bff5ea7cd5850d3e9899",
        // OK Go - A Million Ways
        "2e4dfc7c7232c02d1299055ff46e1ba2",
        // Evanescence - Bring Me to Life
        "ac4dc8d57656657c1dcc9cd1d8096868",
        // David Bowie - Let's Dance
        "07f5cf0f554fe15b6dae186b34701351",
        // Roy Orbison - Pretty Woman
        "ff8636fe982c99449232b6ee68ea0b03",
      ],
      isUnlocked: lastAreEncores(2, allCompleted) 
    },
    {
      title: "Smoke & Water Festival", 
      songs: [
        // Hilary Duff - So Yesterday
        "6e06b69514812347f04932c80c87af1f",
        // The Bravery - Believe
        "9f59ba89a4ad155131179e5a61927238",
        // Hinder - Lips of an Angel
        "14061f5d64481f3e5aff1c4b93f6581b",
        // N.E.R.D. - Rock Star
        "2ab88413d4d3f3f1b3911c0e8cad2e95",
        // The Go-Go's - Our Lips Are Sealed
        "25bac9b4a05ac35f31c1d6fb19c197c5",
        // Marvin Gaye - I Heard It Through the Grapevine
        "a3985e70164eb511565ee8b3e4ba4a2a",
        // The Jackson 5 - ABC
        "289e8029fc67ae7ea81133c5250f8228",
      ],
      isUnlocked: lastAreEncores(2, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Club La Noza",
      songs: [
        // The All-American Rejects - Dirty Little Secret
        "45f995d6a55a490ca0736fa72dc0186f",
        // Santigold - L.E.S. Artistes
        "51b118257908ffeeedb2ff0cd992efb6",
        // Spice Girls - Wannabe
        "32787158f435696fd7b0904b59f8440d",
        // Tonic - If You Could Only See
        "51d6305317cb7b43d1b60c559e10349a",
        // Everclear - Santa Monica
        "aeeb28c3f7aa55e605726bfd3f55e5cc",
        // Lily Allen - Take What You Take
        "8859016b6f28bfeaca4791cd50c5fcde",
        // Joan Jett - Bad Reputation
        "9425fbdf6e6a22810162cc438233973a"
      ],
      isUnlocked: lastAreEncores(2, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Summer Park Festival",
      songs: [
        // Snow Patrol - Take Back the City
        "98134198310804b17eefd9b5471fed48",
        // Devo - Whip It
        "8a3fc20786ed40a794063accb320c055",
        // The Last Goodnight - Pictures of You
        "9893f5d9d713a4cf42f302ae54f8a212",
        // Ben Harper and the Innocent Criminals - Steal My Kisses
        "1cca5be9154ca29887e28c8c2d7a79b5",
        // Taylor Swift - Love Story
        "db68c753b6f25aef84b638fc0eaf17ef",
        // The Rolling Stones - Honky Tonk Women
        "09b226a735b3118124d00c3afbbb3cbd",
        // Styx - Mr. Roboto
        "9344acd27f5a024bf109857a48c39f9e"
      ],
      isUnlocked: lastAreEncores(2, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Harajuku",
      songs: [
        // Poison - Every Rose Has Its Thorn
        "8f581c4e7e13e1a5957372176f23f2f8",
        // Yellowcard - Love is a Battlefield
        "c3a92cfd893b1ff6f7752d714a34af1c",
        // Corinne Bailey Rae - Put Your Records On
        "488722ff89723188e660b3d6eec340b4",
        // Maroon 5 - She Will Be Loved
        "c689b15b3540333607986b0277569e41",
        // Janet Jackson - Black Cat
        "55f18ea266e78b6c9fede73aab82d1bb",
        // Fall Out Boy - Sugar, We're Goin Down
        "ac87c7e8b6ee407e91d72602642ec763"
      ],
      isUnlocked: lastAreEncores(2, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "La Luz De Madrid",
      songs: [
        // 3 Doors Down - When I'm Gone
        "a7328a44c9bc4df40740cc931d5bd98e",
        // Duffy - Warwick Avenue
        "15c9e4dc4c7d2b33d8f5c87122852797",
        // Big Country - In a Big Country
        "c1665d7191a1527bf944e81f0e822841",
        // Finger Eleven - Paralyzer
        "10433203ec68ba508d18f394db5caeed",
        // Cheap Trick - I Want You to Want Me (Live)
        "fa670f2757c3c2bebd5d959753cd06a0",
        // Katrina and the Waves - Walking on Sunshine
        "e3f8cb7337b3d3d6f1a83b3d23b9ba8c",
        // Angels & Airwaves - The Adventure
        "afb94e8495ece07e625c6929e0d6c692"
      ],
      isUnlocked: lastAreEncores(2, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Everpop Awards",
      songs: [
        // Counting Crows - Angels of the Silences
        "969e445e7ad22d1c3b267b605d1d0b3f",
        // Jesse McCartney - Beautiful Soul
        "20ab7ab296438e6b2fbc31f4194d541b",
        // Alphabeat - Fascination
        "98bef72ff44adf56adfc08050db159bd",
        // Filter - Take a Picture
        "7b621b20b4e34c2e3dc24511bd32f498",
        // No Doubt - Don't Speak
        "2d3a10a75aa12bdbe791a96fce8d01e5",
        // Nelly Furtado - Turn Off the Light
        "9ee8c180926b1c862ad11927f4ff82f4",
        // Village People - Y.M.C.A.
        "4768cb1cae1ed2941f01c9be02693619"
      ],
      isUnlocked: lastAreEncores(2, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Red River Canyon",
      songs: [
        // Dashboard Confessional - Hands Down
        "3974741993c70aec2e3c30840e278f94",
        // Carl Douglas - Kung Fu Fighting
        "7746e94c42cc578276390a8a748013a2",
        // Aly & AJ - Like Whoa
        "f4a5a4c1b130f8dc9857bb7aecc11267",
        // Taylor Swift - Picture to Burn
        "18232b69aaafbc0c26acd330a9ed4087",
        // Taylor Swift - You Belong With Me
        "65680b7042b5d6b927c15fff5de2a02f",
        // Duran Duran - Rio
        "5606dda416ad2e1e7f3461e7a92202ba"
      ],
      isUnlocked: lastAreEncores(1, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Paris",
      songs: [
        // KT Tunstall - Black Horse and the Cherry Tree
        "233d8b5a9a7ef453666679e1f0c8ff2b",
        // Robbie Williams and Kylie Minogue - Kids
        "52e16f9866845a2efd137d97076bde43",
        // Parachute - Back Again
        "a3645930bf71c96682ed2df4d36035e0",
        // No Doubt - Just a Girl
        "3ee4ee95ac053a54c7ebd59b1427cc29",
        // Papa Roach - Lifeline
        "bcc42dc6e7a09af5384428c768b6010b"
      ],
      isUnlocked: lastAreEncores(1, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Amp Orbiter",
      songs: [
        // The Airborne Toxic Event - Gasoline
        "bd5c3ef296f5dd8498d04c92e703678d",
        // The Kooks - Naïve
        "0e87cae1a18579ef639bea1aa4950f68",
        // Joss Stone - You Had Me
        "e4fa770a9a86957d8a35cb95349d33c8",
        // Mighty Mighty Bosstones - The Impression That I Get
        "2043f4877c3299312bcef39b4133adb8"
      ],
      isUnlocked: lastAreEncores(1, allCompleted, previousGroupMeets(allCompleted))
    },
    {
      title: "Hypersphere",
      songs: [
        // Don McLean - American Pie
        "33c619bd67021cdd4340c2d1fc3bc9c8"
      ],
      isUnlocked: previousGroupMeets(allCompleted)
    }
  ]
})
