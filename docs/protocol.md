# Protocol: wfc-gamestats

Internal HTTP defaults to WFC_HTTP_ADDR=:20205. GET /healthz is unauthenticated;
protocol/internal paths require WFC_INTERNAL_KEY. Console HTTP reaches this owner
through wfc-nas; TCP/UDP ports belong directly to their protocol's container.
Runtime frontend/backend net/rpc endpoints stay on loopback inside each TCP
service container. Restarting another protocol does not restart these sockets.

## Mario Kart DS

Game codes AMCE/AMCP/AMCJ, GameSpy game names `mariokartds` (id 1238) and `mariokartdsam`
(id 1262, Automatch). Nothing is worth storing or serving over a read API; no `/api/mkds/v1`
exists.

- NAS: DS authentication/login (`ac`, `rhgamecd` AMCE etc.), then GPCM login.
- GPCM/GPSP: profile login, friend codes, buddy add/status and messages (friend rooms and
  rival lists).
- QR2 and server browser: matchmaking sessions. The DS DWC build does not check the matching
  status field, which is why `wfc-qr2` special-cases AMCE/AMCP/AMCJ.
- NATNEG: NAT traversal between racers; races are peer-to-peer.
- The game makes no GameStats, SAKE or DLS requests. Online ghosts and rankings belong to the
  Wii Mario Kart Channel (MKWii), not the DS game.

A 30-day production log check found no GameStats, SAKE or DLS request for MKDS. That is absence
of captured traffic, not a console capture: if a DS ever sends one for AMC*, wfc-dls and
wfc-gamestats log it with the game name and this decision should be revisited.
