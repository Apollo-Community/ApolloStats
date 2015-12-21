ApolloStats
--------------------------------------------------------------------------------

Webpage for showing various stats from the Apollo Station SS13 game database.
With heavy inspiration from other servers' webpages such as [Goon](http://goonhub.com/) and [/vg/station](http://ss13.pomf.se/index.php/bans).

TODO
--------------------------------------------------------------------------------

Templates:
- Should try to use existing css from forum.
- Show some nice error pages.

Ban log:
- Format rows as: Player (ip, cid) got bantype until time by admin for reason.

Account item log:
- Format rows as: Player was awarded item, at time.
- Would be nice to show why a player got an item too.

Player deaths:
- Format rows as: Name (job) died at time, in room (pos) with damage.
- Really nice if we could show a heatmap of deaths.
- Ask @HiddenKn how he made his python version.

End round stats:
- Main window displays rows of summaries of ended rounds.
- Format rows as: A round of type (started on time) has ended after length.
- Clicking on a row opens new page with full details of the round.
- Straight copy of the existing end round stats window from the game, if we
  could store it in the DB after a round has ended.
- Copy the stats used on goonhub: Antags, custom AI laws, deaths, round type
  and length.

Game map:
- Huge, zoomable map of the main station (only).
- Store the map as picture tiles?
- Investigate how goon made their map.
- Need to rebuild the map after any new map changes from a commit.
