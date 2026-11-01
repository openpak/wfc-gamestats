# Evidence: wfc-gamestats

## Mario Kart DS online data

Game: Mario Kart DS, game codes AMCE/AMCP/AMCJ, GameSpy game name
`mariokartds` (id 1238) and `mariokartdsam` (id 1262, Automatch); see
`game_list.tsv`. Conclusion: nothing worth storing or serving over a read API.
No `/api/mkds/v1` is implemented.

Message flows MKDS uses on WFC (facts):

- NAS: DS authentication/login (`ac`, `rhgamecd` AMCE etc.) then GPCM login.
- GPCM/GPSP: profile login, friend codes, buddy add/status and buddy messages
  (friend rooms and rival lists are driven by these). Friend-code-linked
  records are only the profile/buddy lists the GPCM service already owns.
- QR2 (heartbeat/availability) and server browser: matchmaking sessions
  (worldwide/regional/friends). The DS DWC build does not check the matching
  status field, which is why `wfc-qr2` special-cases AMCE/AMCP/AMCJ.
- NATNEG: NAT traversal between racers. Races themselves are peer-to-peer.
- No GameStats, SAKE or DLS requests are made by the game. The Mario Kart DS
  ghost/ranking features that Nintendo shipped were local (Time Trial ghosts
  over local wireless); the online ghost and ranking service was the Wii
  Mario Kart Channel (MKWii), not the DS game.

Sources:

- Wikipedia, "Mario Kart DS": online play via Nintendo WFC, up to four
  players, ended 2014-05-20; no mention of ghosts, rankings or downloads.
  https://en.wikipedia.org/wiki/Mario_Kart_DS
- Wikipedia, "Nintendo Wi-Fi Connection": ghost racing and Time Trial
  ranking are attributed only to the Mario Kart Channel (Wii); MKDS is
  listed only as a multiplayer example.
  https://en.wikipedia.org/wiki/Nintendo_Wi-Fi_Connection
- MKWiiki (Tockdom) "Wiimmfi" / Wiimmfi protocol notes: GPCM for login and
  buddy messages, NATNEG for peer connection.
  https://mkwiiki.org/wiki/Wiimmfi
- OpenPak code: `wfc-common/common/game_list.go` (DoesGameNeedExploit),
  `wfc-qr2/qr2/heartbeat.go` (AMC* special case).

Production logs (checked 2026-11-01, last 30 days, all wfc-* containers):
the only mention of MKDS is a single NAS login line (`rhgamecd AMCE`,
2026-10-25) in wfc-dls/wfc-nas output. No GameStats, SAKE or DLS requests for
MKDS exist. This is absence of captured traffic, consistent with the
documentation but not a console capture; if a DS ever sends a GameStats or
DLS request for AMC*, wfc-dls/wfc-gamestats logs it with the game name and
this decision should be revisited.
